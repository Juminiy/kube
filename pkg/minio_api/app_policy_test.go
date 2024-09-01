package minio_api

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPolicyConfig_IBAPAccessKeyWithOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "chisato-x",
			ID:   strconv.FormatUint(8, 10),
		},
		BucketName: "s3fs-mount-bucket-chisato-x",
	}
	fmt.Println(pc.IBAPAccessKeyWithOneBucketObjectCRUDPolicy())
}

func TestPolicyConfig_RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy(t *testing.T) {
	pc := PolicyConfig{
		BusinessUser: BusinessUser{
			Name: "chisato-x",
			ID:   strconv.FormatUint(8, 10),
		},
		BucketName: "s3fs-mount-bucket-chisato-x",
	}
	fmt.Println(pc.RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy())
}
