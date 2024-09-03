package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

func SilentHandleError(handle string, err error) {
	if err != nil {
		consoleLogError(handle, err)
	}
}

func consoleLogError(detail string, err error) {
	stdlog.ErrorF("%s: %v\n", detail, err)
}
