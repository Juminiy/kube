package minio_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_go"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	"strings"
)

type BucketConfig struct {
	// username to make bucket
	// +required BusinessUser.ID
	BusinessUser BusinessUser

	// restrict bucket Quota size: /Byte
	// +required
	Quota uint64

	// bucket name to get, delete, update
	// +optional
	BucketName string
}

var (
	errBucketConfigNull  = errors.New("bucket config is null")
	errBucketConfigName  = errors.New("bucket config bucket name error")
	errBucketConfigQuota = errors.New("bucket config bucket quota error")
)

// +self define
// v1.1.1->v1.1.2 update name rule
func (c *BucketConfig) setDefaultBucketName() {
	c.BucketName = strings.Join([]string{
		"s3fs",
		"mount",
		"bucket",
		c.BusinessUser.ID,
	}, "-")
}

// MakeBucket
// 1. make a new bucket
// 2. set bucket quota: size(B)
// 3. set bucket access policy
func (c *Client) MakeBucket(bucketConfig *BucketConfig) error {
	if len(bucketConfig.BusinessUser.ID) == 0 && len(bucketConfig.BucketName) == 0 {
		return errBucketConfigName
	}
	if bucketConfig.Quota <= 0 {
		return errBucketConfigQuota
	}
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

var _forcedRemoveBucket = minio.RemoveBucketOptions{ForceDelete: true}

// RemoveBucket
// Remove Bucket itself and Bucket Objects, Versions, Markers
// All objects (including all object versions and delete markers)
func (c *Client) RemoveBucket(bucketName string) error {
	return c.mc.RemoveBucketWithOptions(
		c.ctx,
		bucketName,
		_forcedRemoveBucket,
	)
}

func (c *Client) UpdateBucketQuota(bucketConfig *BucketConfig) error {
	return c.setBucketQuota(bucketConfig)
}

func (c *Client) setBucketQuota(bucketConfig *BucketConfig) error {
	if len(bucketConfig.BusinessUser.ID) == 0 && len(bucketConfig.BucketName) == 0 {
		return errBucketConfigName
	}
	if bucketConfig.Quota <= 0 {
		return errBucketConfigQuota
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

func (c *Client) ListBucket() ([]minio.BucketInfo, error) {
	bucketsInfo, err := c.mc.ListBuckets(c.ctx)
	if err != nil {
		return nil, err
	}

	limitSize := c.page.SizeIntValue()
	if limitSize >= len(bucketsInfo) {
		return bucketsInfo, nil
	}
	return bucketsInfo[:limitSize], nil
}

func (c *Client) BatchRemoveBucket(bucketNames map[string]struct{}) error {
	fns := make([]util.Func, 0, len(bucketNames))
	for bucketName := range bucketNames {
		fns = append(fns, func() error {
			return c.RemoveBucket(bucketName)
		})
	}
	return safe_go.DryRun(fns...)
}
