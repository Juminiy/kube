package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed array and pointer to array
func TestTypVal_SliceSet(t *testing.T) {
	sl := []string{"aaa", "bbb", "ccc"}
	t.Logf("before %v", sl)

	// no pointer
	Of(sl).SliceSet(2, "vvv")
	t.Logf("no pointer %v", sl)

	// pointer
	Of(&sl).SliceSet(2, "xxx")
	t.Logf("pointer %v", sl)

}

// +passed array pointer and pointer to array pointer
func TestTypVal_SliceSet2(t *testing.T) {
	slp := []*string{util.NewString("aaa"), util.NewString("bbb"), util.NewString("ccc")}
	t.Logf("before %v", *slp[2])

	// no pointer
	Of(slp).SliceSet(2, util.NewString("vvv"))
	t.Logf("no pointer %v %v", *slp[2], slp)

	// pointer
	Of(&slp).SliceSet(2, util.NewString("xxx"))
	t.Logf("pointer %v %v", *slp[2], slp)
}

// +passed all
func TestTypVal_SliceSetStructFields(t *testing.T) {
	sl := make([]t0, 3)
	//src := t0{F0: "no pointer", F1: 69} // no pointer
	//srcPtr := &src                      // p
	//srcPPtr := &srcPtr                  // pp

	Of(sl).SliceSetStructFields(map[string]any{
		"F0": "field F0",
		"F1": "999", // value_type mismatch
	})
	t.Log(sl)

	Of(&sl).SliceSetStructFields(map[string]any{
		"F0": util.NewString("field F0 ptr"), // value_type indirect
		"F1": "999",                          // value_type mismatch
	})
	t.Log(sl)

	Of(&sl).SliceSetStructFields(map[string]any{
		"F0": util.NewString("field F0 pointer"), // value_type indirect
		"F1": 999,                                // value_type ok
	})
	t.Log(sl)
}

func TestTypVal_SliceSetOol(t *testing.T) {
	sl := make([]int, 2, 8)

	t.Log(sl)
	// no pointer
	Of(sl).SliceSet(3, 333) // out of bound length
	t.Log(sl)

	Of(sl).sliceSetLen(3) // out of bound length
	t.Log(sl)

	Of(sl).SliceSetOol(3, 333) // out of bound length
	t.Log(sl)

	Of(sl).SliceSetOol(7, 777) // out of bound capacity
	t.Log(sl)

	Of(sl).SliceSetOol(4, 444) // in bound
	t.Log(sl)

	Of(sl).SliceSet(1, util.New[int](111)) // in bound
	t.Log(sl)

	Of(sl).SliceSetOol(8, 888) // out of bound capacity
	t.Log(sl)

	Of(sl).SliceSetOol(9, 999) // out of bound capacity
	t.Log(sl)
}

func TestTypVal_SliceSetOol2(t *testing.T) {
	sl := make([]int, 2, 8)

	t.Log(sl)
	// no pointer
	Of(&sl).SliceSet(3, 333) // out of bound length
	t.Log(sl)

	Of(&sl).sliceSetLen(3) // out of bound length
	t.Log(sl)

	Of(&sl).SliceSetOol(3, 333) // out of bound length
	t.Log(sl)

	Of(&sl).SliceSetOol(7, 777) // out of bound capacity
	t.Log(sl)

	Of(sl).SliceSetOol(4, 444) // in bound
	t.Log(sl)

	Of(sl).SliceSetOol(1, util.New[int](111)) // in bound
	t.Log(sl)

	Of(&sl).SliceSetOol(8, 888) // out of bound capacity
	t.Log(sl)

	Of(&sl).SliceSetOol(9, 999) // out of bound capacity
	t.Log(sl)
}

func TestTypVal_SliceSetLen(t *testing.T) {
	sl := make([]int, 2, 8)

	Of(&sl).sliceSetLen(3)
	t.Log(len(sl), cap(sl))

	Of(&sl).sliceSetLen(9)
	t.Log(len(sl), cap(sl))

	Of(&sl).sliceShiftLen2Cap()
	t.Log(len(sl), cap(sl))

	sl[7] = 777
	t.Log(len(sl), cap(sl))

	Of(&sl).sliceSetLen(5)
	t.Log(len(sl), cap(sl))

	Of(&sl).sliceShiftLen2Cap()
	t.Log(len(sl), cap(sl))
}

func TestTypVal_SliceSet3(t *testing.T) {
	sl := make([]int, 2, 8)

	// no pointer
	Of(sl).SliceSet(1, 111) // in bound length
	t.Log(sl)

	Of(sl).SliceSet(3, 333) // out of bound length, in bound capacity
	t.Log(sl)

	Of(sl).SliceSet(9, 999) // out of bound capacity
	t.Log(sl)

	// pointer
	Of(&sl).SliceSet(1, 111) // in bound length
	t.Log(sl)

	Of(&sl).SliceSet(3, 333) // out of bound length, in bound capacity
	t.Log(sl)

	Of(&sl).SliceSet(9, 999) // out of bound capacity
	t.Log(sl)
}

func TestTypVal_SliceSetOol3(t *testing.T) {
	sl := make([]int, 2, 8)

	// no pointer
	Of(sl).SliceSetOol(1, 111) // in bound length
	t.Log(sl)

	Of(sl).SliceSetOol(3, 333) // out of bound length, in bound capacity
	t.Log(sl)

	Of(sl).SliceSetOol(9, 999) // out of bound capacity
	t.Log(sl)

	// pointer
	Of(&sl).SliceSetOol(1, 111) // in bound length
	t.Log(sl)

	Of(&sl).SliceSetOol(3, 333) // out of bound length, in bound capacity
	t.Log(sl)

	Of(&sl).SliceSetOol(9, 999) // out of bound capacity
	t.Log(sl)
}

func TestTypVal_SliceSetOoc3(t *testing.T) {
	sl := make([]int, 2, 8)

	// no pointer
	Of(sl).SliceSetOoc(1, 111) // in bound length
	t.Log(sl)

	Of(sl).SliceSetOoc(3, 333) // out of bound length, in bound capacity
	t.Log(sl)

	Of(sl).SliceSetOoc(9, 999) // out of bound capacity
	t.Log(sl)

	// pointer
	Of(&sl).SliceSetOoc(1, 111) // in bound length
	t.Log(sl)

	Of(&sl).SliceSetOoc(3, 333) // out of bound length, in bound capacity
	t.Log(sl)

	Of(&sl).SliceSetOoc(9, 999) // out of bound capacity
	t.Log(sl)
}

func TestTypVal_SliceSetMake(t *testing.T) {
	var sl []int
	Of(sl).SliceSetMake(10, 1)
	t.Log(sl)

	Of(&sl).SliceSetMake(10, "sss")
	t.Log(sl)

	Of(&sl).SliceSetMake(10, 10)
	t.Log(sl)
}

func TestSliceMake(t *testing.T) {
	t.Log(SliceMake(nil, 10, 20))
	t.Log(SliceMake(nil, 5, -1))
	sl := SliceMake(t0{}, 10, 20)
	t.Log(len(sl.([]t0)), cap(sl.([]t0)))
}
