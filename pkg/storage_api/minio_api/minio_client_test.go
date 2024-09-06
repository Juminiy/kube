package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

const (
	endpoint        = "192.168.31.110:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	sessionToken    = ""
	secure          = false
)

var (
	testMinioClient, _ = New(
		endpoint,
		accessKeyID,
		secretAccessKey,
		sessionToken,
		secure,
	)
)

func TestClient_UpdateBucketQuota(t *testing.T) {
	util.SilentHandleError("update quota error",
		testMinioClient.UpdateBucketQuota(&BucketConfig{
			Quota:      util.Ti * 5,
			BucketName: "bin",
		}))

}
func TestClient_MakeBucket(t *testing.T) {
	util.SilentHandleError("create bucket error",
		testMinioClient.MakeBucket(&BucketConfig{
			Quota:            util.Gi * 30,
			BusinessUserName: "chisato",
		}))
}

func TestClient_RemoveBucket(t *testing.T) {
	util.SilentHandleError("remove bucket error",
		testMinioClient.RemoveBucket(&BucketConfig{
			BucketName: "s3fs-mount-bucket-chisato",
		}),
	)
}

func TestClient_BucketWorkFlow(t *testing.T) {

}
