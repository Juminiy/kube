package minio_api

import (
	"fmt"
	"testing"
)

func TestBuiltInPolicyVars(t *testing.T) {
	fmt.Println(consoleAdmin.String())
	fmt.Println(diagnostics.String())
	fmt.Println(readOnly.String())
	fmt.Println(readWrite.String())
	fmt.Println(writeOnly.String())
}
