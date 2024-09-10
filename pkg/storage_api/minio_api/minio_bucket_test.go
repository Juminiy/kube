package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
	"strconv"
	"testing"
)

// +passed
func TestClient_MakeBucket(t *testing.T) {
	util.SilentHandleError("create bucket error",
		testMinioClient.MakeBucket(&BucketConfig{
			Quota: util.Gi * 30,
			BusinessUser: BusinessUser{
				Name: "chisato",
			},
		}))
}

// +passed
func TestClient_UpdateBucketQuota(t *testing.T) {
	util.SilentHandleError("update quota error",
		testMinioClient.UpdateBucketQuota(&BucketConfig{
			Quota:      util.Gi * 60,
			BucketName: "s3fs-mount-bucket-chisato",
		}))

}

// +passed
func TestClient_RemoveBucket(t *testing.T) {
	util.SilentHandleError("remove bucket error",
		testMinioClient.RemoveBucket(&BucketConfig{
			BucketName: "s3fs-mount-bucket-chisato",
		}),
	)
}

// +passed
func TestClient_SetBucketPolicy(t *testing.T) {
	util.SilentHandleError("create bucket policy error", testMinioClient.SetBucketPolicy(
		&PolicyConfig{
			BusinessUser: BusinessUser{
				Name: "chisato",
				ID:   strconv.Itoa(11),
			},
			Cred: miniocred.Value{
				AccessKeyID: "uUDC29bGJj3v15K33rAmM1urgRk6c924eov0IrF6PZz3BnHj24",
			},
			BucketName: "s3fs-mount-bucket-chisato",
		},
	))
}

// FINISH: create bucket by businessUser{ID, Name} and set bucket quota
// FINISH: create policy attach to bucket
// +passed
func TestClient_BucketWorkflow(t *testing.T) {
	businessUser := BusinessUser{
		Name: "chisatox0129",
		ID:   strconv.Itoa(33),
	}

	// create bucket, set quota
	bucketConfig := BucketConfig{
		Quota:        util.Gi * 114514,
		BusinessUser: businessUser,
	}
	util.SilentHandleError("create bucket error", testMinioClient.MakeBucket(&bucketConfig))
	stdlog.Info(bucketConfig)

	policyConfig := PolicyConfig{
		BusinessUser: businessUser,
		Cred: miniocred.Value{
			AccessKeyID: "uUDC29bGJj3v15K33rAmM1urgRk6c924eov0IrF6PZz3BnHj24",
		},
		BucketName: bucketConfig.BucketName,
	}

	//create bucket policy
	util.SilentHandleError("create bucket policy error", testMinioClient.SetBucketPolicy(&policyConfig))

}
