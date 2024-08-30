package s3_api

import (
	"encoding/json"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
)

// gen from:
// https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html

// AWS S3 IAM JSON Policy Element Reference

// RBAP: Resource-based access policy

// IBAP: Identity-based access policy

// SP: Session Policy

const (
	dir                 = "/"
	matchAny            = "*"
	singlePolicyMaxSize = 20 * util.Ki
)

type Policy struct {
	Version       string        `json:"Version,omitempty"`
	StatementList StatementList `json:"Statement,omitempty"`
}

func (p *Policy) String() string {
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if len(bs) >= singlePolicyMaxSize {
		fmt.Println("policy size bigger than 20KiB")
		return ""
	}
	return util.Bytes2StringNoCopy(bs)
}

type AdminStatement struct {
	SASRStatement
}

type AccessKeyWithOneBucketRWStatement struct {
	MASRStatement
}

type BucketStatement struct {
	SAMRStatement
}

type StatementList []Statement

type Statement interface {
	WithName(string)
}

// SASRStatement
// Single Action Single Resource Statement
type SASRStatement struct {
	Sid       string       `json:"Sid,omitempty"`
	Effect    string       `json:"Effect,omitempty"`
	Action    string       `json:"Action,omitempty"`
	Resource  string       `json:"Resource,omitempty"`
	Condition ConditionMap `json:"Condition,omitempty"`
}

func (s *SASRStatement) WithName(name string) {
	s.Sid = name
}

// MASRStatement
// Multiple Action Single Resource Statement
type MASRStatement struct {
	Sid       string       `json:"Sid,omitempty"`
	Effect    string       `json:"Effect,omitempty"`
	Action    []string     `json:"Action,omitempty"`
	Resource  string       `json:"Resource,omitempty"`
	Condition ConditionMap `json:"Condition,omitempty"`
}

func (s *MASRStatement) WithName(name string) {
	s.Sid = name
}

// SAMRStatement
// Single Action Multiple Resource Statement
type SAMRStatement struct {
	Sid       string       `json:"Sid,omitempty"`
	Effect    string       `json:"Effect,omitempty"`
	Action    string       `json:"Action,omitempty"`
	Resource  []string     `json:"Resource,omitempty"`
	Condition ConditionMap `json:"Condition,omitempty"`
}

func (s *SAMRStatement) WithName(name string) {
	s.Sid = name
}

// MAMRStatement
// Multiple Action Multiple Resource Statement
type MAMRStatement struct {
	Sid       string       `json:"Sid,omitempty"`
	Effect    string       `json:"Effect,omitempty"`
	Action    []string     `json:"Action,omitempty"`
	Resource  []string     `json:"Resource,omitempty"`
	Condition ConditionMap `json:"Condition,omitempty"`
}

func (s *MAMRStatement) WithName(name string) {
	s.Sid = name
}

type ConditionMap map[string]any
