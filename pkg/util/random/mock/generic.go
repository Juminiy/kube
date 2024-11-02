package mock

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
)

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
