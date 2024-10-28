package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestTypVal_unpack(t *testing.T) {
	var i32 int32 = 10
	unpackEqual(i32, pointerCast(i32, 2)) // dead-loop found
}

type t2 struct{}

func (t t2) IFaceImpl() {}

type t22 struct{}

func (t *t22) IFaceImpl() {}

type iFace2 interface {
	IFaceImpl()
}

var _iface2a = iFace2(nil)
var _iface2 = (*iFace2)(nil)

func TestImpl(t *testing.T) {
	t.Log(Impl(t2{}, _iface2))   // value receiver impl, value receiver ok
	t.Log(Impl(&t2{}, _iface2))  // value receiver impl, pointer receiver ok
	t.Log(Impl(t22{}, _iface2))  // pointer receiver impl, value receiver not
	t.Log(Impl(&t22{}, _iface2)) // pointer receiver impl, pointer receiver ok

	var iface2 iFace2
	iface2 = t2{}
	t.Log(Impl(t2{}, iface2)) // only same one

	util.TestLongHorizontalLine(t)

	switch iface2.(type) {
	case t2:
		t.Log("t2")
	case *t2:
		t.Log("*t2")
	}

	util.TestLongHorizontalLine(t)

	iface2 = &t2{}
	switch iface2.(type) {
	case t2:
		t.Log("t2")
	case *t2:
		t.Log("*t2")
	}

	util.TestLongHorizontalLine(t)
	//iface2 = t22{} // error
	iface2 = &t22{} // ok

}
