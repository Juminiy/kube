package minio_api

import (
	"testing"

	"github.com/Juminiy/kube/pkg/util"
)

const (
	endpoint        = "192.168.31.110:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	sessionToken    = ""
	secure          = false
)

var (
	testMinioClient, testMinioClientError = New(
		endpoint,
		accessKeyID,
		secretAccessKey,
		sessionToken,
		secure,
	)
)

func TestClient_UpdateBucketQuota(t *testing.T) {
	util.SilentPanicError(testMinioClientError)
	util.SilentHandleError("update quota error",
		testMinioClient.UpdateBucketQuota(&BucketConfig{
			Quota:      util.Ti * 5,
			BucketName: "bin",
		}))

}
func TestClient_MakeBucket(t *testing.T) {
	util.SilentPanicError(testMinioClientError)
	util.SilentHandleError("create bucket error",
		testMinioClient.MakeBucket(&BucketConfig{
			Quota: util.Gi * 30,
			BusinessUser: BusinessUser{
				Name: "chisato",
			},
		}))
}

func TestClient_RemoveBucket(t *testing.T) {
	util.SilentPanicError(testMinioClientError)
	util.SilentHandleError("remove bucket error",
		testMinioClient.RemoveBucket(&BucketConfig{
			BucketName: "s3fs-mount-bucket-chisato",
		}),
	)
}

func TestClient_BucketWorkFlow(t *testing.T) {

}
