package util

import (
	"testing"
	"unsafe"
)

// *Typ <-> unsafe.Pointer <-> uintptr
func TestUnsafeSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	t.Log(arr)
	for idx := range cap(arr) {
		t.Log(*unsafe.SliceData(arr[idx:]))
	}
	_ = append(arr, 6, 7, 8)
	t.Log(arr)
	// not I want
	arrAddr := unsafe.Pointer(unsafe.SliceData(arr))
	for idx := range cap(arr) + 3 {
		elemAddr := unsafe.Add(arrAddr, unsafe.Sizeof(int(0))*uintptr(idx))
		t.Log(*(*int)(elemAddr))
	}
}
