package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestClient_UpdateBucketQuota(t *testing.T) {
	util.SilentHandleError("update quota error",
		New().UpdateBucketQuota(&BucketConfig{
			SizeB: util.Ti * 5,
			BName: "bin",
		}))

}
func TestClient_MakeBucket(t *testing.T) {
	util.SilentHandleError("create bucket error",
		New().MakeBucket(&BucketConfig{
			SizeB: util.Gi * 30,
			UName: "chisato",
		}))
}

func TestClient_RemoveBucket(t *testing.T) {
	util.SilentHandleError("remove bucket error",
		New().RemoveBucket(&BucketConfig{
			BName: "s3fs-mount-bucket-chisato",
		}),
	)
}

func TestClient_GetBucketAccess(t *testing.T) {
}
