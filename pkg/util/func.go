package util

type (
	Fn   func()
	Func func() error
)

var (
	NothingFn Fn = func() {}
	DoNothing Fn = func() {}
)
