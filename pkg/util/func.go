package util

type (
	Fn   func()
	Func func() error
)

var (
	_nothingFn   = func() {}
	_nothingFunc = func() error { return nil }
)

func NothingFn() Fn {
	return _nothingFn
}

func DoNothing() Fn {
	return _nothingFn
}

func NothingFunc() Func {
	return _nothingFunc
}

func SeqRun(fns ...Fn) {
	for _, fn := range fns {
		fn()
	}
}

func ConRun(fns ...Fn) {
	for _, fn := range fns {
		go fn()
	}
}
