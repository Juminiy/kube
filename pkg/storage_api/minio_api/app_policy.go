package minio_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/s3_api"
	s3apiv2 "github.com/Juminiy/kube/pkg/storage_api/s3_api/v2"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/google/uuid"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

type (
	PolicyConfig struct {
		BusinessUser BusinessUser

		// Minio:UserName equals to Minio:AccessKeyID
		// +optional
		Cred miniocred.Value
		// +optional
		GroupName string

		BucketName string
	}

	BusinessUser struct {
		Name string
		ID   string
	}
)

const (
	IAMIBAP = "IBAP"
	IAMRBAP = "RBAP"

	IdentifyAdmin     = "Admin"
	IdentifyAccessKey = "AccessKey"
)

func (c *PolicyConfig) IBAPAccessKeyWithOneBucketObjectCRUDPolicy() (string, error) {
	policy := s3apiv2.IBAPolicy{
		Version: Version,
		Statement: []s3apiv2.Statement{
			s3apiv2.Statement{
				Sid: makeStatementSid(
					c.BusinessUser.Name,
					IAMIBAP,
					IdentifyAccessKey,
					"WithOneBucketObjectCRUD",
				),
				Effect:       Allow,
				Principal:    nil,
				NotPrincipal: nil,
				Action: []string{
					s3_api.S3ListBucket,
					s3_api.S3DeleteObject,
					s3_api.S3GetObject,
					s3_api.S3PutObject,
				},
				NotAction: nil,
				Resource: []string{
					s3_api.GetBucketAnyResource(c.BucketName),
				},
				NotResource: nil,
			},
		},
	}
	return policy.String()
}

func (c *PolicyConfig) RBAPBucketWithAdminAllWithAccessKeyOneBucketObjectCRUDPolicy() (string, error) {
	policy := s3apiv2.RBAPolicy{
		Version: Version,
		Id:      makePolicyId(c.BusinessUser.ID),
		Statement: []s3apiv2.Statement{
			c.RBAPBucketWithAdminAllStatement(),
			c.RBAPBucketWithAccessKeyOneBucketObjectCRUDStatement(),
		},
	}
	return policy.String()
}

func (c *PolicyConfig) RBAPBucketWithAdminAllStatement() s3apiv2.Statement {
	return s3apiv2.Statement{
		Sid: makeStatementSid(
			c.BusinessUser.Name,
			IAMRBAP,
			IdentifyAdmin,
		),
		Effect:       Allow,
		Principal:    map[string]any{},
		NotPrincipal: nil,
		Action:       s3apiv2.ActionAll,
		NotAction:    nil,
		Resource: []string{
			s3_api.GetBucketResource(c.BucketName),
			s3_api.GetBucketAnyResource(c.BucketName),
		},
		NotResource: nil,
	}
}

func (c *PolicyConfig) RBAPBucketWithAccessKeyOneBucketObjectCRUDStatement() s3apiv2.Statement {
	return s3apiv2.Statement{
		Sid: makeStatementSid(
			c.BusinessUser.Name,
			IAMRBAP,
			IdentifyAccessKey,
		),
		Effect:       Allow,
		Principal:    map[string]any{},
		NotPrincipal: nil,
		Action: []string{
			s3_api.S3ListBucket,
			s3_api.S3DeleteObject,
			s3_api.S3GetObject,
			s3_api.S3PutObject,
		},
		NotAction: nil,
		Resource: []string{
			s3_api.GetBucketAnyResource(c.BucketName),
		},
		NotResource: nil,
	}
}

func makePolicyId(bUID string) string {
	return util.StringJoin(
		"-",
		bUID,
		uuid.NewString(),
	)
}

func makeStatementSid(s ...string) string {
	return util.StringJoin(
		"-",
		s...,
	)
}
