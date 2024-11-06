// Package mockv2 was generated
package mockv2

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
)

func Slice(v any) {
	indirTv := indir(v)
	if indirTv.Typ.Kind() != safe_reflect.Slice {
		return
	}

	sliceOrArraySet(indirTv)
}

func sliceOrArraySet(indirTv safe_reflect.TypVal) {
	for i := range indirTv.FieldLen() {
		elemTv := safe_reflect.Wrap(indirTv.Val.Index(i))
		structSet(elemTv, structTv(elemTv).typ)
	}
}
