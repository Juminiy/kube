package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"strconv"
	"testing"
)

// TODO: create policy attach to bucket
// TODO: create policy attach to accessKey
func TestClient_BusinessWorkflow(t *testing.T) {
	util.SilentPanicError(testMinioClientError)

	businessUser := BusinessUser{
		Name: "chisatox0129",
		ID:   strconv.Itoa(33),
	}

	// create bucket, set quota
	bucketConfig := BucketConfig{
		Quota:        util.Gi * 114514,
		BusinessUser: businessUser,
	}
	util.SilentPanicError(testMinioClient.MakeBucket(&bucketConfig))

	// create access key
	//userCred, err := testMinioClient.CreateAccessKey()
	//util.SilentPanicError(err)

	//policyConfig := PolicyConfig{
	//	BusinessUser: businessUser,
	//	Cred:         userCred,
	//	BucketName:   bucketConfig.BucketName,
	//}
	// create access policy
	//util.SilentPanicError(testMinioClient.CreateAccessPolicy(&policyConfig))

	// create bucket policy
	//util.SilentPanicError(testMinioClient.CreateBucketPolicy(&policyConfig))

}

func TestClient_RemoveBucket2(t *testing.T) {
	util.SilentPanicError(testMinioClientError)
	util.SilentHandleError("remove bucket error",
		testMinioClient.RemoveBucket(&BucketConfig{
			BucketName: "s3fs-mount-bucket-chisatox0129",
		}),
	)
}

func TestClient_DeleteAccessKey(t *testing.T) {
	util.SilentPanicError(testMinioClientError)
	util.SilentHandleError("delete access key error",
		testMinioClient.DeleteAccessKey(""),
	)
}
