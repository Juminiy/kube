package minio_api

import (
	"strconv"
	"testing"
)

// +passed
func TestPolicyConfig_IBAPAccessKeyWithOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "chisato-x",
			ID:   strconv.FormatUint(8, 10),
		},
		BucketName: "s3fs-mount-bucket-chisato-x",
	}
	t.Log(pc.IBAPAccessKeyWithOneBucketObjectCRUDPolicy())
}

// +passed
func TestPolicyConfig_RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "chisato-x",
			ID:   strconv.FormatUint(8, 10),
		},
		BucketName: "s3fs-mount-bucket-chisato-x",
	}
	t.Log(pc.RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy())
}
