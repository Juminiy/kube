package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed array and pointer to array
func TestTypVal_SliceSet(t *testing.T) {
	arr := []string{"aaa", "bbb", "ccc"}
	t.Logf("before %v", arr)

	// no pointer
	Of(arr).SliceSet(2, "vvv")
	t.Logf("no pointer %v", arr)

	// pointer
	Of(&arr).SliceSet(2, "xxx")
	t.Logf("pointer %v", arr)

}

// +passed array pointer and pointer to array pointer
func TestTypVal_SliceSet2(t *testing.T) {
	arrp := []*string{util.NewString("aaa"), util.NewString("bbb"), util.NewString("ccc")}
	t.Logf("before %v", *arrp[2])

	// no pointer
	Of(arrp).SliceSet(2, util.NewString("vvv"))
	t.Logf("no pointer %v %v", *arrp[2], arrp)

	// pointer
	Of(&arrp).SliceSet(2, util.NewString("xxx"))
	t.Logf("pointer %v %v", *arrp[2], arrp)
}

// +passed all
func TestTypVal_SliceSetStructFields(t *testing.T) {
	arr := make([]t0, 3)
	//src := t0{F0: "no pointer", F1: 69} // no pointer
	//srcPtr := &src                      // p
	//srcPPtr := &srcPtr                  // pp

	Of(arr).SliceSetStructFields(map[string]any{
		"F0": "field F0",
		"F1": "999", // value_type mismatch
	})
	t.Log(arr)

	Of(&arr).SliceSetStructFields(map[string]any{
		"F0": util.NewString("field F0 ptr"), // value_type indirect
		"F1": "999",                          // value_type mismatch
	})
	t.Log(arr)

	Of(&arr).SliceSetStructFields(map[string]any{
		"F0": util.NewString("field F0 pointer"), // value_type indirect
		"F1": 999,                                // value_type ok
	})
	t.Log(arr)
}

func TestTypVal_SliceSetOob(t *testing.T) {
	arrp := []*string{util.NewString("aaa"), util.NewString("bbb"), util.NewString("ccc"), nil}
	t.Logf("before %v", *arrp[2])

	// no pointer
	Of(arrp).SliceSetOob(3, util.NewString("vvv"))
	t.Logf("no pointer %v", arrp)

	// pointer
	Of(&arrp).SliceSetOob(3, util.NewString("xxx"))
	t.Logf("pointer %v", arrp)
}
