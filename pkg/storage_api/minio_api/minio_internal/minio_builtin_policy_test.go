package minio_internal

import (
	"testing"
)

func TestBuiltInPolicyVars(t *testing.T) {
	t.Log(ConsoleAdminPolicy())
	t.Log(DiagnosticsPolicy())
	t.Log(ReadOnlyPolicy())
	t.Log(ReadWritePolicy())
	t.Log(WriteOnlyPolicy())
}
