package s3_api

import (
	awsv2iam "github.com/aws/aws-sdk-go-v2/service/iam"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

// AWSIAMClient
// convinced that PolicyDocument describe:
// - gotype: string and *string
// - format: JSON encode
type AWSIAMClient struct {
	Policy   awsiam.CreatePolicyInput
	PolicyV2 awsv2iam.CreatePolicyInput
}
