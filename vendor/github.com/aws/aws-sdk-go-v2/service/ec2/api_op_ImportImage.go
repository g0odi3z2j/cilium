// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// To import your virtual machines (VMs) with a console-based experience, you can
// use the Import virtual machine images to Amazon Web Services template in the
// Migration Hub Orchestrator console (https://console.aws.amazon.com/migrationhub/orchestrator)
// . For more information, see the Migration Hub Orchestrator User Guide  (https://docs.aws.amazon.com/migrationhub-orchestrator/latest/userguide/import-vm-images.html)
// . Import single or multi-volume disk images or EBS snapshots into an Amazon
// Machine Image (AMI). Amazon Web Services VM Import/Export strongly recommends
// specifying a value for either the --license-type or --usage-operation parameter
// when you create a new VM Import task. This ensures your operating system is
// licensed appropriately and your billing is optimized. For more information, see
// Importing a VM as an image using VM Import/Export (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html)
// in the VM Import/Export User Guide.
func (c *Client) ImportImage(ctx context.Context, params *ImportImageInput, optFns ...func(*Options)) (*ImportImageOutput, error) {
	if params == nil {
		params = &ImportImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportImage", params, optFns, c.addOperationImportImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportImageInput struct {

	// The architecture of the virtual machine. Valid values: i386 | x86_64
	Architecture *string

	// The boot mode of the virtual machine. The uefi-preferred boot mode isn't
	// supported for importing images. For more information, see Boot modes (https://docs.aws.amazon.com/vm-import/latest/userguide/prerequisites.html#vmimport-boot-modes)
	// in the VM Import/Export User Guide.
	BootMode types.BootModeValues

	// The client-specific data.
	ClientData *types.ClientData

	// The token to enable idempotency for VM import requests.
	ClientToken *string

	// A description string for the import image task.
	Description *string

	// Information about the disk containers.
	DiskContainers []types.ImageDiskContainer

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Specifies whether the destination AMI of the imported image should be
	// encrypted. The default KMS key for EBS is used unless you specify a non-default
	// KMS key using KmsKeyId . For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool

	// The target hypervisor platform. Valid values: xen
	Hypervisor *string

	// An identifier for the symmetric KMS key to use when creating the encrypted AMI.
	// This parameter is only required if you want to use a non-default KMS key; if
	// this parameter is not specified, the default KMS key for EBS is used. If a
	// KmsKeyId is specified, the Encrypted flag must also be set. The KMS key
	// identifier may be provided in any of the following formats:
	//   - Key ID
	//   - Key alias
	//   - ARN using key ID. The ID ARN contains the arn:aws:kms namespace, followed by
	//   the Region of the key, the Amazon Web Services account ID of the key owner, the
	//   key namespace, and then the key ID. For example,
	//   arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//   - ARN using key alias. The alias ARN contains the arn:aws:kms namespace,
	//   followed by the Region of the key, the Amazon Web Services account ID of the key
	//   owner, the alias namespace, and then the key alias. For example,
	//   arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	// Amazon Web Services parses KmsKeyId asynchronously, meaning that the action you
	// call may appear to complete even though you provided an invalid identifier. This
	// action will eventually report failure. The specified KMS key must exist in the
	// Region that the AMI is being copied to. Amazon EBS does not support asymmetric
	// KMS keys.
	KmsKeyId *string

	// The ARNs of the license configurations.
	LicenseSpecifications []types.ImportImageLicenseConfigurationRequest

	// The license type to be used for the Amazon Machine Image (AMI) after importing.
	// Specify AWS to replace the source-system license with an Amazon Web Services
	// license or BYOL to retain the source-system license. Leaving this parameter
	// undefined is the same as choosing AWS when importing a Windows Server operating
	// system, and the same as choosing BYOL when importing a Windows client operating
	// system (such as Windows 10) or a Linux operating system. To use BYOL , you must
	// have existing licenses with rights to use these licenses in a third party cloud,
	// such as Amazon Web Services. For more information, see Prerequisites (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html#prerequisites-image)
	// in the VM Import/Export User Guide.
	LicenseType *string

	// The operating system of the virtual machine. If you import a VM that is
	// compatible with Unified Extensible Firmware Interface (UEFI) using an EBS
	// snapshot, you must specify a value for the platform. Valid values: Windows |
	// Linux
	Platform *string

	// The name of the role to use when not using the default role, 'vmimport'.
	RoleName *string

	// The tags to apply to the import image task during creation.
	TagSpecifications []types.TagSpecification

	// The usage operation value. For more information, see Licensing options (https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#prerequisites)
	// in the VM Import/Export User Guide.
	UsageOperation *string

	noSmithyDocumentSerde
}

type ImportImageOutput struct {

	// The architecture of the virtual machine.
	Architecture *string

	// A description of the import task.
	Description *string

	// Indicates whether the AMI is encrypted.
	Encrypted *bool

	// The target hypervisor of the import task.
	Hypervisor *string

	// The ID of the Amazon Machine Image (AMI) created by the import task.
	ImageId *string

	// The task ID of the import image task.
	ImportTaskId *string

	// The identifier for the symmetric KMS key that was used to create the encrypted
	// AMI.
	KmsKeyId *string

	// The ARNs of the license configurations.
	LicenseSpecifications []types.ImportImageLicenseConfigurationResponse

	// The license type of the virtual machine.
	LicenseType *string

	// The operating system of the virtual machine.
	Platform *string

	// The progress of the task.
	Progress *string

	// Information about the snapshots.
	SnapshotDetails []types.SnapshotDetail

	// A brief status of the task.
	Status *string

	// A detailed status message of the import task.
	StatusMessage *string

	// Any tags assigned to the import image task.
	Tags []types.Tag

	// The usage operation value.
	UsageOperation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpImportImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpImportImage{}, middleware.After)
	if err != nil {
		return err
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
	if err = addImportImageResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportImage(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opImportImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ImportImage",
	}
}

type opImportImageResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opImportImageResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opImportImageResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addImportImageResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opImportImageResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
