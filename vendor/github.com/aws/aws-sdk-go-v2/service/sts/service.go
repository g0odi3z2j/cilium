// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// STS provides the API operation methods for making requests to
// AWS Security Token Service. See this package's package overview docs
// for details on the service.
//
// STS methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type STS struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*STS)

// Used for custom request initialization logic
var initRequest func(*STS, *aws.Request)

// Service information constants
const (
	ServiceName = "sts"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the STS client with a config.
//
// Example:
//     // Create a STS client from just a config.
//     svc := sts.New(myConfig)
func New(config aws.Config) *STS {
	var signingName string
	signingRegion := config.Region

	svc := &STS{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2011-06-15",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a STS operation and runs any
// custom request initialization.
func (c *STS) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
