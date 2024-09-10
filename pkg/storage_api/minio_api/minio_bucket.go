package minio_api

import (
	"errors"
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

// +self define
func (c *BucketConfig) setDefaultBucketName() {
	c.BucketName = strings.Join([]string{
		"s3fs",
		"mount",
		"bucket",
		c.BusinessUser.Name,
	}, "-")
}

// MakeBucket
// 1. make a new bucket
// 2. set bucket quota: size(B)
// 3. set bucket access policy
func (c *Client) MakeBucket(bucketConfig *BucketConfig) error {
	if len(bucketConfig.BucketName) == 0 {
		bucketConfig.setDefaultBucketName()
	}

	err := c.mc.MakeBucket(
		c.ctx,
		bucketConfig.BucketName,
		minio.MakeBucketOptions{},
	)
	if err != nil {
		return err
	}

	err = c.setBucketQuota(bucketConfig)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveBucket(bucketConfig *BucketConfig) error {
	return c.mc.RemoveBucket(
		c.ctx,
		bucketConfig.BucketName,
	)
}

func (c *Client) UpdateBucketQuota(bucketConfig *BucketConfig) error {
	return c.setBucketQuota(bucketConfig)
}

func (c *Client) setBucketQuota(bucketConfig *BucketConfig) error {
	if bucketConfig == nil ||
		(len(bucketConfig.BusinessUser.Name) == 0 && len(bucketConfig.BucketName) == 0) ||
		bucketConfig.Quota <= 0 {
		return errors.New("bucket config error")
	}
	if len(bucketConfig.BucketName) == 0 {
		bucketConfig.setDefaultBucketName()
	}

	return c.ma.SetBucketQuota(
		c.ctx,
		bucketConfig.BucketName,
		&madmin.BucketQuota{
			Quota: bucketConfig.Quota, // Deprecated, but set it
			Size:  bucketConfig.Quota,
			Type:  madmin.HardQuota,
		},
	)
}

func (c *Client) SetBucketPolicy(policyConfig *PolicyConfig) error {
	policy, err := policyConfig.RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy()
	if err != nil {
		return err
	}

	return c.mc.SetBucketPolicy(
		c.ctx,
		policyConfig.BucketName,
		policy,
	)
}
