package s3_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
)

// referred from:
// AWS IAM JSON policy elements reference
// https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json

/*
Please Find out the following Concepts:

	IAM: identity access management

	PBAC: Policy-based access control

	RBAP: Resource-based access policy

	IBAP: Identity-based access policy

	SP: Session Policy

*/

const (
	dirSlash            = "/"
	matchAny            = "*"
	singlePolicyMaxSize = 20 * util.Ki
)

var (
	singlePolicySizeError = errors.New("single policy size bigger than 20KiB")
)

var (
	VersionError         = errors.New("Version value is not valid")
	VersionOutdatedError = errors.New("Version value has outdated")
	StatementError       = errors.New("Statement size can not be zero")
	SidError             = errors.New("Sid are duplicated in full PolicyDocument")
	EffectError          = errors.New("Effect value is not valid")
	ActionError          = errors.New("Action and NotAction cannot be used simultaneously in a same Statement")
	ResourceError        = errors.New("Resource and NotResource cannot be used simultaneously in a same Statement")
	PrincipalError       = errors.New("Principal and NotPrincipal cannot be used simultaneously in a same Statement")
	PrincipalV2Error     = errors.New("Principal can not be used in IBAPolicy Statement")
)

// AWS IAM PolicyDocument JSON
type Policy interface {
	// String to get packed json string
	String() (string, error)
	VersionString() string
	StatementLen() int
}

// RBAPolicy
// Resource-based access policy
type RBAPolicy struct {
	// always set to "2012-10-17"
	// more detail referred to local: testdata/Version
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
	Version string `json:"Version,omitempty"`

	// can be used in Resource-based access policy
	// can not be used in Identity-based access policy
	// suggest to use: UUID/GUID or combine of UUID&ID
	// more detail referred to local: testdata/Id
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_id.html
	// +optional
	Id string `json:"Id,omitempty"`

	// declaration array
	// more detail referred to local: testdata/Statement
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_statement.html
	StatementList StatementList `json:"Statement,omitempty"`
}

func (p *RBAPolicy) String() (string, error) {
	return MarshalPolicy(p)
}

func (p *RBAPolicy) VersionString() string {
	return p.Version
}

func (p *RBAPolicy) StatementLen() int {
	return len(p.StatementList)
}

// IBAPolicy
// Identity-based access policy
type IBAPolicy struct {
	// always set to "2012-10-17"
	// more detail referred to local: testdata/Version
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
	Version string `json:"Version,omitempty"`

	// declaration array
	// more detail referred to local: testdata/Statement
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_statement.html
	StatementList StatementList `json:"Statement,omitempty"`
}

func (p *IBAPolicy) String() (string, error) {
	return MarshalPolicy(p)
}

func (p *IBAPolicy) VersionString() string {
	return p.Version
}

func (p *IBAPolicy) StatementLen() int {
	return len(p.StatementList)
}

type StatementList []Statement

type Statement struct {
	// policy optional identifier
	// +optional each policy statement with a Sid value
	// +optional Sid value as description of its policy statement
	// permit to use: SQS or SNS, Sid value is policy file ID's child-ID
	// must unique in JSON Policy
	// more detail referred to local: testdata/Statement/Sid
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_sid.html
	// +optional
	Sid string `json:"Sid,omitempty"`

	// only valid of: "Allow" and "Deny"
	// any others are invalid
	// more detail referred to local: testdata/Statement/Effect
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_effect.html
	Effect string `json:"Effect,omitempty"`

	// must be used in Resource-based access policy
	// RBAPolicy for example: in Amazon S3 Bucket or AWS KMS Key
	// can not be used in Identity-based access policy
	// IBAPolicy is attached to IAM Identification(Users, Groups or Roles) policy
	// more detail referred to local: testdata/Statement/Principal
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html
	// +optional
	Principal

	// more detail referred to local: testdata/Statement/Action
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_action.html
	Action

	// more detail referred to local: testdata/Statement/Resource
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_resource.html
	Resource

	// more detail referred to local: testdata/Statement/Condition
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html
	// +optional
	Condition ConditionType `json:"Condition,omitempty"`
}

// PrincipalType
// possible gotype: string, map[string]any
type PrincipalType any
type PrincipalTypeConstraint interface {
	~string | ~map[string]any
}
type Principal struct {
	// more detail referred to local: testdata/Statement/Principal
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html
	// +optional
	Principal PrincipalType `json:"Principal,omitempty"`

	// more detail referred to local: testdata/Statement/Principal
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_notprincipal.html
	// +optional
	NotPrincipal PrincipalType `json:"NotPrincipal,omitempty"`
}

// ActionType
// possible gotype: string, []string
type ActionType any
type ActionTypeConstraint interface {
	~string | ~[]string
}
type Action struct {
	// more detail referred to local: testdata/Statement/Action
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_action.html
	// +optional
	Action ActionType `json:"Action,omitempty"`

	// more detail referred to local: testdata/Statement/Action
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_notaction.html
	// +optional
	NotAction ActionType `json:"NotAction,omitempty"`
}

// ResourceType
// possible gotype: string, []string
type ResourceType any
type ResourceTypeConstraint interface {
	~string | ~[]string
}
type Resource struct {
	// more detail referred to local: testdata/Statement/Resource
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_resource.html
	// +optional
	Resource ResourceType `json:"Resource,omitempty"`

	// more detail referred to local: testdata/Statement/Resource
	// and more detail referred to web: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_notresource.html
	// +optional
	NotResource ResourceType `json:"NotResource,omitempty"`
}

// ConditionType
// Extremely Complex, do it when really need it!!!
type ConditionType map[string]any

func PolicyValid(policy Policy) error {
	version := policy.VersionString()
	if version == EarlierVersion {
		return VersionOutdatedError
	}
	if version != Version {
		return VersionError
	}
	if policy.StatementLen() == 0 {
		return StatementError
	}
	return nil
}

func MarshalPolicy(policy Policy) (string, error) {
	policyJSONString, err := util.MarshalJSONPretty(&policy)
	if err != nil {
		return "", err
	}
	if len(policyJSONString) > singlePolicyMaxSize {
		return "", singlePolicySizeError
	}
	return policyJSONString, err
}
