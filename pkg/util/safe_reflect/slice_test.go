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

func TestTypVal_SliceAppend(t *testing.T) {
	var sl []int
	Of(sl).SliceAppend(1)
	t.Log(sl)

	Of(&sl).SliceAppend(2)
	t.Log(sl)

	Of(&sl).SliceAppend("sss")
	t.Log(sl)
}

func TestTypVal_SliceAppendSlice(t *testing.T) {
	var sl []int
	Of(sl).SliceAppendSlice([]int{1, 2, 3})
	t.Log(sl)

	Of(&sl).SliceAppendSlice([]int{1, 2, 3})
	t.Log(sl)

	Of(&sl).SliceAppendSlice([]string{"111", "222", "333"})
	t.Log(sl)
}

func TestTypVal_SliceAppends(t *testing.T) {
	var sl []int

	Of(sl).SliceAppends(1, 2, 3)
	t.Log(sl)

	Of(&sl).SliceAppends(1, 2, 3)
	t.Log(sl)

	Of(&sl).SliceAppends("111", "222", "333")
	t.Log(sl)
}

func TestTypVal_SliceStructFieldsValues(t *testing.T) {
	var sl []int

	t.Log(Of(sl).SliceStructFieldsValues(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	t.Log(Of(&sl).SliceStructFieldsValues(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	var t0sl = []t0{{F0: "v1", F1: 1}, {F0: "v2", F1: 2}}
	t.Log(Of(t0sl).SliceStructFieldsValues(map[string]struct{}{}))
	t.Log(Of(t0sl).SliceStructFieldsValues(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	t.Log(Of(&t0sl).SliceStructFieldsValues(map[string]struct{}{}))
	t.Log(Of(&t0sl).SliceStructFieldsValues(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	t.Log(Of(&t0sl).SliceStructFieldValues("F0"))
	t.Log(Of(&t0sl).SliceStructFieldValues("F1"))
	t.Log(Of(&t0sl).SliceStructFieldValues("Fx"))
}

func TestTypVal_SliceStruct2SliceMap(t *testing.T) {
	var sl []int = []int{1, 2, 3}

	t.Log(Of(sl).SliceStruct2SliceMap(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	t.Log(Of(&sl).SliceStruct2SliceMap(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	var t0sl = []t0{{F0: "v1", F1: 1}, {F0: "v2", F1: 2}}
	t.Log(Of(t0sl).SliceStruct2SliceMap(map[string]struct{}{}))
	t.Log(Of(t0sl).SliceStruct2SliceMap(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

	t.Log(Of(&t0sl).SliceStruct2SliceMap(map[string]struct{}{}))
	t.Log(Of(&t0sl).SliceStruct2SliceMap(map[string]struct{}{
		"F0": {},
		"F1": {},
	}))

}

func TestTypVal_StructMakeSlice(t *testing.T) {
	v := Of(t0{}).StructMakeSlice(10, 10)
	vassertV := v.([]t0)
	vassertV[1] = t0{F0: "1f0", F1: 111}
	vassertV[8] = t0{F0: "8f0", F1: 888}
	t.Log(v)
}

func TestTypVal_StructMakeSlice2(t *testing.T) {
	v := Of(&t0{}).StructMakeSlice(10, 10)
	vassertV := v.([]t0)
	vassertV[1] = t0{F0: "1f0", F1: 111}
	vassertV[8] = t0{F0: "8f0", F1: 888}
	t.Log(v)
}
