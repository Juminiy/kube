package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

var sl0 []ExportStruct
var sl1 = []ExportStruct{sv}

var slIS = []Tv{
	Indirect(sl0),
	Indirect(&sl0),
	Indirect(util.New(&sl0)),
	Indirect(sl1),
	Indirect(&sl1),
	Indirect(util.New(&sl1)),
}

func TestT_SliceStructFields(t *testing.T) {
	for _, etv := range slIS {
		t.Log(String(etv.SliceStructFields()))
	}
}

func TestT_SliceStructTags(t *testing.T) {
	for _, etv := range slIS {
		t.Log(String(etv.SliceStructTags(`gorm`)))
		t.Log(String(etv.SliceStructTags(`json`)))
	}
}

func TestT_SliceStructTypes(t *testing.T) {
	for _, etv := range slIS {
		t.Log(etv.SliceStructTypes())
	}
}

func TestTv_SliceStructValues(t *testing.T) {
	for _, etv := range slIS {
		t.Log(String(etv.SliceStructValues()))
	}
}
