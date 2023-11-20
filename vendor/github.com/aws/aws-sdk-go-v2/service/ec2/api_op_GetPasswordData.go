// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Retrieves the encrypted administrator password for a running Windows instance.
// The Windows password is generated at boot by the EC2Config service or EC2Launch
// scripts (Windows Server 2016 and later). This usually only happens the first
// time an instance is launched. For more information, see EC2Config (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/UsingConfig_WinAMI.html)
// and EC2Launch (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2launch.html)
// in the Amazon EC2 User Guide. For the EC2Config service, the password is not
// generated for rebundled AMIs unless Ec2SetPassword is enabled before bundling.
// The password is encrypted using the key pair that you specified when you
// launched the instance. You must provide the corresponding key pair file. When
// you launch an instance, password generation and encryption may take a few
// minutes. If you try to retrieve the password before it's available, the output
// returns an empty string. We recommend that you wait up to 15 minutes after
// launching an instance before trying to retrieve the generated password.
func (c *Client) GetPasswordData(ctx context.Context, params *GetPasswordDataInput, optFns ...func(*Options)) (*GetPasswordDataOutput, error) {
	if params == nil {
		params = &GetPasswordDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPasswordData", params, optFns, c.addOperationGetPasswordDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPasswordDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPasswordDataInput struct {

	// The ID of the Windows instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type GetPasswordDataOutput struct {

	// The ID of the Windows instance.
	InstanceId *string

	// The password of the instance. Returns an empty string if the password is not
	// available.
	PasswordData *string

	// The time the data was last updated.
	Timestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPasswordDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetPasswordData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetPasswordData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPasswordData"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetPasswordDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPasswordData(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// GetPasswordDataAPIClient is a client that implements the GetPasswordData
// operation.
type GetPasswordDataAPIClient interface {
	GetPasswordData(context.Context, *GetPasswordDataInput, ...func(*Options)) (*GetPasswordDataOutput, error)
}

var _ GetPasswordDataAPIClient = (*Client)(nil)

// PasswordDataAvailableWaiterOptions are waiter options for
// PasswordDataAvailableWaiter
type PasswordDataAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// PasswordDataAvailableWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, PasswordDataAvailableWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetPasswordDataInput, *GetPasswordDataOutput, error) (bool, error)
}

// PasswordDataAvailableWaiter defines the waiters for PasswordDataAvailable
type PasswordDataAvailableWaiter struct {
	client GetPasswordDataAPIClient

	options PasswordDataAvailableWaiterOptions
}

// NewPasswordDataAvailableWaiter constructs a PasswordDataAvailableWaiter.
func NewPasswordDataAvailableWaiter(client GetPasswordDataAPIClient, optFns ...func(*PasswordDataAvailableWaiterOptions)) *PasswordDataAvailableWaiter {
	options := PasswordDataAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = passwordDataAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &PasswordDataAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for PasswordDataAvailable waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *PasswordDataAvailableWaiter) Wait(ctx context.Context, params *GetPasswordDataInput, maxWaitDur time.Duration, optFns ...func(*PasswordDataAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PasswordDataAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *PasswordDataAvailableWaiter) WaitForOutput(ctx context.Context, params *GetPasswordDataInput, maxWaitDur time.Duration, optFns ...func(*PasswordDataAvailableWaiterOptions)) (*GetPasswordDataOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetPasswordData(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for PasswordDataAvailable waiter")
}

func passwordDataAvailableStateRetryable(ctx context.Context, input *GetPasswordDataInput, output *GetPasswordDataOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(PasswordData) > `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetPasswordData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPasswordData",
	}
}
