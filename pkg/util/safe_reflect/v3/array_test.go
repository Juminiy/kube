package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

var arr0 [1]ExportStruct
var arr1 = [1]ExportStruct{sv}

var arrIS = []Tv{
	Indirect(arr0),
	Indirect(&arr0),
	Indirect(util.New(&arr0)),
	Indirect(arr1),
	Indirect(&arr1),
	Indirect(util.New(&arr1)),
}

func TestT_ArrayStructFields(t *testing.T) {
	for _, etv := range arrIS {
		t.Log(String(etv.ArrayStructFields()))
	}
}

func TestT_ArrayStructTags(t *testing.T) {
	for _, etv := range arrIS {
		t.Log(String(etv.ArrayStructTags(`gorm`)))
		t.Log(String(etv.ArrayStructTags(`json`)))
	}
}

func TestT_ArrayStructTypes(t *testing.T) {
	for _, etv := range arrIS {
		t.Log(etv.ArrayStructTypes())
	}
}

func TestTv_ArrayStructValues(t *testing.T) {
	for _, etv := range arrIS {
		t.Log(String(etv.ArrayStructValues()))
	}
}
