package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestTypVal_unpack(t *testing.T) {
	var i32 int32 = 10
	unpackEqual(i32, pointerCast(i32, 2)) // dead-loop found
}

type t2 struct{}

func (t t2) IFaceImpl() {
	stdlog.Info("value receiver t2 impl")
}

type t22 struct{}

func (t *t22) IFaceImpl() {
	stdlog.Info("pointer receiver t22 impl")
}

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
	t.Log(Impl(t2{}, iface2)) // only same one, not

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

func testTypeIFace2(t *testing.T, t2v any) {
	inst, inst2 := IndirectImpl(t2v, _iface2)
	if instiface, ok := inst.(iFace2); ok {
		t.Log("value inst")
		instiface.IFaceImpl()
	} else if inst2iface, ok := inst2.(iFace2); ok {
		t.Log("pointer inst")
		inst2iface.IFaceImpl()
	}
	util.TestLongHorizontalLine(t)
}

// value receiver allow any level pointer or slice
func TestIndirectImpl(t *testing.T) {
	testTypeIFace2(t, t2{})
	testTypeIFace2(t, &t2{})
	ptr2 := &t2{}
	testTypeIFace2(t, &ptr2)

	testTypeIFace2(t, [2]t2{})
	testTypeIFace2(t, []t2{})

	testTypeIFace2(t, [2]*t2{})
	testTypeIFace2(t, []*t2{})

	testTypeIFace2(t, [2]**t2{})
	testTypeIFace2(t, []**t2{})

	testTypeIFace2(t, [2]***t2{})
	testTypeIFace2(t, []***t2{})

	testTypeIFace2(t, &[2]t2{})
	testTypeIFace2(t, &[]t2{})

	testTypeIFace2(t, &[2]*t2{})
	testTypeIFace2(t, &[]*t2{})

	testTypeIFace2(t, &[2]**t2{})
	testTypeIFace2(t, &[]**t2{})

	testTypeIFace2(t, &[2]***t2{})
	testTypeIFace2(t, &[]***t2{})
}

// pointer receiver only allow *T, **T, ***T
func TestIndirectImpl2(t *testing.T) {
	testTypeIFace2(t, t22{})
	testTypeIFace2(t, &t22{})
	ptr22 := &t22{}
	testTypeIFace2(t, &ptr22)

	testTypeIFace2(t, [2]t22{})
	testTypeIFace2(t, []t22{})

	testTypeIFace2(t, [2]*t22{})
	testTypeIFace2(t, []*t22{})

	testTypeIFace2(t, [2]**t22{})
	testTypeIFace2(t, []**t22{})

	testTypeIFace2(t, [2]***t22{})
	testTypeIFace2(t, []***t22{})

	testTypeIFace2(t, &[2]t22{})
	testTypeIFace2(t, &[]t22{})

	testTypeIFace2(t, &[2]*t22{})
	testTypeIFace2(t, &[]*t22{})

	testTypeIFace2(t, &[2]**t22{})
	testTypeIFace2(t, &[]**t22{})

	testTypeIFace2(t, &[2]***t22{})
	testTypeIFace2(t, &[]***t22{})
}
