package s3_api

import "kube/pkg/util"

// gen from:
// https://min.io/docs/minio/linux/administration/identity-access-management/policy-based-access-control.html

const (
	Version = "2012-10-17"
	Allow   = "Allow"
	Deny    = "Deny"
)

// assigning to users or groups:
// consoleAdmin
const (
	S3All    = "s3:*"
	AdminAll = "admin:*"
)

// readonly
const (
	S3GetBucketLocation = "s3:GetBucketLocation"
	S3GetObject         = "s3:GetObject"
)

// readwrite
const (
	S3RW = "s3:*"
	S3Rw
)

// diagnostics
const (
	AdminServerTrace      = "admin:ServerTrace"
	AdminProfiling        = "admin:Profiling"
	AdminConsoleLog       = "admin:ConsoleLog"
	AdminServerInfo       = "admin:ServerInfo"
	AdminTopLocksInfo     = "admin:TopLocksInfo"
	AdminOBDInfo          = "admin:OBDInfo"
	AdminBandwidthMonitor = "admin:BandwidthMonitor"
	AdminPrometheus       = "admin:Prometheus"
)

// writeonly
const (
	S3PutObject = "s3:PutObject"
)

// S3:Bucket
const (
	S3CreateBucket      = "s3:CreateBucket"
	S3DeleteBucket      = "s3:DeleteBucket"
	S3ForceDeleteBucket = "s3:ForceDeleteBucket"
	//S3GetBucketLocation   = "s3:GetBucketLocation"
	S3ListAllMyBuckets = "s3:ListAllMyBuckets"
)

// S3:Object
const (
	S3DeleteObject = "s3:DeleteObject"
	//S3GetObject           = "s3:GetObject"
	S3ListBucket = "s3:ListBucket"
	//S3PutObject           = "s3:PutObject"
	S3PutObjectTagging    = "s3:PutObjectTagging"
	S3GetObjectTagging    = "s3:GetObjectTagging"
	S3DeleteObjectTagging = "s3:DeleteObjectTagging"
)

// S3:BucketConfig
const (
	S3GetBucketPolicy    = "s3:GetBucketPolicy"
	S3PutBucketPolicy    = "s3:PutBucketPolicy"
	S3DeleteBucketPolicy = "s3:DeleteBucketPolicy"
	S3GetBucketTagging   = "s3:GetBucketTagging"
	S3PutBucketTagging   = "s3:PutBucketTagging"
)

const (
	ResourceAll      = matchAny
	ResourceS3Prefix = "arn:aws:s3:::"
)

// WARNING: not api expose, only test for internal use
func AdminPolicy(policyName ...string) string {
	p := Policy{
		Version: Version,
		StatementList: StatementList{
			&AdminStatement{
				SASRStatement{
					Sid:       getSid(policyName...),
					Effect:    Allow,
					Action:    S3All,
					Resource:  ResourceAll,
					Condition: nil,
				},
			},
		},
	}
	return p.String()
}

func AccessKeyWithBucketRWPolicy(bucketName string, policyName ...string) string {
	p := Policy{
		Version: Version,
		StatementList: StatementList{
			&AccessKeyWithOneBucketRWStatement{
				MASRStatement{
					Sid:    getSid(policyName...),
					Effect: Allow,
					Action: []string{
						S3DeleteObject,
						S3GetObject,
						S3ListBucket,
						S3PutObject,
					},
					Resource: getBucketResource(bucketName),
				},
			},
		},
	}
	return p.String()
}

func BucketRWPolicy(bucketName string, policyName ...string) string {
	p := Policy{
		Version: Version,
		StatementList: StatementList{
			&BucketStatement{
				SAMRStatement{
					Sid:    getSid(policyName...),
					Effect: Allow,
					Action: S3All,
					Resource: []string{
						getBucketResource(bucketName),
						getBucketDirResource(bucketName),
					},
				},
			},
		},
	}
	return p.String()
}

func getSid(policyName ...string) string {
	pName := ""
	if len(policyName) > 0 {
		pName = policyName[0]
	}
	return pName
}

func getBucketResource(bucketName string) string {
	return util.StringConcat(ResourceS3Prefix, bucketName)
}

func getBucketDirResource(bucketName string) string {
	return util.StringConcat(ResourceS3Prefix, bucketName, dir, ResourceAll)
}
