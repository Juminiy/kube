package util

type (
	Fn   func()
	Func func() error
)

var (
	NothingFn Fn = func() {}
	DoNothing Fn = func() {}
)

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
