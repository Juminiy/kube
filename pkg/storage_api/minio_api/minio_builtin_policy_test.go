package minio_api

import (
	"testing"

	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

func TestBuiltInPolicyVars(t *testing.T) {
	stdlog.Info(consoleAdmin.String())
	stdlog.Info(diagnostics.String())
	stdlog.Info(readOnly.String())
	stdlog.Info(readWrite.String())
	stdlog.Info(writeOnly.String())
}
