package v2

import (
	s3_api2 "github.com/Juminiy/kube/pkg/storage_api/s3_api"
	"k8s.io/apimachinery/pkg/util/sets"
)

// the following JSON reference are compatible with s3_api.Policy
// for much easier to use
// s3_api/v2 can refer api in s3_api
// s3_api can not refer api in s3_api/v2

var (
	ActionAll    = []string{s3_api2.ActionAll}
	PrincipalAll = []string{s3_api2.PrincipalAll}
	ResourceAll  = []string{s3_api2.ResourceAll}
)

type RBAPolicy struct {
	Version   string      `json:"Version,omitempty"`
	Id        string      `json:"Id,omitempty"`
	Statement []Statement `json:"Statement,omitempty"`
}

func (p *RBAPolicy) String() (string, error) {
	if validErr := p.Valid(); validErr != nil {
		return "", validErr
	}
	return s3_api2.MarshalPolicy(p)
}

func (p *RBAPolicy) VersionString() string {
	return p.Version
}

func (p *RBAPolicy) StatementLen() int {
	return len(p.Statement)
}

func (p *RBAPolicy) Valid() error {
	if policyErr := s3_api2.PolicyValid(p); policyErr != nil {
		return policyErr
	}
	sidMap := sets.Set[string]{}
	for _, sm := range p.Statement {
		if sidMap.Has(sm.Sid) {
			return s3_api2.SidError
		}
		sidMap.Insert(sm.Sid)
		if smErr := sm.Valid(); smErr != nil {
			return smErr
		}
	}
	return nil
}

type IBAPolicy struct {
	Version   string      `json:"Version,omitempty"`
	Statement []Statement `json:"Statement,omitempty"`
}

func (p *IBAPolicy) String() (string, error) {
	if validErr := p.Valid(); validErr != nil {
		return "", validErr
	}
	return s3_api2.MarshalPolicy(p)
}

func (p *IBAPolicy) VersionString() string {
	return p.Version
}

func (p *IBAPolicy) StatementLen() int {
	return len(p.Statement)
}

func (p *IBAPolicy) Valid() error {
	if policyErr := s3_api2.PolicyValid(p); policyErr != nil {
		return policyErr
	}
	for _, sm := range p.Statement {
		if smErr := sm.Valid(); smErr != nil {
			return smErr
		}
		if principalErr := sm.IBAPPrincipalValid(); principalErr != nil {
			return principalErr
		}
	}
	return nil
}

type Statement struct {
	Sid          string         `json:"Sid,omitempty"`
	Effect       string         `json:"Effect,omitempty"`
	Principal    map[string]any `json:"Principal,omitempty"`
	NotPrincipal map[string]any `json:"NotPrincipal,omitempty"`
	Action       []string       `json:"Action,omitempty"`
	NotAction    []string       `json:"NotAction,omitempty"`
	Resource     []string       `json:"Resource,omitempty"`
	NotResource  []string       `json:"NotResource,omitempty"`
}

func (s *Statement) Valid() error {
	if s.Effect != s3_api2.Allow &&
		s.Effect != s3_api2.Deny {
		return s3_api2.EffectError
	}
	if len(s.Principal) != 0 &&
		len(s.NotPrincipal) != 0 {
		return s3_api2.PrincipalError
	}
	if len(s.Action) != 0 &&
		len(s.NotAction) != 0 {
		return s3_api2.ActionError
	}
	if len(s.Resource) != 0 &&
		len(s.NotResource) != 0 {
		return s3_api2.ResourceError
	}
	return nil
}

func (s *Statement) IBAPPrincipalValid() error {
	if len(s.Principal) != 0 ||
		len(s.NotPrincipal) != 0 {
		return s3_api2.PrincipalV2Error
	}
	return nil
}
