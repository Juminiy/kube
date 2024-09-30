package minio_internal

import (
	"testing"

	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

func TestBuiltInPolicyVars(t *testing.T) {
	stdlog.Info(ConsoleAdminPolicy())
	stdlog.Info(DiagnosticsPolicy())
	stdlog.Info(ReadOnlyPolicy())
	stdlog.Info(ReadWritePolicy())
	stdlog.Info(WriteOnlyPolicy())
}
