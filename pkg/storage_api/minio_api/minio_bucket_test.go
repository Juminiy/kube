package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"github.com/brianvoe/gofakeit/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/s3utils"
	"strconv"
	"testing"
)

// +passed
func TestClient_MakeBucket(t *testing.T) {
	util.SilentFatalf("create bucket error",
		_cli.MakeBucket(&BucketConfig{
			Quota: util.Gi * 30,
			BusinessUser: BusinessUser{
				Name: "chisato",
			},
		}))
}

// +passed
func TestClient_UpdateBucketQuota(t *testing.T) {
	util.SilentFatalf("update quota error",
		_cli.UpdateBucketQuota(&BucketConfig{
			Quota:      util.Gi * 60,
			BucketName: "s3fs-mount-bucket-chisato",
		}))

}

// +passed
func TestClient_RemoveBucket(t *testing.T) {
	util.SilentFatalf("remove bucket error",
		_cli.RemoveBucket("s3fs-mount-bucket-chisato"),
	)
}

// +passed
func TestClient_SetBucketPolicy(t *testing.T) {
	util.SilentFatalf("create bucket policy error", _cli.SetBucketPolicy(
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
	util.SilentFatalf("create bucket error", _cli.MakeBucket(&bucketConfig))
	t.Log(bucketConfig)

	policyConfig := PolicyConfig{
		BusinessUser: businessUser,
		Cred: miniocred.Value{
			AccessKeyID: "uUDC29bGJj3v15K33rAmM1urgRk6c924eov0IrF6PZz3BnHj24",
		},
		BucketName: bucketConfig.BucketName,
	}

	//create bucket policy
	util.SilentFatalf("create bucket policy error", _cli.SetBucketPolicy(&policyConfig))

}

// +passed
func TestClient_ListBucket(t *testing.T) {
	buckets, err := _cli.WithPage(util.NewPageConfig(0, 128)).ListBucket()
	util.Must(err)

	t.Log(buckets)
}

func TestClient_BatchRemoveBucket(t *testing.T) {
	bucketName := make([]string, 0, util.MagicSliceCap)
	for range 32 {
		namei := random.Integer(8)
		util.Must(s3utils.CheckValidBucketNameStrict(namei))
		bucketName = append(bucketName, namei)
	}

	for i, namei := range bucketName {
		util.Must(_cli.MakeBucket(&BucketConfig{
			BusinessUser: BusinessUser{ID: "test_minio" + strconv.Itoa(i), Name: gofakeit.Username()},
			Quota:        100 * util.Ki,
			BucketName:   namei,
		}))
	}
	util.TestLongHorizontalLine(t)

	t.Log(_cli.BatchRemoveBucket(util.Slice2Map[[]string, map[string]struct{}](bucketName)))
}
