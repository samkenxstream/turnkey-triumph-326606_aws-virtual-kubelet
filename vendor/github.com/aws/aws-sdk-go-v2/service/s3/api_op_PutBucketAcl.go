// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the permissions on an existing bucket using access control lists (ACL). For
// more information, see Using ACLs
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html). To set
// the ACL of a bucket, you must have WRITE_ACP permission. You can use one of the
// following two ways to set a bucket's permissions:
//
// * Specify the ACL in the
// request body
//
// * Specify permissions using request headers
//
// You cannot specify
// access permission using both the body and the request headers. Depending on your
// application needs, you may choose to set the ACL on a bucket using either the
// request body or the headers. For example, if you have an existing application
// that updates a bucket ACL using the request body, then you can continue to use
// that approach. Access Permissions You can set access permissions using one of
// the following methods:
//
// * Specify a canned ACL with the x-amz-acl request
// header. Amazon S3 supports a set of predefined ACLs, known as canned ACLs. Each
// canned ACL has a predefined set of grantees and permissions. Specify the canned
// ACL name as the value of x-amz-acl. If you use this header, you cannot use other
// access control-specific headers in your request. For more information, see
// Canned ACL
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
// *
// Specify access permissions explicitly with the x-amz-grant-read,
// x-amz-grant-read-acp, x-amz-grant-write-acp, and x-amz-grant-full-control
// headers. When using these headers, you specify explicit access permissions and
// grantees (AWS accounts or Amazon S3 groups) who will receive the permission. If
// you use these ACL-specific headers, you cannot use the x-amz-acl header to set a
// canned ACL. These parameters map to the set of permissions that Amazon S3
// supports in an ACL. For more information, see Access Control List (ACL) Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html). You specify
// each grantee as a type=value pair, where the type is one of the following:
//
// * id
// – if the value specified is the canonical user ID of an AWS account
//
// * uri – if
// you are granting permissions to a predefined group
//
// * emailAddress – if the
// value specified is the email address of an AWS account Using email addresses to
// specify a grantee is only supported in the following AWS Regions:
//
// * US East (N.
// Virginia)
//
// * US West (N. California)
//
// * US West (Oregon)
//
// * Asia Pacific
// (Singapore)
//
// * Asia Pacific (Sydney)
//
// * Asia Pacific (Tokyo)
//
// * Europe
// (Ireland)
//
// * South America (São Paulo)
//
// For a list of all the Amazon S3
// supported Regions and endpoints, see Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) in the AWS
// General Reference.
//
// For example, the following x-amz-grant-write header grants
// create, overwrite, and delete objects permission to LogDelivery group predefined
// by Amazon S3 and two AWS accounts identified by their email addresses.
// x-amz-grant-write: uri="http://acs.amazonaws.com/groups/s3/LogDelivery",
// id="111122223333", id="555566667777"
//
// You can use either a canned ACL or specify
// access permissions explicitly. You cannot do both. Grantee Values You can
// specify the person (grantee) to whom you're assigning access rights (using
// request elements) in the following ways:
//
// * By the person's ID:
// <>ID<><>GranteesEmail<>  DisplayName is optional and ignored in the request
//
// *
// By URI: <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// * By
// Email address: <>Grantees@email.com<>lt;/Grantee> The grantee is resolved to the
// CanonicalUser and, in a response to a GET Object acl request, appears as the
// CanonicalUser. Using email addresses to specify a grantee is only supported in
// the following AWS Regions:
//
// * US East (N. Virginia)
//
// * US West (N.
// California)
//
// * US West (Oregon)
//
// * Asia Pacific (Singapore)
//
// * Asia Pacific
// (Sydney)
//
// * Asia Pacific (Tokyo)
//
// * Europe (Ireland)
//
// * South America (São
// Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see
// Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) in the AWS
// General Reference.
//
// Related Resources
//
// * CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
//
// *
// DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//
// *
// GetObjectAcl
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html)
func (c *Client) PutBucketAcl(ctx context.Context, params *PutBucketAclInput, optFns ...func(*Options)) (*PutBucketAclOutput, error) {
	if params == nil {
		params = &PutBucketAclInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketAcl", params, optFns, c.addOperationPutBucketAclMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAclInput struct {

	// The bucket to which to apply the ACL.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL

	// Contains the elements that set the ACL permissions for an object per grantee.
	AccessControlPolicy *types.AccessControlPolicy

	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to RFC 1864.
	// (http://www.ietf.org/rfc/rfc1864.txt) For requests made using the AWS Command
	// Line Interface (CLI) or AWS SDKs, this field is calculated automatically.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket. For the bucket and object
	// owners of existing objects, also allows deletions and overwrites of those
	// objects.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string
}

type PutBucketAclOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutBucketAclMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAcl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAcl{}, middleware.After)
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
	if err = addOpPutBucketAclValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAcl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketAclUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketAcl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketAcl",
	}
}

// getPutBucketAclBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketAclBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketAclInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketAclUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketAclBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		TargetS3ObjectLambda:    false,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
