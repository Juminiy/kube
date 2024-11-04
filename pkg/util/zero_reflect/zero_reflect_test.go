package zero_reflect

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/modern-go/reflect2"
	"testing"
	"unsafe"
)

func TestGetSet(t *testing.T) {
	var i int = 114
	var j int = 514
	reflect2.TypeOf(i).Set(&i, &j)
	t.Log(i, j)

	t.Log(reflect2.TypeOf(nil))

	t.Log(reflect2.TypeOf(10))

	reflect2.TypeOf(i).Set(&i, &j)
	t.Log()
}

func printBinary(v any) {
	pv := unsafe.Pointer(&v)
	for addByte := range unsafe.Sizeof(v) {
		vbyteaddr := unsafe.Pointer(uintptr(pv) + uintptr(addByte))
		vbyteval := *(*uint8)(vbyteaddr)
		stdlog.InfoF("%x", vbyteval)
	}
	stdlog.Info("--------")
}

func TestUnsafe(t *testing.T) {
	type t0 struct {
		Int  int
		Int8 int8
	}
	v0 := t0{Int: 0x0000FFFF, Int8: 0xF}
	//t.Log(unsafe.Sizeof(v0))

	//pv0 := unsafe.Pointer(&v0)
	//for addByte := range 16 {
	//	v0addr := unsafe.Pointer(uintptr(pv0) + uintptr(addByte))
	//	v0addrval := *(*uint8)(v0addr)
	//	t.Logf("%x", v0addrval)
	//}
	printBinary(v0)

	type t1 struct {
		Int8 int8
		Int  int
	}
	v1 := t1{Int8: 0xF, Int: 0x0000FFFF}
	//t.Log(unsafe.Sizeof(v1))

	printBinary(v1)

	t.Logf("%x %x", v1.Int8, v1.Int)

	v1 = *(*t1)(unsafe.Pointer(&v0))

	t.Logf("%x %x", v1.Int8, v1.Int)
}
