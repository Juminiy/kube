package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"github.com/minio/madmin-go/v3"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

func NewCred(id, secret string) miniocred.Value {
	return miniocred.Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
	}
}

func (c *Client) CreateAccessKey() (miniocred.Value, error) {
	cred := NewCred(randAccessKeyID(), randSecretAccessKey())
	return cred, c.AddIAMUser(&cred)
}

func (c *Client) DeleteAccessKey(accessKeyID string) error {
	return c.RemoveIAMUser(accessKeyID)
}

func (c *Client) AddIAMUser(cred *miniocred.Value) error {
	return c.ma.AddUser(
		c.ctx,
		cred.AccessKeyID,
		cred.SecretAccessKey,
	)
}

func (c *Client) RemoveIAMUser(accessKeyID string) error {
	return c.ma.RemoveUser(
		c.ctx,
		accessKeyID,
	)
}

// SetAccessPolicy
// 1. create an IBA AccessKey Policy(IAM Policy) from business user information
// 2. and attach the created policy to the business user's AccessKey
func (c *Client) SetAccessPolicy(config *PolicyConfig) error {
	policy, err := config.IBAPAccessKeyWithOneBucketObjectCRUDPolicy()
	if err != nil {
		return err
	}
	config.setPolicyName()

	err = c.ma.AddCannedPolicy(
		c.ctx,
		config.PolicyName,
		util.String2BytesNoCopy(policy),
	)
	if err != nil {
		return err
	}

	_, err = c.ma.AttachPolicy(
		c.ctx,
		madmin.PolicyAssociationReq{
			Policies: []string{config.PolicyName},
			User:     config.Cred.AccessKeyID,
			Group:    config.GroupName,
		},
	)
	//stdlog.Debug(resp)

	return err
}

// DeleteAccessPolicy
// 1. detach the PolicyJSONString from an exists user
// 1. delete the PolicyJSONString from Minio Server(Minio IAM)
func (c *Client) DeleteAccessPolicy(config *PolicyConfig) error {
	_, err := c.ma.DetachPolicy(
		c.ctx,
		madmin.PolicyAssociationReq{
			Policies: []string{config.PolicyName},
			User:     config.Cred.AccessKeyID,
			Group:    config.GroupName,
		},
	)
	if err != nil {
		return err
	}

	err = c.ma.RemoveCannedPolicy(
		c.ctx,
		config.PolicyName,
	)
	//stdlog.Debug(resp)

	return err
}

func randAccessKeyID() string {
	return random.IDString(AccessKeyIDMaxLen)
}

func randSecretAccessKey() string {
	return random.PasswordString(SecretAccessKeyMaxLen)
}
