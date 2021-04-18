// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Provides information about your AWS account.
type AccountInfo struct {

	// The identifier of the AWS account that is assigned to the user.
	AccountId *string

	// The display name of the AWS account that is assigned to the user.
	AccountName *string

	// The email address of the AWS account that is assigned to the user.
	EmailAddress *string
}

// Provides information about the role credentials that are assigned to the user.
type RoleCredentials struct {

	// The identifier used for the temporary security credentials. For more
	// information, see Using Temporary Security Credentials to Request Access to AWS
	// Resources
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	AccessKeyId *string

	// The date on which temporary security credentials expire.
	Expiration int64

	// The key that is used to sign the request. For more information, see Using
	// Temporary Security Credentials to Request Access to AWS Resources
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	SecretAccessKey *string

	// The token used for temporary credentials. For more information, see Using
	// Temporary Security Credentials to Request Access to AWS Resources
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	SessionToken *string
}

// Provides information about the role that is assigned to the user.
type RoleInfo struct {

	// The identifier of the AWS account assigned to the user.
	AccountId *string

	// The friendly name of the role that is assigned to the user.
	RoleName *string
}
