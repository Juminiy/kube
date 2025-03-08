package minio_api

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cast"
)

type Req struct {
	// +required
	UserID uint64
	// +required
	UserName string
	// +required
	BucketQuotaByte uint64
	// +optional
	BucketName string
}

type Resp struct {
	Req
	CredValue      miniocred.Value
	CredPolicyName string
}

// AtomicWorkflow
// functionality:
//  1. create a bucket, set bucket quota by byte{BucketQuotaByte} with optional bucket name{BucketName}
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
// so call to with go func() { resp, err := minioClient.AtomicWorkflow() } or use channel is a good practice
func (c *Client) AtomicWorkflow(req Req) (resp Resp, err error) {

	resp.Req = req
	businessUser := BusinessUser{
		ID:   util.U64toa(req.UserID),
		Name: req.UserName,
	}

	bucketConfig := BucketConfig{
		BusinessUser: businessUser,
		Quota:        req.BucketQuotaByte,
		BucketName:   req.BucketName,
	}

	credValue := miniocred.Value{}
	policyConfig := PolicyConfig{}

	progErrFn := func(step int, desc string, err error) {
		stdlog.ErrorF("create minio atomic workflow progress %d/4, step %d desc: %s error: %s", step-1, step, desc, err.Error())
	}

	//1.
	err = c.MakeBucket(&bucketConfig)
	if err != nil {
		progErrFn(1, "make bucket", err)
		goto makeBucketRollback
	}
	resp.BucketName = bucketConfig.BucketName

	//2.
	credValue, err = c.CreateAccessKey()
	if err != nil {
		progErrFn(2, "create accessKey", err)
		goto createAccessKeyRollback
	}
	resp.CredValue = credValue

	policyConfig = PolicyConfig{
		BusinessUser: businessUser,
		Cred:         credValue,
		BucketName:   bucketConfig.BucketName,
	}

	//3.
	err = c.SetBucketPolicy(&policyConfig)
	if err != nil {
		progErrFn(3, "set bucket policy", err)
		goto createAccessKeyRollback
	}

	//4.
	err = c.SetAccessPolicy(&policyConfig)
	if err != nil {
		progErrFn(4, "create and set accessKey policy", err)
		goto setAccessPolicyRollback
	}
	resp.CredPolicyName = policyConfig.PolicyName

	return

	//5.
	// rollback and return
setAccessPolicyRollback:
	c.DeleteAccessPolicy(&policyConfig)
createAccessKeyRollback:
	c.DeleteAccessKey(credValue.AccessKeyID)
makeBucketRollback:
	c.RemoveBucket(bucketConfig.BucketName)

	return
}

func (c *Client) AtomicDeleteFlow(resp Resp) error {
	errH := util.NewErrHandle()
	progErrFn := func(step int, desc string, err error) {
		errH.HasStr(fmt.Sprintf("delete minio atomic workflow progress %d/3, step %d desc: %s error: %s", step-1, step, desc, err.Error()))
	}

	err := c.DeleteAccessPolicy(&PolicyConfig{
		BusinessUser: BusinessUser{
			ID:   cast.ToString(resp.UserID),
			Name: resp.UserName,
		},
		Cred:       resp.CredValue,
		BucketName: resp.BucketName,
		PolicyName: resp.CredPolicyName,
	})
	if err != nil {
		progErrFn(1, "delete access policy", err)
	}

	err = c.DeleteAccessKey(resp.CredValue.AccessKeyID)
	if err != nil {
		progErrFn(2, "delete access key", err)
	}

	err = c.RemoveBucket(resp.BucketName)
	if err != nil {
		progErrFn(3, "remove bucket", err)
	}
	return errH.All(", ")
}

func (c *Client) ClientDo(fc func(mc *minio.Client) error) error {
	return fc(c.mc)
}

func (c *Client) AdminClientDo(fc func(ma *madmin.AdminClient) error) error {
	return fc(c.ma)
}
