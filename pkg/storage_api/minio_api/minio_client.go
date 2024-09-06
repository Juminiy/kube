package minio_api

import (
	"context"
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/log_api/zaplog"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	AccessKeyIDMaxLen     = 50
	SecretAccessKeyMaxLen = 128
)

type (
	Client struct {
		Endpoint string
		miniocred.Value
		Secure bool

		mc  *minio.Client
		ma  *madmin.AdminClient
		ctx context.Context
	}
)

func New(
	endpoint string,
	accessKeyID string,
	secretAccessKey string,
	sessionToken string,
	secure bool,
) (*Client, error) {
	mOpts := &minio.Options{
		Creds:  miniocred.NewStaticV4(accessKeyID, secretAccessKey, sessionToken),
		Secure: secure,
	}
	mc, err := minio.New(endpoint, mOpts)
	if err != nil {
		stdlog.ErrorF("minio client error: %s", err.Error())
		return nil, err
	}

	ma, err := madmin.NewWithOptions(endpoint, &madmin.Options{Creds: mOpts.Creds, Secure: mOpts.Secure})
	if err != nil {
		stdlog.ErrorF("minio admin client error: %s", err.Error())
		return nil, err
	}

	return &Client{
		mc:  mc,
		ma:  ma,
		ctx: context.TODO(),
	}, nil
}

// MakeBucket
// 1. make bucket
// 2. set bucket quota: size(B)
// 3. set bucket policy: readwrite
func (c *Client) MakeBucket(bucket *BucketConfig) error {
	bucket.BucketName = bucket.Name()
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
		(len(config.BusinessUserName) == 0 && len(config.BucketName) == 0) ||
		config.Quota <= 0 {
		return errors.New("bucket config error")
	}
	if len(config.BucketName) == 0 {
		config.BucketName = config.Name()
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

	return c.mc.SetBucketPolicy(
		c.ctx,
		config.BucketName,
		policy,
	)
}

func (c *Client) CreateAccessKey() (*miniocred.Value, error) {
	cred := NewCred(randAccessKeyID(), randSecretAccessKey())
	return cred, c.CreateIDPUser(cred)
}

func (c *Client) DeleteAccessKey(accessKeyID string) error {
	return c.DeleteIDPUser(accessKeyID)
}

func (c *Client) CreateIDPUser(cred *miniocred.Value) error {
	return c.ma.AddUser(
		c.ctx,
		randAccessKeyID(),
		randSecretAccessKey(),
	)
}

func (c *Client) DeleteIDPUser(accessKeyID string) error {
	return c.ma.RemoveUser(
		c.ctx,
		accessKeyID,
	)
}

func (c *Client) CreateAccessPolicy(config *PolicyConfig) error {
	policy, err := config.IBAPAccessKeyWithOneBucketObjectCRUDPolicy()
	if err != nil {
		return err
	}

	resp, err := c.ma.AttachPolicy(
		c.ctx,
		madmin.PolicyAssociationReq{
			Policies: []string{policy},
			User:     config.Cred.AccessKeyID,
			Group:    config.GroupName,
		},
	)
	zaplog.Info(resp)
	return err
}

func (c *Client) DeleteAccessPolicy(config *PolicyConfig) error {
	//c.ma.AttachPolicy()
	//c.ma.DetachPolicy()
	//
	//c.ma.AssignPolicy()
	//c.ma.DeletePolicy()
	//
	//c.ma.AddCannedPolicy()
	//c.ma.RemoveCannedPolicy()
	//

	return nil
}
