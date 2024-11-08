package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	gostruntime "github.com/dubbogo/gost/runtime"
	"runtime/debug"
)

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

// Recover any func panic anyway, never hangout
func Recover(fn Fn) {
	defer func() {
		PanicHandler()
	}()
	if fn != nil {
		fn()
	}
}

func GoSafe(fn Fn) {
	gostruntime.GoSafely(nil, false, fn, nil)
}

func PanicHandler(v ...any) {
	if r := recover(); r != nil {
		stdlog.ErrorF("panic sth: %v, hangup from recover panic: %v, stack: %s", v, r, Bytes2StringNoCopy(debug.Stack()))
	}
}
