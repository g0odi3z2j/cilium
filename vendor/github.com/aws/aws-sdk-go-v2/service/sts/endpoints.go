// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalConfig "github.com/aws/aws-sdk-go-v2/internal/configsources"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints/awsrulesfn"
	internalendpoints "github.com/aws/aws-sdk-go-v2/service/sts/internal/endpoints"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// EndpointResolverOptions is the service endpoint resolver options
type EndpointResolverOptions = internalendpoints.Options

// EndpointResolver interface for resolving service endpoints.
type EndpointResolver interface {
	ResolveEndpoint(region string, options EndpointResolverOptions) (aws.Endpoint, error)
}

var _ EndpointResolver = &internalendpoints.Resolver{}

// NewDefaultEndpointResolver constructs a new service endpoint resolver
func NewDefaultEndpointResolver() *internalendpoints.Resolver {
	return internalendpoints.New()
}

// EndpointResolverFunc is a helper utility that wraps a function so it satisfies
// the EndpointResolver interface. This is useful when you want to add additional
// endpoint resolving logic, or stub out specific endpoints with custom values.
type EndpointResolverFunc func(region string, options EndpointResolverOptions) (aws.Endpoint, error)

func (fn EndpointResolverFunc) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	return fn(region, options)
}

// EndpointResolverFromURL returns an EndpointResolver configured using the
// provided endpoint url. By default, the resolved endpoint resolver uses the
// client region as signing region, and the endpoint source is set to
// EndpointSourceCustom.You can provide functional options to configure endpoint
// values for the resolved endpoint.
func EndpointResolverFromURL(url string, optFns ...func(*aws.Endpoint)) EndpointResolver {
	e := aws.Endpoint{URL: url, Source: aws.EndpointSourceCustom}
	for _, fn := range optFns {
		fn(&e)
	}

	return EndpointResolverFunc(
		func(region string, options EndpointResolverOptions) (aws.Endpoint, error) {
			if len(e.SigningRegion) == 0 {
				e.SigningRegion = region
			}
			return e, nil
		},
	)
}

type ResolveEndpoint struct {
	Resolver EndpointResolver
	Options  EndpointResolverOptions
}

func (*ResolveEndpoint) ID() string {
	return "ResolveEndpoint"
}

func (m *ResolveEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if !awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.Resolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	eo := m.Options
	eo.Logger = middleware.GetLogger(ctx)

	var endpoint aws.Endpoint
	endpoint, err = m.Resolver.ResolveEndpoint(awsmiddleware.GetRegion(ctx), eo)
	if err != nil {
		nf := (&aws.EndpointNotFoundError{})
		if errors.As(err, &nf) {
			ctx = awsmiddleware.SetRequiresLegacyEndpoints(ctx, false)
			return next.HandleSerialize(ctx, in)
		}
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL, err = url.Parse(endpoint.URL)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}

	if len(awsmiddleware.GetSigningName(ctx)) == 0 {
		signingName := endpoint.SigningName
		if len(signingName) == 0 {
			signingName = "sts"
		}
		ctx = awsmiddleware.SetSigningName(ctx, signingName)
	}
	ctx = awsmiddleware.SetEndpointSource(ctx, endpoint.Source)
	ctx = smithyhttp.SetHostnameImmutable(ctx, endpoint.HostnameImmutable)
	ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)
	ctx = awsmiddleware.SetPartitionID(ctx, endpoint.PartitionID)
	return next.HandleSerialize(ctx, in)
}
func addResolveEndpointMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&ResolveEndpoint{
		Resolver: o.EndpointResolver,
		Options:  o.EndpointOptions,
	}, "OperationSerializer", middleware.Before)
}

func removeResolveEndpointMiddleware(stack *middleware.Stack) error {
	_, err := stack.Serialize.Remove((&ResolveEndpoint{}).ID())
	return err
}

type wrappedEndpointResolver struct {
	awsResolver aws.EndpointResolverWithOptions
}

func (w *wrappedEndpointResolver) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	return w.awsResolver.ResolveEndpoint(ServiceID, region, options)
}

type awsEndpointResolverAdaptor func(service, region string) (aws.Endpoint, error)

func (a awsEndpointResolverAdaptor) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return a(service, region)
}

var _ aws.EndpointResolverWithOptions = awsEndpointResolverAdaptor(nil)

// withEndpointResolver returns an aws.EndpointResolverWithOptions that first delegates endpoint resolution to the awsResolver.
// If awsResolver returns aws.EndpointNotFoundError error, the v1 resolver middleware will swallow the error,
// and set an appropriate context flag such that fallback will occur when EndpointResolverV2 is invoked
// via its middleware.
//
// If another error (besides aws.EndpointNotFoundError) is returned, then that error will be propagated.
func withEndpointResolver(awsResolver aws.EndpointResolver, awsResolverWithOptions aws.EndpointResolverWithOptions) EndpointResolver {
	var resolver aws.EndpointResolverWithOptions

	if awsResolverWithOptions != nil {
		resolver = awsResolverWithOptions
	} else if awsResolver != nil {
		resolver = awsEndpointResolverAdaptor(awsResolver.ResolveEndpoint)
	}

	return &wrappedEndpointResolver{
		awsResolver: resolver,
	}
}

func finalizeClientEndpointResolverOptions(options *Options) {
	options.EndpointOptions.LogDeprecated = options.ClientLogMode.IsDeprecatedUsage()

	if len(options.EndpointOptions.ResolvedRegion) == 0 {
		const fipsInfix = "-fips-"
		const fipsPrefix = "fips-"
		const fipsSuffix = "-fips"

		if strings.Contains(options.Region, fipsInfix) ||
			strings.Contains(options.Region, fipsPrefix) ||
			strings.Contains(options.Region, fipsSuffix) {
			options.EndpointOptions.ResolvedRegion = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
				options.Region, fipsInfix, "-"), fipsPrefix, ""), fipsSuffix, "")
			options.EndpointOptions.UseFIPSEndpoint = aws.FIPSEndpointStateEnabled
		}
	}

}

func resolveEndpointResolverV2(options *Options) {
	if options.EndpointResolverV2 == nil {
		options.EndpointResolverV2 = NewDefaultEndpointResolverV2()
	}
}

func resolveBaseEndpoint(cfg aws.Config, o *Options) {
	if cfg.BaseEndpoint != nil {
		o.BaseEndpoint = cfg.BaseEndpoint
	}

	_, g := os.LookupEnv("AWS_ENDPOINT_URL")
	_, s := os.LookupEnv("AWS_ENDPOINT_URL_STS")

	if g && !s {
		return
	}

	value, found, err := internalConfig.ResolveServiceBaseEndpoint(context.Background(), "STS", cfg.ConfigSources)
	if found && err == nil {
		o.BaseEndpoint = &value
	}
}

// EndpointParameters provides the parameters that influence how endpoints are
// resolved.
type EndpointParameters struct {
	// The AWS region used to dispatch the request.
	//
	// Parameter is
	// required.
	//
	// AWS::Region
	Region *string

	// When true, use the dual-stack endpoint. If the configured endpoint does not
	// support dual-stack, dispatching the request MAY return an error.
	//
	// Defaults to
	// false if no value is provided.
	//
	// AWS::UseDualStack
	UseDualStack *bool

	// When true, send this request to the FIPS-compliant regional endpoint. If the
	// configured endpoint does not have a FIPS compliant endpoint, dispatching the
	// request will return an error.
	//
	// Defaults to false if no value is
	// provided.
	//
	// AWS::UseFIPS
	UseFIPS *bool

	// Override the endpoint used to send this request
	//
	// Parameter is
	// required.
	//
	// SDK::Endpoint
	Endpoint *string

	// Whether the global endpoint should be used, rather then the regional endpoint
	// for us-east-1.
	//
	// Defaults to false if no value is
	// provided.
	//
	// AWS::STS::UseGlobalEndpoint
	UseGlobalEndpoint *bool
}

// ValidateRequired validates required parameters are set.
func (p EndpointParameters) ValidateRequired() error {
	if p.UseDualStack == nil {
		return fmt.Errorf("parameter UseDualStack is required")
	}

	if p.UseFIPS == nil {
		return fmt.Errorf("parameter UseFIPS is required")
	}

	if p.UseGlobalEndpoint == nil {
		return fmt.Errorf("parameter UseGlobalEndpoint is required")
	}

	return nil
}

// WithDefaults returns a shallow copy of EndpointParameterswith default values
// applied to members where applicable.
func (p EndpointParameters) WithDefaults() EndpointParameters {
	if p.UseDualStack == nil {
		p.UseDualStack = ptr.Bool(false)
	}

	if p.UseFIPS == nil {
		p.UseFIPS = ptr.Bool(false)
	}

	if p.UseGlobalEndpoint == nil {
		p.UseGlobalEndpoint = ptr.Bool(false)
	}
	return p
}

// EndpointResolverV2 provides the interface for resolving service endpoints.
type EndpointResolverV2 interface {
	// ResolveEndpoint attempts to resolve the endpoint with the provided options,
	// returning the endpoint if found. Otherwise an error is returned.
	ResolveEndpoint(ctx context.Context, params EndpointParameters) (
		smithyendpoints.Endpoint, error,
	)
}

// resolver provides the implementation for resolving endpoints.
type resolver struct{}

func NewDefaultEndpointResolverV2() EndpointResolverV2 {
	return &resolver{}
}

// ResolveEndpoint attempts to resolve the endpoint with the provided options,
// returning the endpoint if found. Otherwise an error is returned.
func (r *resolver) ResolveEndpoint(
	ctx context.Context, params EndpointParameters,
) (
	endpoint smithyendpoints.Endpoint, err error,
) {
	params = params.WithDefaults()
	if err = params.ValidateRequired(); err != nil {
		return endpoint, fmt.Errorf("endpoint parameters are not valid, %w", err)
	}
	_UseDualStack := *params.UseDualStack
	_UseFIPS := *params.UseFIPS
	_UseGlobalEndpoint := *params.UseGlobalEndpoint

	if _UseGlobalEndpoint == true {
		if !(params.Endpoint != nil) {
			if exprVal := params.Region; exprVal != nil {
				_Region := *exprVal
				_ = _Region
				if exprVal := awsrulesfn.GetPartition(_Region); exprVal != nil {
					_PartitionResult := *exprVal
					_ = _PartitionResult
					if _UseFIPS == false {
						if _UseDualStack == false {
							if _Region == "ap-northeast-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "ap-south-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "ap-southeast-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "ap-southeast-2" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "aws-global" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "ca-central-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "eu-central-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "eu-north-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "eu-west-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "eu-west-2" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "eu-west-3" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "sa-east-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "us-east-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "us-east-2" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "us-west-1" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							if _Region == "us-west-2" {
								uriString := "https://sts.amazonaws.com"

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
											{
												SchemeID: "aws.auth#sigv4",
												SignerProperties: func() smithy.Properties {
													var sp smithy.Properties
													smithyhttp.SetSigV4SigningName(&sp, "sts")
													smithyhttp.SetSigV4ASigningName(&sp, "sts")

													smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
													return sp
												}(),
											},
										})
										return out
									}(),
								}, nil
							}
							uriString := func() string {
								var out strings.Builder
								out.WriteString("https://sts.")
								out.WriteString(_Region)
								out.WriteString(".")
								out.WriteString(_PartitionResult.DnsSuffix)
								return out.String()
							}()

							uri, err := url.Parse(uriString)
							if err != nil {
								return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
							}

							return smithyendpoints.Endpoint{
								URI:     *uri,
								Headers: http.Header{},
								Properties: func() smithy.Properties {
									var out smithy.Properties
									smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
										{
											SchemeID: "aws.auth#sigv4",
											SignerProperties: func() smithy.Properties {
												var sp smithy.Properties
												smithyhttp.SetSigV4SigningName(&sp, "sts")
												smithyhttp.SetSigV4ASigningName(&sp, "sts")

												smithyhttp.SetSigV4SigningRegion(&sp, _Region)
												return sp
											}(),
										},
									})
									return out
								}(),
							}, nil
						}
					}
				}
			}
		}
	}
	if exprVal := params.Endpoint; exprVal != nil {
		_Endpoint := *exprVal
		_ = _Endpoint
		if _UseFIPS == true {
			return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: FIPS and custom endpoint are not supported")
		}
		if _UseDualStack == true {
			return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: Dualstack and custom endpoint are not supported")
		}
		uriString := _Endpoint

		uri, err := url.Parse(uriString)
		if err != nil {
			return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
		}

		return smithyendpoints.Endpoint{
			URI:     *uri,
			Headers: http.Header{},
		}, nil
	}
	if exprVal := params.Region; exprVal != nil {
		_Region := *exprVal
		_ = _Region
		if exprVal := awsrulesfn.GetPartition(_Region); exprVal != nil {
			_PartitionResult := *exprVal
			_ = _PartitionResult
			if _UseFIPS == true {
				if _UseDualStack == true {
					if true == _PartitionResult.SupportsFIPS {
						if true == _PartitionResult.SupportsDualStack {
							uriString := func() string {
								var out strings.Builder
								out.WriteString("https://sts-fips.")
								out.WriteString(_Region)
								out.WriteString(".")
								out.WriteString(_PartitionResult.DualStackDnsSuffix)
								return out.String()
							}()

							uri, err := url.Parse(uriString)
							if err != nil {
								return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
							}

							return smithyendpoints.Endpoint{
								URI:     *uri,
								Headers: http.Header{},
							}, nil
						}
					}
					return endpoint, fmt.Errorf("endpoint rule error, %s", "FIPS and DualStack are enabled, but this partition does not support one or both")
				}
			}
			if _UseFIPS == true {
				if _PartitionResult.SupportsFIPS == true {
					if _PartitionResult.Name == "aws-us-gov" {
						uriString := func() string {
							var out strings.Builder
							out.WriteString("https://sts.")
							out.WriteString(_Region)
							out.WriteString(".amazonaws.com")
							return out.String()
						}()

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
						}, nil
					}
					uriString := func() string {
						var out strings.Builder
						out.WriteString("https://sts-fips.")
						out.WriteString(_Region)
						out.WriteString(".")
						out.WriteString(_PartitionResult.DnsSuffix)
						return out.String()
					}()

					uri, err := url.Parse(uriString)
					if err != nil {
						return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
					}

					return smithyendpoints.Endpoint{
						URI:     *uri,
						Headers: http.Header{},
					}, nil
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "FIPS is enabled but this partition does not support FIPS")
			}
			if _UseDualStack == true {
				if true == _PartitionResult.SupportsDualStack {
					uriString := func() string {
						var out strings.Builder
						out.WriteString("https://sts.")
						out.WriteString(_Region)
						out.WriteString(".")
						out.WriteString(_PartitionResult.DualStackDnsSuffix)
						return out.String()
					}()

					uri, err := url.Parse(uriString)
					if err != nil {
						return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
					}

					return smithyendpoints.Endpoint{
						URI:     *uri,
						Headers: http.Header{},
					}, nil
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "DualStack is enabled but this partition does not support DualStack")
			}
			if _Region == "aws-global" {
				uriString := "https://sts.amazonaws.com"

				uri, err := url.Parse(uriString)
				if err != nil {
					return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
				}

				return smithyendpoints.Endpoint{
					URI:     *uri,
					Headers: http.Header{},
					Properties: func() smithy.Properties {
						var out smithy.Properties
						smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
							{
								SchemeID: "aws.auth#sigv4",
								SignerProperties: func() smithy.Properties {
									var sp smithy.Properties
									smithyhttp.SetSigV4SigningName(&sp, "sts")
									smithyhttp.SetSigV4ASigningName(&sp, "sts")

									smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
									return sp
								}(),
							},
						})
						return out
					}(),
				}, nil
			}
			uriString := func() string {
				var out strings.Builder
				out.WriteString("https://sts.")
				out.WriteString(_Region)
				out.WriteString(".")
				out.WriteString(_PartitionResult.DnsSuffix)
				return out.String()
			}()

			uri, err := url.Parse(uriString)
			if err != nil {
				return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
			}

			return smithyendpoints.Endpoint{
				URI:     *uri,
				Headers: http.Header{},
			}, nil
		}
		return endpoint, fmt.Errorf("Endpoint resolution failed. Invalid operation or environment input.")
	}
	return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: Missing Region")
}

type endpointParamsBinder interface {
	bindEndpointParams(*EndpointParameters)
}

func bindEndpointParams(input interface{}, options Options) *EndpointParameters {
	params := &EndpointParameters{}

	params.Region = aws.String(endpoints.MapFIPSRegion(options.Region))
	params.UseDualStack = aws.Bool(options.EndpointOptions.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled)
	params.UseFIPS = aws.Bool(options.EndpointOptions.UseFIPSEndpoint == aws.FIPSEndpointStateEnabled)
	params.Endpoint = options.BaseEndpoint

	if b, ok := input.(endpointParamsBinder); ok {
		b.bindEndpointParams(params)
	}

	return params
}

type resolveEndpointV2Middleware struct {
	options Options
}

func (*resolveEndpointV2Middleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *resolveEndpointV2Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.options.EndpointResolverV2 == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := bindEndpointParams(getOperationInput(ctx), m.options)
	endpt, err := m.options.EndpointResolverV2.ResolveEndpoint(ctx, *params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	if endpt.URI.RawPath == "" && req.URL.RawPath != "" {
		endpt.URI.RawPath = endpt.URI.Path
	}
	req.URL.Scheme = endpt.URI.Scheme
	req.URL.Host = endpt.URI.Host
	req.URL.Path = smithyhttp.JoinPath(endpt.URI.Path, req.URL.Path)
	req.URL.RawPath = smithyhttp.JoinPath(endpt.URI.RawPath, req.URL.RawPath)
	for k := range endpt.Headers {
		req.Header.Set(k, endpt.Headers.Get(k))
	}

	rscheme := getResolvedAuthScheme(ctx)
	if rscheme == nil {
		return out, metadata, fmt.Errorf("no resolved auth scheme")
	}

	opts, _ := smithyauth.GetAuthOptions(&endpt.Properties)
	for _, o := range opts {
		rscheme.SignerProperties.SetAll(&o.SignerProperties)
	}

	return next.HandleFinalize(ctx, in)
}
