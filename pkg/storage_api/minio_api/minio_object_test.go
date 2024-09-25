package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed
func TestClient_TempGetObject(t *testing.T) {
	url, err := testMinioClient.TempGetObject(&ObjectConfig{"kube-env", "shell", "info.sh"}, util.DurationMinute*10)
	if err != nil {
		panic(err)
	}
	stdlog.Info(url)
}

// +passed
func TestClient_TempPutObject(t *testing.T) {
	url, err := testMinioClient.TempPutObject(&ObjectConfig{"kube-env", "/shell/folder/", "/cc.sh"}, util.DurationMinute*10)
	if err != nil {
		panic(err)
	}
	stdlog.Info(url)
}
