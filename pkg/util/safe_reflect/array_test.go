package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed array and pointer to array
func TestTypVal_ArraySet(t *testing.T) {
	arr := [5]string{"aaa", "bbb", "ccc"}
	t.Logf("before %v", arr)

	// no pointer
	Of(arr).ArraySet(3, "vvv")
	t.Logf("no pointer %v", arr)

	// pointer
	Of(&arr).ArraySet(3, "xxx")
	t.Logf("pointer %v", arr)

}

// +passed array pointer and pointer to array pointer
func TestTypVal_ArraySet2(t *testing.T) {
	arrp := [5]*string{util.NewString("aaa"), util.NewString("bbb"), util.NewString("ccc")}
	t.Logf("before %v", arrp)

	// no pointer
	Of(arrp).ArraySet(3, util.NewString("vvv"))
	t.Logf("no pointer %v", arrp)

	// pointer
	Of(&arrp).ArraySet(3, util.NewString("xxx"))
	t.Logf("pointer %v", arrp)

	t.Logf("before %v %v", *arrp[2], arrp)
	// no pointer
	Of(arrp).ArraySet(2, util.NewString("mmm"))
	t.Logf("no pointer %v %v", *arrp[2], arrp)

	// pointer
	Of(&arrp).ArraySet(2, util.NewString("nnn"))
	t.Logf("pointer %v %v", *arrp[2], arrp)
}

// +passed all
func TestTypVal_ArraySetStructFields(t *testing.T) {
	arr := [5]t0{}
	//src := t0{F0: "no pointer", F1: 69} // no pointer
	//srcPtr := &src                      // p
	//srcPPtr := &srcPtr                  // pp

	Of(arr).ArraySetStructFields(map[string]any{
		"F2": -123,
		"F0": "field F0",
		"F1": "999", // value_type mismatch
	})
	t.Log(arr)

	Of(&arr).ArraySetStructFields(map[string]any{
		"F0": util.NewString("field F0"), // value_type indirect
		"F1": "999",                      // value_type mismatch
	})
	t.Log(arr)

	Of(&arr).ArraySetStructFields(map[string]any{
		"F0": util.NewString("field F0 pointer"), // value_type indirect
		"F1": 999,                                // value_type ok
	})
	t.Log(arr)
}
