// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	internalConfig "github.com/aws/aws-sdk-go-v2/internal/configsources"
	smithy "github.com/aws/smithy-go"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net"
	"net/http"
	"time"
)

const ServiceID = "SSO"
const ServiceAPIVersion = "2019-06-10"

// Client provides the API client to make operations call for AWS Single Sign-On.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	setResolvedDefaultsMode(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The optional application specific identifier appended to the User-Agent header.
	AppID string

	// This endpoint will be given as input to an EndpointResolverV2. It is used for
	// providing a custom base endpoint that is subject to modifications by the
	// processing EndpointResolverV2.
	BaseEndpoint *string

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The configuration DefaultsMode that the SDK should use when constructing the
	// clients initial default settings.
	DefaultsMode aws.DefaultsMode

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	//
	// Deprecated: Deprecated: EndpointResolver and WithEndpointResolver. Providing a
	// value for this field will likely prevent you from using any endpoint-related
	// service features released after the introduction of EndpointResolverV2 and
	// BaseEndpoint. To migrate an EndpointResolver implementation that uses a custom
	// endpoint, set the client option BaseEndpoint instead.
	EndpointResolver EndpointResolver

	// Resolves the endpoint used for a particular service. This should be used over
	// the deprecated EndpointResolver
	EndpointResolverV2 EndpointResolverV2

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// RetryMaxAttempts specifies the maximum number attempts an API client will call
	// an operation that fails with a retryable error. A value of 0 is ignored, and
	// will not be used to configure the API client created default retryer, or modify
	// per operation call's retry max attempts. When creating a new API Clients this
	// member will only be used if the Retryer Options member is nil. This value will
	// be ignored if Retryer is not nil. If specified in an operation call's functional
	// options with a value that is different than the constructed client's Options,
	// the Client's Retryer will be wrapped to use the operation's specific
	// RetryMaxAttempts value.
	RetryMaxAttempts int

	// RetryMode specifies the retry mode the API client will be created with, if
	// Retryer option is not also specified. When creating a new API Clients this
	// member will only be used if the Retryer Options member is nil. This value will
	// be ignored if Retryer is not nil. Currently does not support per operation call
	// overrides, may in the future.
	RetryMode aws.RetryMode

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer. The kind of
	// default retry created by the API client can be changed with the RetryMode
	// option.
	Retryer aws.Retryer

	// The RuntimeEnvironment configuration, only populated if the DefaultsMode is set
	// to DefaultsModeAuto and is initialized using config.LoadDefaultConfig . You
	// should not populate this structure programmatically, or rely on the values here
	// within your applications.
	RuntimeEnvironment aws.RuntimeEnvironment

	// The initial DefaultsMode used when the client options were constructed. If the
	// DefaultsMode was set to aws.DefaultsModeAuto this will store what the resolved
	// value was at that point in time. Currently does not support per operation call
	// overrides, may in the future.
	resolvedDefaultsMode aws.DefaultsMode

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

// WithAPIOptions returns a functional option for setting the Client's APIOptions
// option.
func WithAPIOptions(optFns ...func(*middleware.Stack) error) func(*Options) {
	return func(o *Options) {
		o.APIOptions = append(o.APIOptions, optFns...)
	}
}

// Deprecated: EndpointResolver and WithEndpointResolver. Providing a value for
// this field will likely prevent you from using any endpoint-related service
// features released after the introduction of EndpointResolverV2 and BaseEndpoint.
// To migrate an EndpointResolver implementation that uses a custom endpoint, set
// the client option BaseEndpoint instead.
func WithEndpointResolver(v EndpointResolver) func(*Options) {
	return func(o *Options) {
		o.EndpointResolver = v
	}
}

// WithEndpointResolverV2 returns a functional option for setting the Client's
// EndpointResolverV2 option.
func WithEndpointResolverV2(v EndpointResolverV2) func(*Options) {
	return func(o *Options) {
		o.EndpointResolverV2 = v
	}
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)

	return to
}
func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	ctx = middleware.ClearStackValues(ctx)
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()
	resolveEndpointResolverV2(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	finalizeRetryMaxAttemptOptions(&options, *c)

	finalizeClientEndpointResolverOptions(&options)

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}

type noSmithyDocumentSerde = smithydocument.NoSerde

type legacyEndpointContextSetter struct {
	LegacyResolver EndpointResolver
}

func (*legacyEndpointContextSetter) ID() string {
	return "legacyEndpointContextSetter"
}

func (m *legacyEndpointContextSetter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.LegacyResolver != nil {
		ctx = awsmiddleware.SetRequiresLegacyEndpoints(ctx, true)
	}

	return next.HandleInitialize(ctx, in)

}
func addlegacyEndpointContextSetter(stack *middleware.Stack, o Options) error {
	return stack.Initialize.Add(&legacyEndpointContextSetter{
		LegacyResolver: o.EndpointResolver,
	}, middleware.Before)
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

func setResolvedDefaultsMode(o *Options) {
	if len(o.resolvedDefaultsMode) > 0 {
		return
	}

	var mode aws.DefaultsMode
	mode.SetFromString(string(o.DefaultsMode))

	if mode == aws.DefaultsModeAuto {
		mode = defaults.ResolveDefaultsModeAuto(o.Region, o.RuntimeEnvironment)
	}

	o.resolvedDefaultsMode = mode
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:             cfg.Region,
		DefaultsMode:       cfg.DefaultsMode,
		RuntimeEnvironment: cfg.RuntimeEnvironment,
		HTTPClient:         cfg.HTTPClient,
		Credentials:        cfg.Credentials,
		APIOptions:         cfg.APIOptions,
		Logger:             cfg.Logger,
		ClientLogMode:      cfg.ClientLogMode,
		AppID:              cfg.AppID,
	}
	resolveAWSRetryerProvider(cfg, &opts)
	resolveAWSRetryMaxAttempts(cfg, &opts)
	resolveAWSRetryMode(cfg, &opts)
	resolveAWSEndpointResolver(cfg, &opts)
	resolveUseDualStackEndpoint(cfg, &opts)
	resolveUseFIPSEndpoint(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	var buildable *awshttp.BuildableClient

	if o.HTTPClient != nil {
		var ok bool
		buildable, ok = o.HTTPClient.(*awshttp.BuildableClient)
		if !ok {
			return
		}
	} else {
		buildable = awshttp.NewBuildableClient()
	}

	modeConfig, err := defaults.GetModeConfiguration(o.resolvedDefaultsMode)
	if err == nil {
		buildable = buildable.WithDialerOptions(func(dialer *net.Dialer) {
			if dialerTimeout, ok := modeConfig.GetConnectTimeout(); ok {
				dialer.Timeout = dialerTimeout
			}
		})

		buildable = buildable.WithTransportOptions(func(transport *http.Transport) {
			if tlsHandshakeTimeout, ok := modeConfig.GetTLSNegotiationTimeout(); ok {
				transport.TLSHandshakeTimeout = tlsHandshakeTimeout
			}
		})
	}

	o.HTTPClient = buildable
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}

	if len(o.RetryMode) == 0 {
		modeConfig, err := defaults.GetModeConfiguration(o.resolvedDefaultsMode)
		if err == nil {
			o.RetryMode = modeConfig.RetryMode
		}
	}
	if len(o.RetryMode) == 0 {
		o.RetryMode = aws.RetryModeStandard
	}

	var standardOptions []func(*retry.StandardOptions)
	if v := o.RetryMaxAttempts; v != 0 {
		standardOptions = append(standardOptions, func(so *retry.StandardOptions) {
			so.MaxAttempts = v
		})
	}

	switch o.RetryMode {
	case aws.RetryModeAdaptive:
		var adaptiveOptions []func(*retry.AdaptiveModeOptions)
		if len(standardOptions) != 0 {
			adaptiveOptions = append(adaptiveOptions, func(ao *retry.AdaptiveModeOptions) {
				ao.StandardOptions = append(ao.StandardOptions, standardOptions...)
			})
		}
		o.Retryer = retry.NewAdaptiveMode(adaptiveOptions...)

	default:
		o.Retryer = retry.NewStandard(standardOptions...)
	}
}

func resolveAWSRetryerProvider(cfg aws.Config, o *Options) {
	if cfg.Retryer == nil {
		return
	}
	o.Retryer = cfg.Retryer()
}

func resolveAWSRetryMode(cfg aws.Config, o *Options) {
	if len(cfg.RetryMode) == 0 {
		return
	}
	o.RetryMode = cfg.RetryMode
}
func resolveAWSRetryMaxAttempts(cfg aws.Config, o *Options) {
	if cfg.RetryMaxAttempts == 0 {
		return
	}
	o.RetryMaxAttempts = cfg.RetryMaxAttempts
}

func finalizeRetryMaxAttemptOptions(o *Options, client Client) {
	if v := o.RetryMaxAttempts; v == 0 || v == client.options.RetryMaxAttempts {
		return
	}

	o.Retryer = retry.AddWithMaxAttempts(o.Retryer, o.RetryMaxAttempts)
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil && cfg.EndpointResolverWithOptions == nil {
		return
	}
	o.EndpointResolver = withEndpointResolver(cfg.EndpointResolver, cfg.EndpointResolverWithOptions)
}

func addClientUserAgent(stack *middleware.Stack, options Options) error {
	if err := awsmiddleware.AddSDKAgentKeyValue(awsmiddleware.APIMetadata, "sso", goModuleVersion)(stack); err != nil {
		return err
	}

	if len(options.AppID) > 0 {
		return awsmiddleware.AddSDKAgentKey(awsmiddleware.ApplicationIdentifier, options.AppID)(stack)
	}

	return nil
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) error {
	mw := v4.NewSignHTTPRequestMiddleware(v4.SignHTTPRequestMiddlewareOptions{
		CredentialsProvider: o.Credentials,
		Signer:              o.HTTPSignerV4,
		LogSigning:          o.ClientLogMode.IsSigning(),
	})
	return stack.Finalize.Add(mw, middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time, optFns ...func(*v4.SignerOptions)) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = newDefaultV4Signer(*o)
}

func newDefaultV4Signer(o Options) *v4.Signer {
	return v4.NewSigner(func(so *v4.SignerOptions) {
		so.Logger = o.Logger
		so.LogSigning = o.ClientLogMode.IsSigning()
	})
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer:          o.Retryer,
		LogRetryAttempts: o.ClientLogMode.IsRetries(),
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

// resolves dual-stack endpoint configuration
func resolveUseDualStackEndpoint(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := internalConfig.ResolveUseDualStackEndpoint(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.EndpointOptions.UseDualStackEndpoint = value
	}
	return nil
}

// resolves FIPS endpoint configuration
func resolveUseFIPSEndpoint(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := internalConfig.ResolveUseFIPSEndpoint(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.EndpointOptions.UseFIPSEndpoint = value
	}
	return nil
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) error {
	return awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return awshttp.AddResponseErrorMiddleware(stack)
}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}

type endpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*endpointDisableHTTPSMiddleware) ID() string {
	return "endpointDisableHTTPSMiddleware"
}

func (m *endpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointDisableHTTPS && !smithyhttp.GetHostnameImmutable(ctx) {
		req.URL.Scheme = "http"
	}

	return next.HandleSerialize(ctx, in)

}
func addendpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&endpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "OperationSerializer", middleware.Before)
}
