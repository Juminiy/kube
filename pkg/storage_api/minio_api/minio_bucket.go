package minio_api

import "strings"

type BucketConfig struct {
	// restrict bucket Quota size: /Byte
	Quota uint64
	// username to make bucket
	BusinessUserName string
	// bucket name to get,delete,update
	BucketName string
}

func (c *BucketConfig) Name() string {
	return strings.Join([]string{
		"s3fs",
		"mount",
		"bucket",
		c.BusinessUserName,
	}, "-")
}
