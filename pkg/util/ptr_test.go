package util

import (
	"maps"
	"slices"
	"testing"
)

func TestToElemPtrSlice(t *testing.T) {
	s := []string{"sq", "rw", "cs"}
	slices.All(ToElemPtrSlice(s))(func(i int, s *string) bool {
		t.Log(*s)
		return true
	})
}

func TestToElemPtrMap(t *testing.T) {
	m := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	maps.All(ToElemPtrMap(m))(func(k string, v *string) bool {
		t.Log(k, *v)
		return true
	})
}

func testLogPtrAndValue[T any](t *testing.T, ptr *T) {
	t.Log(ptr, PtrValue(ptr))
}

func TestPtrFunc(t *testing.T) {
	testLogPtrAndValue(t, PtrFunc(Min, nil, New(0), New(1), nil))
	testLogPtrAndValue(t, PtrFunc(Max, nil, New(-9), New(111), nil))

	var i0ptr *int
	testLogPtrAndValue(t, PtrFunc(Min, nil, i0ptr))
	testLogPtrAndValue(t, PtrFunc(Max, nil, i0ptr))
}

func TestPtrPairMin(t *testing.T) {
	testLogPtrAndValue(t, PtrPairMin(nil, New(0)))
	testLogPtrAndValue(t, PtrPairMin(New(1), New(0)))

	var i0ptr *int
	testLogPtrAndValue(t, PtrPairMin(nil, i0ptr))
}

func TestPtrPairMax(t *testing.T) {
	testLogPtrAndValue(t, PtrPairMax(nil, New(0)))
	testLogPtrAndValue(t, PtrPairMax(New(1), New(0)))

	var i0ptr *int
	testLogPtrAndValue(t, PtrPairMax(nil, i0ptr))
}
