package minio_api

import (
	"errors"
	miniointernal "github.com/Juminiy/kube/pkg/storage_api/minio_api/minio_internal"
	"github.com/Juminiy/kube/pkg/storage_api/s3_api"
	s3apiv2 "github.com/Juminiy/kube/pkg/storage_api/s3_api/v2"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"github.com/google/uuid"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

type PolicyConfig struct {
	// +required
	BusinessUser BusinessUser

	// Minio's UserName equals to Minio's AccessKeyID
	// +required Cred.AccessKeyID
	Cred miniocred.Value

	// +optional
	GroupName string

	// +required
	BucketName string

	// when delete policy from a user, must provide the PolicyName
	// +required
	// when create policy from a user, api will ignore the PolicyName
	// +optional
	PolicyName string
}

// BusinessUser
// assume the business user is
type BusinessUser struct {
	ID   string
	Name string
}

const (
	IAMIBAP = "IBAP"
	IAMRBAP = "RBAP"

	IdentifyAdmin     = "Admin"
	IdentifyAccessKey = "AccessKey"

	OneBucketObjectCRUD = "WithOneBucketObjectCRUD"
)

var (
	rbapBusinessUserIDError   = errors.New("RBAPolicy, BusinessUser.ID is nil")
	rbapBusinessUserNameError = errors.New("RBAPolicy, BusinessUser.Name is nil")
	rbapBucketNameError       = errors.New("RBAPolicy, BucketName is nil")
	rbapAccessKeyIDError      = errors.New("RBAPolicy, AccessKeyID is nil")
)

var (
	ibapBusinessUserNameError = errors.New("IBAPolicy, BusinessUser.Name is nil")
	ibapBucketNameError       = errors.New("IBAPolicy, BucketName is nil")
)

func (c *PolicyConfig) IBAPAccessKeyWithOneBucketObjectCRUDPolicy() (string, error) {
	if err := c.validateIBAP(); err != nil {
		return "", err
	}
	policy := s3apiv2.IBAPolicy{
		Version: miniointernal.Version,
		Statement: []s3apiv2.Statement{
			{
				Sid: makeStatementSid(
					c.BusinessUser.Name,
					IAMIBAP,
					IdentifyAccessKey,
					OneBucketObjectCRUD,
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
	if err := c.validateRBAP(); err != nil {
		return "", err
	}
	policy := s3apiv2.RBAPolicy{
		Version: miniointernal.Version,
		Id:      makePolicyId(c.BusinessUser.ID),
		Statement: []s3apiv2.Statement{
			//c.RBAPBucketWithAdminAllStatement(), //admin already allow any, not required any more
			c.RBAPBucketWithAccessKeyOneBucketObjectCRUDStatement(), //accessKey must be required
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
			OneBucketObjectCRUD,
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

func (c *PolicyConfig) validateRBAP() error {
	if len(c.BusinessUser.ID) == 0 {
		return rbapBusinessUserIDError
	}
	if len(c.BusinessUser.Name) == 0 {
		return rbapBusinessUserNameError
	}
	if len(c.BucketName) == 0 {
		return rbapBucketNameError
	}
	if len(c.Cred.AccessKeyID) == 0 {
		return rbapAccessKeyIDError
	}
	return nil
}

func (c *PolicyConfig) validateIBAP() error {
	if len(c.BusinessUser.Name) == 0 {
		return ibapBusinessUserNameError
	}
	if len(c.BucketName) == 0 {
		return ibapBucketNameError
	}
	return nil
}

// +self define
func (c *PolicyConfig) setPolicyName() {
	c.PolicyName = util.StringJoin(
		"-",
		c.BusinessUser.ID,
		c.BusinessUser.Name,
		c.BucketName,
		random.IDString(len(c.BusinessUser.Name)),
	)
}

// +self define
func makePolicyId(businessUserID string) string {
	return util.StringJoin(
		"-",
		businessUserID,
		uuid.NewString(),
	)
}

// +self define
func makeStatementSid(s ...string) string {
	return util.StringJoin(
		"-",
		s...,
	)
}

// +aws define
// +minio define
func makeAWSAccountPrincipal(minioUserID string) string {
	return s3_api.GetPrincipalAccountRoot(minioUserID)
}
