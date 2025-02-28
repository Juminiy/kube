package gorm_api

import (
	"encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

type Name string

func (n Name) MarshalJSON() ([]byte, error) {
	if n == "" || n == "null" {
		return nil, nil
	}
	return util.String2BytesNoCopy(string(n)), nil
}

func (n *Name) UnmarshalJSON(b []byte) error {
	if str := util.Bytes2StringNoCopy(b); str != "null" {
		*n = Name(str)
	}
	return nil
}

func TestMagicType(t *testing.T) {
	n := Name("null")
	bs, err := json.Marshal(n)
	Err(t, err)
	t.Logf("%s", bs)
}
