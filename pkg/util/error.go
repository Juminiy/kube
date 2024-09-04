package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

// SilentHandleError
// Deprecated
// only used in dev env or test env, _test file
// not to use in production env
func SilentHandleError(handle string, err error) {
	if err != nil {
		consoleLogError(handle, err)
	}
}

// consoleLogError
// Deprecated
// only used in dev env or test env, _test file
// not to use in production env
func consoleLogError(detail string, err error) {
	stdlog.ErrorF("%s: %v\n", detail, err)
}
