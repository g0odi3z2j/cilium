// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of temporary security credentials that you can use to access
// Amazon Web Services resources that you might not normally have access to. These
// temporary credentials consist of an access key ID, a secret access key, and a
// security token. Typically, you use AssumeRole within your account or for
// cross-account access. For a comparison of AssumeRole with other API operations
// that produce temporary credentials, see Requesting Temporary Security
// Credentials
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the Amazon Web Services STS API operations
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide. Permissions The temporary security credentials created by
// AssumeRole can be used to make API calls to any Amazon Web Services service with
// the following exception: You cannot call the Amazon Web Services STS
// GetFederationToken or GetSessionToken API operations. (Optional) You can pass
// inline or managed session policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// to this operation. You can pass a single JSON policy document to use as an
// inline session policy. You can also specify up to 10 managed policies to use as
// managed session policies. The plaintext that you use for both inline and managed
// session policies can't exceed 2,048 characters. Passing policies to this
// operation returns new temporary credentials. The resulting session's permissions
// are the intersection of the role's identity-based policy and the session
// policies. You can use the role's temporary credentials in subsequent Amazon Web
// Services API calls to access resources in the account that owns the role. You
// cannot use session policies to grant more permissions than those allowed by the
// identity-based policy of the role that is being assumed. For more information,
// see Session Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// in the IAM User Guide. When you create a role, you create two policies: A role
// trust policy that specifies who can assume the role and a permissions policy
// that specifies what can be done with the role. You specify the trusted principal
// who is allowed to assume the role in the role trust policy. To assume a role
// from a different account, your Amazon Web Services account must be trusted by
// the role. The trust relationship is defined in the role's trust policy when the
// role is created. That trust policy states which accounts are allowed to delegate
// that access to users in the account. A user who wants to access a role in a
// different account must also have permissions that are delegated from the user
// account administrator. The administrator must attach a policy that allows the
// user to call AssumeRole for the ARN of the role in the other account. To allow a
// user to assume a role in the same account, you can do either of the
// following:
//
// * Attach a policy to the user that allows the user to call
// AssumeRole (as long as the role's trust policy trusts the account).
//
// * Add the
// user as a principal directly in the role's trust policy.
//
// You can do either
// because the role’s trust policy acts as an IAM resource-based policy. When a
// resource-based policy grants access to a principal in the same account, no
// additional identity-based policy is required. For more information about trust
// policies and resource-based policies, see IAM Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html) in the
// IAM User Guide. Tags (Optional) You can pass tag key-value pairs to your
// session. These tags are called session tags. For more information about session
// tags, see Passing Session Tags in STS
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html) in the
// IAM User Guide. An administrator must grant you the permissions necessary to
// pass session tags. The administrator can also create granular permissions to
// allow you to pass only specific session tags. For more information, see
// Tutorial: Using Tags for Attribute-Based Access Control
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html)
// in the IAM User Guide. You can set the session tags as transitive. Transitive
// tags persist during role chaining. For more information, see Chaining Roles with
// Session Tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining)
// in the IAM User Guide. Using MFA with AssumeRole (Optional) You can include
// multi-factor authentication (MFA) information when you call AssumeRole. This is
// useful for cross-account scenarios to ensure that the user that assumes the role
// has been authenticated with an Amazon Web Services MFA device. In that scenario,
// the trust policy of the role being assumed includes a condition that tests for
// MFA authentication. If the caller does not include valid MFA information, the
// request to assume the role is denied. The condition in a trust policy that tests
// for MFA authentication might look like the following example. "Condition":
// {"Bool": {"aws:MultiFactorAuthPresent": true}} For more information, see
// Configuring MFA-Protected API Access
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/MFAProtectedAPI.html) in the
// IAM User Guide guide. To use MFA with AssumeRole, you pass values for the
// SerialNumber and TokenCode parameters. The SerialNumber value identifies the
// user's hardware or virtual MFA device. The TokenCode is the time-based one-time
// password (TOTP) that the MFA device produces.
func (c *Client) AssumeRole(ctx context.Context, params *AssumeRoleInput, optFns ...func(*Options)) (*AssumeRoleOutput, error) {
	if params == nil {
		params = &AssumeRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssumeRole", params, optFns, c.addOperationAssumeRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssumeRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssumeRoleInput struct {

	// The Amazon Resource Name (ARN) of the role to assume.
	//
	// This member is required.
	RoleArn *string

	// An identifier for the assumed role session. Use the role session name to
	// uniquely identify a session when the same role is assumed by different
	// principals or for different reasons. In cross-account scenarios, the role
	// session name is visible to, and can be logged by the account that owns the role.
	// The role session name is also used in the ARN of the assumed role principal.
	// This means that subsequent cross-account API requests that use the temporary
	// security credentials will expose the role session name to the external account
	// in their CloudTrail logs. The regex used to validate this parameter is a string
	// of characters consisting of upper- and lower-case alphanumeric characters with
	// no spaces. You can also include underscores or any of the following characters:
	// =,.@-
	//
	// This member is required.
	RoleSessionName *string

	// The duration, in seconds, of the role session. The value specified can range
	// from 900 seconds (15 minutes) up to the maximum session duration set for the
	// role. The maximum session duration setting can have a value from 1 hour to 12
	// hours. If you specify a value higher than this setting or the administrator
	// setting (whichever is lower), the operation fails. For example, if you specify a
	// session duration of 12 hours, but your administrator set the maximum session
	// duration to 6 hours, your operation fails. Role chaining limits your Amazon Web
	// Services CLI or Amazon Web Services API role session to a maximum of one hour.
	// When you use the AssumeRole API operation to assume a role, you can specify the
	// duration of your role session with the DurationSeconds parameter. You can
	// specify a parameter value of up to 43200 seconds (12 hours), depending on the
	// maximum session duration setting for your role. However, if you assume a role
	// using role chaining and provide a DurationSeconds parameter value greater than
	// one hour, the operation fails. To learn how to view the maximum value for your
	// role, see View the Maximum Session Duration Setting for a Role
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
	// in the IAM User Guide. By default, the value is set to 3600 seconds. The
	// DurationSeconds parameter is separate from the duration of a console session
	// that you might request using the returned credentials. The request to the
	// federation endpoint for a console sign-in token takes a SessionDuration
	// parameter that specifies the maximum length of the console session. For more
	// information, see Creating a URL that Enables Federated Users to Access the
	// Amazon Web Services Management Console
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
	// in the IAM User Guide.
	DurationSeconds *int32

	// A unique identifier that might be required when you assume a role in another
	// account. If the administrator of the account to which the role belongs provided
	// you with an external ID, then provide that value in the ExternalId parameter.
	// This value can be any string, such as a passphrase or account number. A
	// cross-account role is usually set up to trust everyone in an account. Therefore,
	// the administrator of the trusting account might send an external ID to the
	// administrator of the trusted account. That way, only someone with the ID can
	// assume the role, rather than everyone in the account. For more information about
	// the external ID, see How to Use an External ID When Granting Access to Your
	// Amazon Web Services Resources to a Third Party
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html)
	// in the IAM User Guide. The regex used to validate this parameter is a string of
	// characters consisting of upper- and lower-case alphanumeric characters with no
	// spaces. You can also include underscores or any of the following characters:
	// =,.@:/-
	ExternalId *string

	// An IAM policy in JSON format that you want to use as an inline session policy.
	// This parameter is optional. Passing policies to this operation returns new
	// temporary credentials. The resulting session's permissions are the intersection
	// of the role's identity-based policy and the session policies. You can use the
	// role's temporary credentials in subsequent Amazon Web Services API calls to
	// access resources in the account that owns the role. You cannot use session
	// policies to grant more permissions than those allowed by the identity-based
	// policy of the role that is being assumed. For more information, see Session
	// Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide. The plaintext that you use for both inline and managed
	// session policies can't exceed 2,048 characters. The JSON policy characters can
	// be any ASCII character from the space character to the end of the valid
	// character list (\u0020 through \u00FF). It can also include the tab (\u0009),
	// linefeed (\u000A), and carriage return (\u000D) characters. An Amazon Web
	// Services conversion compresses the passed session policies and session tags into
	// a packed binary format that has a separate limit. Your request can fail for this
	// limit even if your plaintext meets the other requirements. The PackedPolicySize
	// response element indicates by percentage how close the policies and tags for
	// your request are to the upper size limit.
	Policy *string

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want to
	// use as managed session policies. The policies must exist in the same account as
	// the role. This parameter is optional. You can provide up to 10 managed policy
	// ARNs. However, the plaintext that you use for both inline and managed session
	// policies can't exceed 2,048 characters. For more information about ARNs, see
	// Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference. An Amazon Web Services conversion
	// compresses the passed session policies and session tags into a packed binary
	// format that has a separate limit. Your request can fail for this limit even if
	// your plaintext meets the other requirements. The PackedPolicySize response
	// element indicates by percentage how close the policies and tags for your request
	// are to the upper size limit. Passing policies to this operation returns new
	// temporary credentials. The resulting session's permissions are the intersection
	// of the role's identity-based policy and the session policies. You can use the
	// role's temporary credentials in subsequent Amazon Web Services API calls to
	// access resources in the account that owns the role. You cannot use session
	// policies to grant more permissions than those allowed by the identity-based
	// policy of the role that is being assumed. For more information, see Session
	// Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	PolicyArns []types.PolicyDescriptorType

	// The identification number of the MFA device that is associated with the user who
	// is making the AssumeRole call. Specify this value if the trust policy of the
	// role being assumed includes a condition that requires MFA authentication. The
	// value is either the serial number for a hardware device (such as GAHT12345678)
	// or an Amazon Resource Name (ARN) for a virtual device (such as
	// arn:aws:iam::123456789012:mfa/user). The regex used to validate this parameter
	// is a string of characters consisting of upper- and lower-case alphanumeric
	// characters with no spaces. You can also include underscores or any of the
	// following characters: =,.@-
	SerialNumber *string

	// The source identity specified by the principal that is calling the AssumeRole
	// operation. You can require users to specify a source identity when they assume a
	// role. You do this by using the sts:SourceIdentity condition key in a role trust
	// policy. You can use source identity information in CloudTrail logs to determine
	// who took actions with a role. You can use the aws:SourceIdentity condition key
	// to further control access to Amazon Web Services resources based on the value of
	// source identity. For more information about using source identity, see Monitor
	// and control actions taken with assumed roles
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_monitor.html)
	// in the IAM User Guide. The regex used to validate this parameter is a string of
	// characters consisting of upper- and lower-case alphanumeric characters with no
	// spaces. You can also include underscores or any of the following characters:
	// =,.@-. You cannot use a value that begins with the text aws:. This prefix is
	// reserved for Amazon Web Services internal use.
	SourceIdentity *string

	// A list of session tags that you want to pass. Each session tag consists of a key
	// name and an associated value. For more information about session tags, see
	// Tagging Amazon Web Services STS Sessions
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html) in the
	// IAM User Guide. This parameter is optional. You can pass up to 50 session tags.
	// The plaintext session tag keys can’t exceed 128 characters, and the values can’t
	// exceed 256 characters. For these and additional limits, see IAM and STS
	// Character Limits
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length)
	// in the IAM User Guide. An Amazon Web Services conversion compresses the passed
	// session policies and session tags into a packed binary format that has a
	// separate limit. Your request can fail for this limit even if your plaintext
	// meets the other requirements. The PackedPolicySize response element indicates by
	// percentage how close the policies and tags for your request are to the upper
	// size limit. You can pass a session tag with the same key as a tag that is
	// already attached to the role. When you do, session tags override a role tag with
	// the same key. Tag key–value pairs are not case sensitive, but case is preserved.
	// This means that you cannot have separate Department and department tag keys.
	// Assume that the role has the Department=Marketing tag and you pass the
	// department=engineering session tag. Department and department are not saved as
	// separate tags, and the session tag passed in the request takes precedence over
	// the role tag. Additionally, if you used temporary credentials to perform this
	// operation, the new session inherits any transitive session tags from the calling
	// session. If you pass a session tag with the same key as an inherited tag, the
	// operation fails. To view the inherited tags for a session, see the CloudTrail
	// logs. For more information, see Viewing Session Tags in CloudTrail
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_ctlogs)
	// in the IAM User Guide.
	Tags []types.Tag

	// The value provided by the MFA device, if the trust policy of the role being
	// assumed requires MFA. (In other words, if the policy includes a condition that
	// tests for MFA). If the role being assumed requires MFA and if the TokenCode
	// value is missing or expired, the AssumeRole call returns an "access denied"
	// error. The format for this parameter, as described by its regex pattern, is a
	// sequence of six numeric digits.
	TokenCode *string

	// A list of keys for session tags that you want to set as transitive. If you set a
	// tag key as transitive, the corresponding key and value passes to subsequent
	// sessions in a role chain. For more information, see Chaining Roles with Session
	// Tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining)
	// in the IAM User Guide. This parameter is optional. When you set session tags as
	// transitive, the session policy and session tags packed binary limit is not
	// affected. If you choose not to specify a transitive tag key, then no tags are
	// passed from this session to any subsequent sessions.
	TransitiveTagKeys []string

	noSmithyDocumentSerde
}

// Contains the response to a successful AssumeRole request, including temporary
// Amazon Web Services credentials that can be used to make Amazon Web Services
// requests.
type AssumeRoleOutput struct {

	// The Amazon Resource Name (ARN) and the assumed role ID, which are identifiers
	// that you can use to refer to the resulting temporary security credentials. For
	// example, you can reference these credentials as a principal in a resource-based
	// policy by using the ARN or assumed role ID. The ARN and ID include the
	// RoleSessionName that you specified when you called AssumeRole.
	AssumedRoleUser *types.AssumedRoleUser

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token. The size of the security token
	// that STS API operations return is not fixed. We strongly recommend that you make
	// no assumptions about the maximum size.
	Credentials *types.Credentials

	// A percentage value that indicates the packed size of the session policies and
	// session tags combined passed in the request. The request fails if the packed
	// size is greater than 100 percent, which means the policies and tags exceeded the
	// allowed space.
	PackedPolicySize *int32

	// The source identity specified by the principal that is calling the AssumeRole
	// operation. You can require users to specify a source identity when they assume a
	// role. You do this by using the sts:SourceIdentity condition key in a role trust
	// policy. You can use source identity information in CloudTrail logs to determine
	// who took actions with a role. You can use the aws:SourceIdentity condition key
	// to further control access to Amazon Web Services resources based on the value of
	// source identity. For more information about using source identity, see Monitor
	// and control actions taken with assumed roles
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_monitor.html)
	// in the IAM User Guide. The regex used to validate this parameter is a string of
	// characters consisting of upper- and lower-case alphanumeric characters with no
	// spaces. You can also include underscores or any of the following characters:
	// =,.@-
	SourceIdentity *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssumeRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAssumeRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAssumeRole{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssumeRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssumeRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssumeRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sts",
		OperationName: "AssumeRole",
	}
}
