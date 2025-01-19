package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
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

func TestV_SliceSet(t *testing.T) {
	for _, isl := range [][]int{{}, {1, 2, 3, 4, 5}} {
		testWatchDo(t, &isl, func() {
			for i := range 10 {
				Indirect(isl).SliceSet(i, (i+1)*111)
			}
		})
		testWatchDo(t, &isl, func() {
			for i := range 10 {
				Indirect(&isl).SliceSet(i, (i+1)*111)
			}
		})
	}
}

func TestV_SliceAppend(t *testing.T) {
	for _, isl := range [][]int{nil, {}, {0, 0, 0, 0, 0}} {
		testWatchDo(t, &isl, func() {
			for i := range 10 {
				Indirect(isl).SliceAppend(i + 1)
			}
		})
		testWatchDo(t, &isl, func() {
			for i := range 10 {
				Indirect(&isl).SliceAppend((i + 1) * 111)
			}
		})
	}
}

func testWatchDo(t *testing.T, v any, do func()) {
	oldv := util.DeepCopyByJSON(safe_json.GoCCY(), v)
	do()
	t.Logf("%v -> %v", oldv, v)
}
