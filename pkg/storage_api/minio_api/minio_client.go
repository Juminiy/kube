package minio_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	AccessKeyIDMaxLen     = 50
	SecretAccessKeyMaxLen = 128
)

type Client struct {
	Endpoint string
	miniocred.Value
	Secure bool

	mc  *minio.Client
	ma  *madmin.AdminClient
	ctx context.Context
}

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

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

type Req struct {
	UserID          uint64
	UserName        string
	BucketQuotaByte uint64
	BucketName      string
}

type Resp struct {
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

// AtomicWorkflow
// functionality:
//  1. create a bucket, set bucket quota by byte{BucketName}
//  2. create an iam credential{AccessKeyID, SecretAccessKey}
//  3. set bucket RBAPolicy, with info(2): Principal{AccessKeyID}, with info(1): Resource{BucketName, BucketName/*}, Action{CRUD}
//  4. iam user policy
//     (1). create an IAM IBAPolicy, with info(1): Resource{BucketName, BucketName/*}, with Action{CRUD}
//     (2). attach policy with info(4) to info(2)
//  5. return Result of all sensitive info(1,2,3,4,5)
//
// atomicity:
//  1. when any of 1,2,3,4,5 failed, rollback all, return error
//  2. when all of 1,2,3,4,5 success, return Resp
//
// synchronization
// 1. the func call may take some time and may be failure with nothing created in Minio Server
// so call to with go func() { resp, err := minioClient.AtomicWorkflow() } is a good practice
func (c *Client) AtomicWorkflow(req Req) (Resp, error) {

	var returnErr error

	businessUser := BusinessUser{
		ID:   util.U64toa(req.UserID),
		Name: req.UserName,
	}

	bucketConfig := BucketConfig{
		BusinessUser: businessUser,
		Quota:        req.BucketQuotaByte,
		BucketName:   req.BucketName,
	}

	//1.
	makeBucketErr := c.MakeBucket(&bucketConfig)
	if makeBucketErr != nil {
		returnErr = makeBucketErr
		//goto makeBucketRollback
	}

	//2.
	credValue, createAccessErr := c.CreateAccessKey()
	if createAccessErr != nil {
		returnErr = createAccessErr
		//goto createAccessKeyRollback
	}

	policyConfig := PolicyConfig{
		BusinessUser: businessUser,
		Cred:         credValue,
		BucketName:   bucketConfig.BucketName,
	}

	//3.
	setBucketPolicyErr := c.SetBucketPolicy(&policyConfig)
	if setBucketPolicyErr != nil {
		returnErr = setBucketPolicyErr
		//goto createAccessKeyRollback
	}

	//4.
	setAccessPolicyErr := c.SetAccessPolicy(&policyConfig)
	if setAccessPolicyErr != nil {
		returnErr = setAccessPolicyErr
		//goto setAccessPolicyRollback
	}

	// rollback
	//setAccessPolicyRollback:
	//	c.DeleteAccessPolicy(&policyConfig)
	//	return Resp{}, returnErr
	//createAccessKeyRollback:
	//	c.DeleteAccessKey(credValue.AccessKeyID)
	//makeBucketRollback:
	//	c.RemoveBucket(&bucketConfig)

	//5.
	return Resp{}, returnErr
}
