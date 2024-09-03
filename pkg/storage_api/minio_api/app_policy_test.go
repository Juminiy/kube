package minio_api

import (
	"strconv"
	"testing"

	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

func TestPolicyConfig_IBAPAccessKeyWithOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "chisato-x",
			ID:   strconv.FormatUint(8, 10),
		},
		BucketName: "s3fs-mount-bucket-chisato-x",
	}
	stdlog.Info(pc.IBAPAccessKeyWithOneBucketObjectCRUDPolicy())
}

func TestPolicyConfig_RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "chisato-x",
			ID:   strconv.FormatUint(8, 10),
		},
		BucketName: "s3fs-mount-bucket-chisato-x",
	}
	stdlog.Info(pc.RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy())
}
