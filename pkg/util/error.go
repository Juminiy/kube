package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"io"
)

// SilentPanicError
// only used in dev env or test env, _test file
// not to use in production env
func SilentPanicError(err error) {
	if err != nil {
		panic(err)
	}
}

// SilentHandleError
// only used in dev env or test env, _test file
// not to use in production env
func SilentHandleError(handle string, err error) {
	if err != nil {
		consoleLogError(handle, err)
	}
}

// consoleLogError
// only used in dev env or test env, _test file
// not to use in production env
func consoleLogError(detail string, err error) {
	stdlog.ErrorF("%s: %v\n", detail, err)
}

// HandleCloseError
// handle io closer error
func HandleCloseError(msg string, closer io.Closer) {
	err := closer.Close()
	if err != nil {
		stdlog.ErrorF(msg+" instance: %#v close error: %s", closer, err)
	}
}
