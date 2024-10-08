package reflect

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
	t.Logf("after deref type: %v", deref2OnePointer(reflect.ValueOf(i32PPPtr)).Type().String())
	t.Logf("after deref type: %v", deref2NoPointer(reflect.ValueOf(i32PPPtr)).Type().String())
}

// +unpassed
func TestDerefInterface(t *testing.T) {
	interfacePtr := cast2Pointer(3, 1)
	iptr := derefInterfacePointer(reflect.ValueOf(interfacePtr))
	t.Log(iptr)
}

func TestMap(t *testing.T) {
	// nil map
	var m1 map[string]string
	mapKeyExistAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Log(m1)

	mapDryAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Log(m1)

	// len0 map assign key exist
	m1 = make(map[string]string)
	mapKeyExistAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Log(m1)

	// len0 map assign
	//m1 = make(map[string]string)
	mapDryAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Log(m1)

	// map delete
	mapDryDelete(reflect.ValueOf(m1), "v1")
	t.Log(m1)

	// panic
	mapDryAssign(reflect.ValueOf(m1), "v2", 1)
	t.Log(m1)
}
