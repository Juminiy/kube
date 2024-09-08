package minio_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	"strings"
)

type BucketConfig struct {
	// username to make bucket
	BusinessUser BusinessUser

	// restrict bucket Quota size: /Byte
	Quota uint64

	// bucket name to get, delete, update
	BucketName string
}

func (c *BucketConfig) setDefaultBucketName() {
	c.BucketName = strings.Join([]string{
		"s3fs",
		"mount",
		"bucket",
		c.BusinessUser.Name,
	}, "-")
}

// MakeBucket
// 1. make bucket
// 2. set bucket quota: size(B)
// 3. set bucket policy: readwrite
func (c *Client) MakeBucket(bucket *BucketConfig) error {
	if len(bucket.BucketName) == 0 {
		bucket.setDefaultBucketName()
	}
	err := c.mc.MakeBucket(
		c.ctx,
		bucket.BucketName,
		minio.MakeBucketOptions{},
	)
	if err != nil {
		return err
	}

	return c.setBucketQuota(bucket)
}

func (c *Client) RemoveBucket(config *BucketConfig) error {
	return c.mc.RemoveBucket(
		c.ctx,
		config.BucketName,
	)
}

func (c *Client) UpdateBucketQuota(config *BucketConfig) error {
	return c.setBucketQuota(config)
}

func (c *Client) setBucketQuota(config *BucketConfig) error {
	if config == nil ||
		(len(config.BusinessUser.Name) == 0 && len(config.BucketName) == 0) ||
		config.Quota <= 0 {
		return errors.New("bucket config error")
	}
	if len(config.BucketName) == 0 {
		config.setDefaultBucketName()
	}

	return c.ma.SetBucketQuota(
		c.ctx,
		config.BucketName,
		&madmin.BucketQuota{
			Quota: config.Quota, // Deprecated, but set it
			Size:  config.Quota,
			Type:  madmin.HardQuota,
		},
	)
}

func (c *Client) CreateBucketPolicy(config *PolicyConfig) error {
	policy, err := config.RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy()
	if err != nil {
		return err
	}

	err = c.ma.AddCannedPolicy(
		c.ctx,
		config.GetPolicyName(),
		util.String2BytesNoCopy(policy),
	)
	if err != nil {
		return err
	}

	return c.mc.SetBucketPolicy(
		c.ctx,
		config.BucketName,
		policy,
	)
}
