// Package mockv2 was generated
package mockv2

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
)

// before call with slice must use with randT(v...)
func randT[T any](v ...T) T {
	if len(v) == 0 {
		return util.Zero[T]()
	}
	return v[gofakeit.IntN(len(v))]
}

type FuncT[T any] func() T

func randFunc[FnT FuncT[T], T any](fns ...FnT) FnT {
	return randT[FnT](fns...)
}
