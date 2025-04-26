package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed
func TestPolicyConfig_IBAPAccessKeyWithOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "juminiy-x",
			ID:   util.U64toa(8),
		},
		BucketName: "s3fs-mount-bucket-juminiy-x",
	}
	t.Log(pc.IBAPAccessKeyWithOneBucketObjectCRUDPolicy())
}

// +passed
func TestPolicyConfig_RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "juminiy-x",
			ID:   util.I64toa(8),
		},
		BucketName: "s3fs-mount-bucket-juminiy-x",
	}
	t.Log(pc.RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy())
}
