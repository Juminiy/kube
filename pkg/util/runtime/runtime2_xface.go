package runtimeutil

import (
	"unsafe"
)

// mirror of runtime/runtime2.go:
// iface and eface 's funcs and extensions

type IFace struct {
	Tab  unsafe.Pointer
	Data unsafe.Pointer
}

type EFace struct {
	Typ  unsafe.Pointer
	Data unsafe.Pointer

	vp *any
}

// EFaceOf
// +param: vp stand for value pointer
func EFaceOf(vp *any) *EFace {
	return (*EFace)(Tp2Up(vp)).WithValuePtr(vp)
}

func (e *EFace) WithValuePtr(vp *any) *EFace {
	e.vp = vp
	return e
}

func (e *EFace) Type() uintptr {
	return Up2Ui(e.Typ)
}

func (e *EFace) Any() any {
	if e.vp != nil {
		return *e.vp
	}
	return FromEface(e.Typ, e.Data)
}

func FromEface(typ, data unsafe.Pointer) any {
	var v any
	ep := EFaceOf(&v)
	ep.Typ = typ
	ep.Data = data
	return v
}
