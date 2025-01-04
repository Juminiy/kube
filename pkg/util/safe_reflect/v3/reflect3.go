package safe_reflectv3

import (
	"reflect"
)

type Tv struct {
	T
	V
}

func (tv Tv) Indirect() Tv {
	if iv := tv.V.Indirect(); iv.IsValid() {
		return Tv{
			T: WrapT(iv.Type()),
			V: iv,
		}
	}
	return tv
}

// Indirect
// if i is-valid; return Indirect T, Indirect V
// if i is-not-valid; return Direct T, Direct V
func Indirect(i any) Tv {
	return Direct(i).Indirect()
}

func Direct(i any) Tv {
	return Tv{
		T: NewT(i),
		V: NewV(i),
	}
}

func Wrap(rv reflect.Value) Tv {
	v := WrapV(rv)
	if v.IsValid() {
		return Tv{
			T: WrapT(v.Type()),
			V: v,
		}
	}
	return Tv{
		T: WrapT(nil),
		V: v,
	}
}

func WrapI(rv reflect.Value) Tv {
	return Wrap(rv).Indirect()
}
