package safe_reflect

import (
	"reflect"
	"testing"
)

// +passed
func TestDeref(t *testing.T) {
	var i32 int32 = 10
	i32Ptr := &i32
	i32PPtr := &i32Ptr
	i32PPPtr := &i32PPtr

	t.Logf("before deref type: %v", reflect.TypeOf(i32PPPtr).String())
	t.Logf("after deref type: %v", onePointer(reflect.ValueOf(i32PPPtr)).Type().String())
	t.Logf("after deref type: %v", noPointer(reflect.ValueOf(i32PPPtr)).Type().String())
}

// +unpassed no-sense
func TestDerefInterface(t *testing.T) {
	interfacePtr := cast2Pointer(3, 1)
	iptr := interfacePointer(reflect.ValueOf(interfacePtr))
	t.Log(iptr)
}
