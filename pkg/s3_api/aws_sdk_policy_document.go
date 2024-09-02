package s3_api

import (
	awssdkv2iam "github.com/aws/aws-sdk-go-v2/service/iam"
	awssdkiam "github.com/aws/aws-sdk-go/service/iam"
)

// AWSIAMClient
// convinced that PolicyDocument describe:
// - gotype: string and *string
// - format: JSON encode
type AWSIAMClient struct {
	Policy   awssdkiam.CreatePolicyInput
	PolicyV2 awssdkv2iam.CreatePolicyInput
}
