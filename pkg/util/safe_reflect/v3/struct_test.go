package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

var sv = ExportStruct{
	Int:    114514,
	Float:  1919.810893,
	String: "IamHaJiMi",
}

var svIS = []Tv{
	Indirect(sv),
	Indirect(&sv),
	Indirect(util.New(&sv)),
}

func TestT_StructFields(t *testing.T) {
	for _, etv := range svIS {
		t.Log(Pretty(etv.StructFields()))
	}
}

func TestT_StructTags(t *testing.T) {
	for _, etv := range svIS {
		t.Log(Pretty(etv.StructTags(`gorm`)))
		t.Log(Pretty(etv.StructTags(`json`)))
	}
}

func Test_parseTagVal(t *testing.T) {
	for _, s := range []string{
		"k1:v1;k2:v2;k3;k4",  // k1:v1 k2:v2 k3: k4:
		"k1:-;k2:v1,v2;",     // k1:- k2:v1,v2
		":v1;k1:;:;k2:v2:v3", // k1: k2:v2:v3
		"name,omitempty",     // name: omitempty:
		",inline",            // inline:
	} {
		t.Log(Pretty(parseTagVal(s)))
	}
}

func TestT_StructTypes(t *testing.T) {
	for _, etv := range svIS {
		t.Log(Pretty(etv.StructTypes()))
	}
}

func TestTv_StructValues(t *testing.T) {
	for _, etv := range svIS {
		t.Log(Pretty(etv.StructValues()))
	}
}
