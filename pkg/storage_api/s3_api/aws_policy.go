package s3_api

import "github.com/Juminiy/kube/pkg/util"

// referred from:
// https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html#amazons3-actions-as-permissions

// Version
const (
	// Version
	// current version of the policy language
	// always include and set to Version "2012-10-17"
	Version = "2012-10-17"

	// EarlierVersion
	// earlier version of the policy language
	// for example: earlier version will not identify ${aws:username} as a variable
	// instead, earlier version identify it is a text-string
	EarlierVersion = "2008-10-17"
)

// Effect
const (
	Allow = "Allow"
	Deny  = "Deny"
)

// Action
// S3:All
const (
	ActionAll = matchAny
	S3All     = "s3:*"
)

// S3:Bucket
const (
	S3CreateBucket      = "s3:CreateBucket"
	S3DeleteBucket      = "s3:DeleteBucket"
	S3ForceDeleteBucket = "s3:ForceDeleteBucket"
	S3GetBucketLocation = "s3:GetBucketLocation"
	S3ListAllMyBuckets  = "s3:ListAllMyBuckets"
	S3ListBucket        = "s3:ListBucket"
)

// S3:Object
const (
	S3DeleteObject        = "s3:DeleteObject"
	S3GetObject           = "s3:GetObject"
	S3PutObject           = "s3:PutObject"
	S3PutObjectTagging    = "s3:PutObjectTagging"
	S3GetObjectTagging    = "s3:GetObjectTagging"
	S3DeleteObjectTagging = "s3:DeleteObjectTagging"
)

// S3:Bucket Configuration
const (
	S3GetBucketPolicy    = "s3:GetBucketPolicy"
	S3PutBucketPolicy    = "s3:PutBucketPolicy"
	S3DeleteBucketPolicy = "s3:DeleteBucketPolicy"
	S3GetBucketTagging   = "s3:GetBucketTagging"
	S3PutBucketTagging   = "s3:PutBucketTagging"
)

// S3:Multipart Upload
const (
	S3AbortMultipartUpload       = "s3:AbortMultipartUpload"
	S3ListMultipartUploadParts   = "s3:ListMultipartUploadParts"
	S3ListBucketMultipartUploads = "s3:ListBucketMultipartUploads"
)

// S3:Versioning and Retention
const (
	S3PutBucketVersioning              = "s3:PutBucketVersioning"
	S3GetBucketVersioning              = "s3:GetBucketVersioning"
	S3DeleteObjectVersion              = "s3:DeleteObjectVersion"
	S3ListBucketVersions               = "s3:ListBucketVersions"
	S3PutObjectVersionTagging          = "s3:PutObjectVersionTagging"
	S3GetObjectVersionTagging          = "s3:GetObjectVersionTagging"
	S3DeleteObjectVersionTagging       = "s3:DeleteObjectVersionTagging"
	S3GetObjectVersion                 = "s3:GetObjectVersion"
	S3BypassGovernanceRetention        = "s3:BypassGovernanceRetention"
	S3PutObjectRetention               = "s3:PutObjectRetention"
	S3GetObjectRetention               = "s3:GetObjectRetention"
	S3GetObjectLegalHold               = "s3:GetObjectLegalHold"
	S3PutObjectLegalHold               = "s3:PutObjectLegalHold"
	S3GetBucketObjectLockConfiguration = "s3:GetBucketObjectLockConfiguration"
	S3PutBucketObjectLockConfiguration = "s3:PutBucketObjectLockConfiguration"
)

// S3:Bucket Notifications
const (
	S3GetBucketNotification    = "s3:GetBucketNotification"
	S3PutBucketNotification    = "s3:PutBucketNotification"
	S3ListenNotification       = "s3:ListenNotification"
	S3ListenBucketNotification = "s3:ListenBucketNotification"
)

// S3:Object Lifecycle Management
const (
	S3PutLifecycleConfiguration = "s3:PutLifecycleConfiguration"
	S3GetLifecycleConfiguration = "s3:GetLifecycleConfiguration"
)

// S3:Object Encryption
const (
	S3GetEncryptionConfiguration = "s3:GetEncryptionConfiguration"
	S3PutEncryptionConfiguration = "s3:PutEncryptionConfiguration"
)

// S3:Bucket Replication
const (
	S3GetReplicationConfiguration    = "s3:GetReplicationConfiguration"
	S3PutReplicationConfiguration    = "s3:PutReplicationConfiguration"
	S3ReplicateObject                = "s3:ReplicateObject"
	S3ReplicateDelete                = "s3:ReplicateDelete"
	S3ReplicateTags                  = "s3:ReplicateTags"
	S3GetObjectVersionForReplication = "s3:GetObjectVersionForReplication"
)

// S3:Condition Keys
const (
	AWSReferer         = "aws:Referer"
	AWSSourceIp        = "aws:SourceIp"
	AWSUserAgent       = "aws:UserAgent"
	AWSSecureTransport = "aws:SecureTransport"
	AWSCurrentTime     = "aws:CurrentTime"
	AWSEpochTime       = "aws:EpochTime"
	AWSPrincipalType   = "aws:PrincipalType"
	AWSUserid          = "aws:userid"
	AWSUsername        = "aws:username"
	XAmzContentSha256  = "x-amz-content-sha256"
	S3signatureAge     = "s3:signatureAge"
)

// Resource
const (
	ResourceAll         = matchAny
	ResourceARNS3All    = "arn:aws:s3:::*"
	ResourceARNS3Prefix = "arn:aws:s3:::"
)

// Principal
const (
	PrincipalAll          = matchAny
	PrincipalARNIAMPrefix = "arn:aws:iam::"
)

/*
 * Func
 */

func GetSid(policyName ...string) string {
	pName := ""
	if len(policyName) > 0 {
		pName = policyName[0]
	}
	return pName
}

func GetBucketResource(bucketName string) string {
	return util.StringConcat(ResourceARNS3Prefix, bucketName)
}

func GetBucketAnyResource(bucketName string) string {
	return util.StringConcat(ResourceARNS3Prefix, bucketName, dirSlash, ResourceAll)
}

func GetPrincipalAccountRoot(accountId string) string {
	return util.StringConcat(PrincipalARNIAMPrefix, accountId, ":root")
}

func GetPrincipalAccountUser(accountId, userName string) string {
	return util.StringConcat(PrincipalARNIAMPrefix, accountId, ":user/", userName)
}
