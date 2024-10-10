package reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed array and pointer to array
func TestTypVal_ArraySet(t *testing.T) {
	arr := [5]string{"aaa", "bbb", "ccc"}
	t.Logf("before %v", arr[3])

	// no pointer
	Of(arr).ArraySet(3, "vvv")
	t.Logf("no pointer %v", arr[3])

	// pointer
	Of(&arr).ArraySet(3, "xxx")
	t.Logf("pointer %v", arr[3])

}

// +passed array pointer and pointer to array pointer
func TestTypVal_ArraySet2(t *testing.T) {
	arrp := [5]*string{util.NewString("aaa"), util.NewString("bbb"), util.NewString("ccc")}
	t.Logf("before %v", arrp[3])

	// no pointer
	Of(arrp).ArraySet(3, util.NewString("vvv"))
	t.Logf("no pointer %v", arrp[3])

	// pointer
	Of(&arrp).ArraySet(3, util.NewString("xxx"))
	t.Logf("pointer %v", arrp[3])
}
