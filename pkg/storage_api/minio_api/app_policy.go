package minio_api

import (
	miniointernal "github.com/Juminiy/kube/pkg/storage_api/minio_api/internal"
	"github.com/Juminiy/kube/pkg/storage_api/s3_api"
	s3apiv2 "github.com/Juminiy/kube/pkg/storage_api/s3_api/v2"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"github.com/google/uuid"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

type (
	PolicyConfig struct {
		BusinessUser BusinessUser

		// Minio's UserName equals to Minio's AccessKeyID
		// +required
		Cred miniocred.Value

		// +optional
		GroupName string

		// +required
		BucketName string

		// when delete policy from a user, must provide the PolicyJSONString
		// +required
		// when create policy from a user, api will ignore the PolicyJSONString
		// +optional
		PolicyJSONString string

		// when delete policy from a user, must provide the PolicyName
		// +required
		// when create policy from a user, api will ignore the PolicyName
		// +optional
		PolicyName string
	}
)

const (
	IAMIBAP = "IBAP"
	IAMRBAP = "RBAP"

	IdentifyAdmin     = "Admin"
	IdentifyAccessKey = "AccessKey"
)

func (c *PolicyConfig) GetPolicyName() string {
	return util.StringJoin(
		"-",
		c.BusinessUser.ID,
		c.BusinessUser.Name,
		c.BucketName,
		random.IDString(len(c.BusinessUser.Name)),
	)
}

func (c *PolicyConfig) IBAPAccessKeyWithOneBucketObjectCRUDPolicy() (string, error) {
	policy := s3apiv2.IBAPolicy{
		Version: miniointernal.Version,
		Statement: []s3apiv2.Statement{
			s3apiv2.Statement{
				Sid: makeStatementSid(
					c.BusinessUser.Name,
					IAMIBAP,
					IdentifyAccessKey,
					"WithOneBucketObjectCRUD",
				),
				Effect:       miniointernal.Allow,
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
					s3_api.GetBucketResource(c.BucketName),
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
		Version: miniointernal.Version,
		Id:      makePolicyId(c.BusinessUser.ID),
		Statement: []s3apiv2.Statement{
			//c.RBAPBucketWithAdminAllStatement(), admin already allow any, not required any more
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
		Effect: miniointernal.Allow,
		Principal: map[string]any{
			miniointernal.AWS: makeAWSAccountPrincipal(""),
		},
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
		Effect: miniointernal.Allow,
		Principal: map[string]any{
			miniointernal.AWS: makeAWSAccountPrincipal(c.Cred.AccessKeyID),
		},
		NotPrincipal: nil,
		Action: []string{
			s3_api.S3ListBucket,
			s3_api.S3DeleteObject,
			s3_api.S3GetObject,
			s3_api.S3PutObject,
		},
		NotAction: nil,
		Resource: []string{
			s3_api.GetBucketResource(c.BucketName),
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

func makeAWSAccountPrincipal(minioUserId string) string {
	return s3_api.GetPrincipalAccountRoot(minioUserId)
}
