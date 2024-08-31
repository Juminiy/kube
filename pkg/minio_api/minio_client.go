package minio_api

import (
	"context"
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"strings"
)

const (
	endpoint        = "192.168.31.110:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	sessionToken    = ""
	secure          = false
)

type (
	Client struct {
		Endpoint string
		ID       string // AccessKeyID
		Secret   string // SecretAccessKey
		Token    string // SessionToken
		Secure   bool

		mc  *minio.Client
		ma  *madmin.AdminClient
		ctx context.Context
	}
	BucketConfig struct {
		// restrict bucket size /Byte
		SizeB uint64
		// username to make bucket
		UName string
		// bucket name to get,delete,update
		BName string
	}
)

func New() *Client {
	mOpts := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, sessionToken),
		Secure: secure,
	}
	mc, err := minio.New(endpoint, mOpts)
	util.SilentHandleError("minio client error: ", err)

	ma, err := madmin.NewWithOptions(endpoint, &madmin.Options{Creds: mOpts.Creds, Secure: mOpts.Secure})
	util.SilentHandleError("minio admin client error: ", err)

	return &Client{mc: mc, ma: ma, ctx: context.TODO()}
}

// MakeBucket
// 1. make bucket
// 2. set bucket quota: size(B)
// 3. set bucket policy: readwrite
func (c *Client) MakeBucket(bucket *BucketConfig) error {
	bucket.BName = bucket.Name()
	err := c.mc.MakeBucket(
		c.ctx,
		bucket.BName,
		minio.MakeBucketOptions{},
	)
	if err != nil {
		return err
	}

	//err := c.mc.SetBucketPolicy(
	//	c.ctx,
	//	bucket.BName,
	//)
	//if err != nil {
	//	return err
	//}

	return c.setBucketQuota(bucket)
}

func (c *Client) RemoveBucket(bucket *BucketConfig) error {
	return c.mc.RemoveBucket(
		c.ctx,
		bucket.BName,
	)
}

func (c *Client) UpdateBucketQuota(bucket *BucketConfig) error {
	return c.setBucketQuota(bucket)
}

func (c *Client) setBucketQuota(bucket *BucketConfig) error {
	if bucket == nil ||
		(len(bucket.UName) == 0 && len(bucket.BName) == 0) ||
		bucket.SizeB <= 0 {
		return errors.New("bucket config error")
	}
	if len(bucket.BName) == 0 {
		bucket.BName = bucket.Name()
	}
	return c.ma.SetBucketQuota(
		c.ctx,
		bucket.BName,
		&madmin.BucketQuota{
			Quota: bucket.SizeB,
			Size:  bucket.SizeB,
			Type:  madmin.HardQuota,
		},
	)
}

func (c *Client) createKey() error {
	//return c.ma.CreateKey(
	//	c.ctx,
	//	s3_api.AccessKeyWithBucketRWPolicy(),
	//)
	return nil
}

func (c *BucketConfig) Name() string {
	return strings.Join([]string{
		"s3fs",
		"mount",
		"bucket",
		c.UName,
	}, "-")
}

func createKeyRandomPolicy() (string, string) {
	return "", ""
}

func createKeyNameEncryptPolicy() (string, string) {
	return "", ""
}
