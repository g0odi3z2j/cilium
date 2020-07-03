// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssumeRoleWithSAMLInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, of the role session. Your role session lasts for
	// the duration that you specify for the DurationSeconds parameter, or until
	// the time specified in the SAML authentication response's SessionNotOnOrAfter
	// value, whichever is shorter. You can provide a DurationSeconds value from
	// 900 seconds (15 minutes) up to the maximum session duration setting for the
	// role. This setting can have a value from 1 hour to 12 hours. If you specify
	// a value higher than this setting, the operation fails. For example, if you
	// specify a session duration of 12 hours, but your administrator set the maximum
	// session duration to 6 hours, your operation fails. To learn how to view the
	// maximum value for your role, see View the Maximum Session Duration Setting
	// for a Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
	// in the IAM User Guide.
	//
	// By default, the value is set to 3600 seconds.
	//
	// The DurationSeconds parameter is separate from the duration of a console
	// session that you might request using the returned credentials. The request
	// to the federation endpoint for a console sign-in token takes a SessionDuration
	// parameter that specifies the maximum length of the console session. For more
	// information, see Creating a URL that Enables Federated Users to Access the
	// AWS Management Console (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
	// in the IAM User Guide.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// An IAM policy in JSON format that you want to use as an inline session policy.
	//
	// This parameter is optional. Passing policies to this operation returns new
	// temporary credentials. The resulting session's permissions are the intersection
	// of the role's identity-based policy and the session policies. You can use
	// the role's temporary credentials in subsequent AWS API calls to access resources
	// in the account that owns the role. You cannot use session policies to grant
	// more permissions than those allowed by the identity-based policy of the role
	// that is being assumed. For more information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	//
	// The plain text that you use for both inline and managed session policies
	// can't exceed 2,048 characters. The JSON policy characters can be any ASCII
	// character from the space character to the end of the valid character list
	// (\u0020 through \u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// An AWS conversion compresses the passed session policies and session tags
	// into a packed binary format that has a separate limit. Your request can fail
	// for this limit even if your plain text meets the other requirements. The
	// PackedPolicySize response element indicates by percentage how close the policies
	// and tags for your request are to the upper size limit.
	Policy *string `min:"1" type:"string"`

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want
	// to use as managed session policies. The policies must exist in the same account
	// as the role.
	//
	// This parameter is optional. You can provide up to 10 managed policy ARNs.
	// However, the plain text that you use for both inline and managed session
	// policies can't exceed 2,048 characters. For more information about ARNs,
	// see Amazon Resource Names (ARNs) and AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// An AWS conversion compresses the passed session policies and session tags
	// into a packed binary format that has a separate limit. Your request can fail
	// for this limit even if your plain text meets the other requirements. The
	// PackedPolicySize response element indicates by percentage how close the policies
	// and tags for your request are to the upper size limit.
	//
	// Passing policies to this operation returns new temporary credentials. The
	// resulting session's permissions are the intersection of the role's identity-based
	// policy and the session policies. You can use the role's temporary credentials
	// in subsequent AWS API calls to access resources in the account that owns
	// the role. You cannot use session policies to grant more permissions than
	// those allowed by the identity-based policy of the role that is being assumed.
	// For more information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	PolicyArns []PolicyDescriptorType `type:"list"`

	// The Amazon Resource Name (ARN) of the SAML provider in IAM that describes
	// the IdP.
	//
	// PrincipalArn is a required field
	PrincipalArn *string `min:"20" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the role that the caller is assuming.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// The base-64 encoded SAML authentication response provided by the IdP.
	//
	// For more information, see Configuring a Relying Party and Adding Claims (https://docs.aws.amazon.com/IAM/latest/UserGuide/create-role-saml-IdP-tasks.html)
	// in the IAM User Guide.
	//
	// SAMLAssertion is a required field
	SAMLAssertion *string `min:"4" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AssumeRoleWithSAMLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssumeRoleWithSAMLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssumeRoleWithSAMLInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}

	if s.PrincipalArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrincipalArn"))
	}
	if s.PrincipalArn != nil && len(*s.PrincipalArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PrincipalArn", 20))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.SAMLAssertion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SAMLAssertion"))
	}
	if s.SAMLAssertion != nil && len(*s.SAMLAssertion) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("SAMLAssertion", 4))
	}
	if s.PolicyArns != nil {
		for i, v := range s.PolicyArns {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PolicyArns", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful AssumeRoleWithSAML request, including
// temporary AWS credentials that can be used to make AWS requests.
type AssumeRoleWithSAMLOutput struct {
	_ struct{} `type:"structure"`

	// The identifiers for the temporary security credentials that the operation
	// returns.
	AssumedRoleUser *AssumedRoleUser `type:"structure"`

	// The value of the Recipient attribute of the SubjectConfirmationData element
	// of the SAML assertion.
	Audience *string `type:"string"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// The size of the security token that STS API operations return is not fixed.
	// We strongly recommend that you make no assumptions about the maximum size.
	Credentials *Credentials `type:"structure"`

	// The value of the Issuer element of the SAML assertion.
	Issuer *string `type:"string"`

	// A hash value based on the concatenation of the Issuer response value, the
	// AWS account ID, and the friendly name (the last part of the ARN) of the SAML
	// provider in IAM. The combination of NameQualifier and Subject can be used
	// to uniquely identify a federated user.
	//
	// The following pseudocode shows how the hash value is calculated:
	//
	// BASE64 ( SHA1 ( "https://example.com/saml" + "123456789012" + "/MySAMLIdP"
	// ) )
	NameQualifier *string `type:"string"`

	// A percentage value that indicates the packed size of the session policies
	// and session tags combined passed in the request. The request fails if the
	// packed size is greater than 100 percent, which means the policies and tags
	// exceeded the allowed space.
	PackedPolicySize *int64 `type:"integer"`

	// The value of the NameID element in the Subject element of the SAML assertion.
	Subject *string `type:"string"`

	// The format of the name ID, as defined by the Format attribute in the NameID
	// element of the SAML assertion. Typical examples of the format are transient
	// or persistent.
	//
	// If the format includes the prefix urn:oasis:names:tc:SAML:2.0:nameid-format,
	// that prefix is removed. For example, urn:oasis:names:tc:SAML:2.0:nameid-format:transient
	// is returned as transient. If the format includes any other prefix, the format
	// is returned with no modifications.
	SubjectType *string `type:"string"`
}

// String returns the string representation
func (s AssumeRoleWithSAMLOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssumeRoleWithSAML = "AssumeRoleWithSAML"

// AssumeRoleWithSAMLRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns a set of temporary security credentials for users who have been authenticated
// via a SAML authentication response. This operation provides a mechanism for
// tying an enterprise identity store or directory to role-based AWS access
// without user-specific credentials or configuration. For a comparison of AssumeRoleWithSAML
// with the other API operations that produce temporary credentials, see Requesting
// Temporary Security Credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS API operations (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// The temporary security credentials returned by this operation consist of
// an access key ID, a secret access key, and a security token. Applications
// can use these temporary security credentials to sign calls to AWS services.
//
// Session Duration
//
// By default, the temporary security credentials created by AssumeRoleWithSAML
// last for one hour. However, you can use the optional DurationSeconds parameter
// to specify the duration of your session. Your role session lasts for the
// duration that you specify, or until the time specified in the SAML authentication
// response's SessionNotOnOrAfter value, whichever is shorter. You can provide
// a DurationSeconds value from 900 seconds (15 minutes) up to the maximum session
// duration setting for the role. This setting can have a value from 1 hour
// to 12 hours. To learn how to view the maximum value for your role, see View
// the Maximum Session Duration Setting for a Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
// in the IAM User Guide. The maximum session duration limit applies when you
// use the AssumeRole* API operations or the assume-role* CLI commands. However
// the limit does not apply when you use those operations to create a console
// URL. For more information, see Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html)
// in the IAM User Guide.
//
// Permissions
//
// The temporary security credentials created by AssumeRoleWithSAML can be used
// to make API calls to any AWS service with the following exception: you cannot
// call the STS GetFederationToken or GetSessionToken API operations.
//
// (Optional) You can pass inline or managed session policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// to this operation. You can pass a single JSON policy document to use as an
// inline session policy. You can also specify up to 10 managed policies to
// use as managed session policies. The plain text that you use for both inline
// and managed session policies can't exceed 2,048 characters. Passing policies
// to this operation returns new temporary credentials. The resulting session's
// permissions are the intersection of the role's identity-based policy and
// the session policies. You can use the role's temporary credentials in subsequent
// AWS API calls to access resources in the account that owns the role. You
// cannot use session policies to grant more permissions than those allowed
// by the identity-based policy of the role that is being assumed. For more
// information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// in the IAM User Guide.
//
// Calling AssumeRoleWithSAML does not require the use of AWS security credentials.
// The identity of the caller is validated by using keys in the metadata document
// that is uploaded for the SAML provider entity for your identity provider.
//
// Calling AssumeRoleWithSAML can result in an entry in your AWS CloudTrail
// logs. The entry includes the value in the NameID element of the SAML assertion.
// We recommend that you use a NameIDType that is not associated with any personally
// identifiable information (PII). For example, you could instead use the persistent
// identifier (urn:oasis:names:tc:SAML:2.0:nameid-format:persistent).
//
// Tags
//
// (Optional) You can configure your IdP to pass attributes into your SAML assertion
// as session tags. Each session tag consists of a key name and an associated
// value. For more information about session tags, see Passing Session Tags
// in STS (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html)
// in the IAM User Guide.
//
// You can pass up to 50 session tags. The plain text session tag keys can’t
// exceed 128 characters and the values can’t exceed 256 characters. For these
// and additional limits, see IAM and STS Character Limits (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length)
// in the IAM User Guide.
//
// An AWS conversion compresses the passed session policies and session tags
// into a packed binary format that has a separate limit. Your request can fail
// for this limit even if your plain text meets the other requirements. The
// PackedPolicySize response element indicates by percentage how close the policies
// and tags for your request are to the upper size limit.
//
// You can pass a session tag with the same key as a tag that is attached to
// the role. When you do, session tags override the role's tags with the same
// key.
//
// An administrator must grant you the permissions necessary to pass session
// tags. The administrator can also create granular permissions to allow you
// to pass only specific session tags. For more information, see Tutorial: Using
// Tags for Attribute-Based Access Control (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html)
// in the IAM User Guide.
//
// You can set the session tags as transitive. Transitive tags persist during
// role chaining. For more information, see Chaining Roles with Session Tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining)
// in the IAM User Guide.
//
// SAML Configuration
//
// Before your application can call AssumeRoleWithSAML, you must configure your
// SAML identity provider (IdP) to issue the claims required by AWS. Additionally,
// you must use AWS Identity and Access Management (IAM) to create a SAML provider
// entity in your AWS account that represents your identity provider. You must
// also create an IAM role that specifies this SAML provider in its trust policy.
//
// For more information, see the following resources:
//
//    * About SAML 2.0-based Federation (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html)
//    in the IAM User Guide.
//
//    * Creating SAML Identity Providers (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml.html)
//    in the IAM User Guide.
//
//    * Configuring a Relying Party and Claims (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml_relying-party.html)
//    in the IAM User Guide.
//
//    * Creating a Role for SAML 2.0 Federation (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_saml.html)
//    in the IAM User Guide.
//
//    // Example sending a request using AssumeRoleWithSAMLRequest.
//    req := client.AssumeRoleWithSAMLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithSAML
func (c *Client) AssumeRoleWithSAMLRequest(input *AssumeRoleWithSAMLInput) AssumeRoleWithSAMLRequest {
	op := &aws.Operation{
		Name:       opAssumeRoleWithSAML,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssumeRoleWithSAMLInput{}
	}

	req := c.newRequest(op, input, &AssumeRoleWithSAMLOutput{})
	req.Config.Credentials = aws.AnonymousCredentials

	return AssumeRoleWithSAMLRequest{Request: req, Input: input, Copy: c.AssumeRoleWithSAMLRequest}
}

// AssumeRoleWithSAMLRequest is the request type for the
// AssumeRoleWithSAML API operation.
type AssumeRoleWithSAMLRequest struct {
	*aws.Request
	Input *AssumeRoleWithSAMLInput
	Copy  func(*AssumeRoleWithSAMLInput) AssumeRoleWithSAMLRequest
}

// Send marshals and sends the AssumeRoleWithSAML API request.
func (r AssumeRoleWithSAMLRequest) Send(ctx context.Context) (*AssumeRoleWithSAMLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssumeRoleWithSAMLResponse{
		AssumeRoleWithSAMLOutput: r.Request.Data.(*AssumeRoleWithSAMLOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssumeRoleWithSAMLResponse is the response type for the
// AssumeRoleWithSAML API operation.
type AssumeRoleWithSAMLResponse struct {
	*AssumeRoleWithSAMLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssumeRoleWithSAML request.
func (r *AssumeRoleWithSAMLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
