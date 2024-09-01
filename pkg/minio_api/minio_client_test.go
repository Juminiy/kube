package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestClient_UpdateBucketQuota(t *testing.T) {
	util.SilentHandleError("update quota error",
		New().UpdateBucketQuota(&BucketConfig{
			Quota:      util.Ti * 5,
			BucketName: "bin",
		}))

}
func TestClient_MakeBucket(t *testing.T) {
	util.SilentHandleError("create bucket error",
		New().MakeBucket(&BucketConfig{
			Quota:            util.Gi * 30,
			BusinessUserName: "chisato",
		}))
}

func TestClient_RemoveBucket(t *testing.T) {
	util.SilentHandleError("remove bucket error",
		New().RemoveBucket(&BucketConfig{
			BucketName: "s3fs-mount-bucket-chisato",
		}),
	)
}

func TestClient_BucketWorkFlow(t *testing.T) {

}
