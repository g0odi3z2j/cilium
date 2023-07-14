// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an event window in which scheduled events for the associated Amazon EC2
// instances can run. You can define either a set of time ranges or a cron
// expression when creating the event window, but not both. All event window times
// are in UTC. You can create up to 200 event windows per Amazon Web Services
// Region. When you create the event window, targets (instance IDs, Dedicated Host
// IDs, or tags) are not yet associated with it. To ensure that the event window
// can be used, you must associate one or more targets with it by using the
// AssociateInstanceEventWindow API. Event windows are applicable only for
// scheduled events that stop, reboot, or terminate instances. Event windows are
// not applicable for:
//   - Expedited scheduled events and network maintenance events.
//   - Unscheduled maintenance such as AutoRecovery and unplanned reboots.
//
// For more information, see Define event windows for scheduled events (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html)
// in the Amazon EC2 User Guide.
func (c *Client) CreateInstanceEventWindow(ctx context.Context, params *CreateInstanceEventWindowInput, optFns ...func(*Options)) (*CreateInstanceEventWindowOutput, error) {
	if params == nil {
		params = &CreateInstanceEventWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstanceEventWindow", params, optFns, c.addOperationCreateInstanceEventWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceEventWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceEventWindowInput struct {

	// The cron expression for the event window, for example, * 0-4,20-23 * * 1,5 . If
	// you specify a cron expression, you can't specify a time range. Constraints:
	//   - Only hour and day of the week values are supported.
	//   - For day of the week values, you can specify either integers 0 through 6 , or
	//   alternative single values SUN through SAT .
	//   - The minute, month, and year must be specified by * .
	//   - The hour value must be one or a multiple range, for example, 0-4 or
	//   0-4,20-23 .
	//   - Each hour range must be >= 2 hours, for example, 0-2 or 20-23 .
	//   - The event window must be >= 4 hours. The combined total time ranges in the
	//   event window must be >= 4 hours.
	// For more information about cron expressions, see cron (https://en.wikipedia.org/wiki/Cron)
	// on the Wikipedia website.
	CronExpression *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The name of the event window.
	Name *string

	// The tags to apply to the event window.
	TagSpecifications []types.TagSpecification

	// The time range for the event window. If you specify a time range, you can't
	// specify a cron expression.
	TimeRanges []types.InstanceEventWindowTimeRangeRequest

	noSmithyDocumentSerde
}

type CreateInstanceEventWindowOutput struct {

	// Information about the event window.
	InstanceEventWindow *types.InstanceEventWindow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInstanceEventWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateInstanceEventWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateInstanceEventWindow{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstanceEventWindow(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateInstanceEventWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateInstanceEventWindow",
	}
}
