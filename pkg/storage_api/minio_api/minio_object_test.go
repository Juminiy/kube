package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed
func TestClient_TempGetObject(t *testing.T) {
	url, err := testMinioClient.TempGetObject(&ObjectConfig{BucketName: "kube-env", ObjectPath: "shell", ObjectName: "info.sh"}, util.DurationMinute*10)
	if err != nil {
		panic(err)
	}
	stdlog.Info(url)
}

// +passed
func TestClient_TempPutObject(t *testing.T) {
	url, err := testMinioClient.TempPutObject(&ObjectConfig{BucketName: "kube-env", ObjectPath: "/shell/folder", ObjectName: "rr.sh"}, util.DurationMinute*10)
	if err != nil {
		panic(err)
	}
	stdlog.Info(url)
}

// +passed
func TestClient_TempGetObjectList(t *testing.T) {
	urls, err := testMinioClient.TempGetObjectList([]ObjectConfig{
		{BucketName: "kube-env", ObjectPath: "/shell/folder", ObjectName: "cc.sh"},
		{BucketName: "kube-env", ObjectPath: "/shell/folder", ObjectName: "rr.sh"}},
		util.DurationMinute*10)
	if err != nil {
		t.Log(err)
	}
	t.Log(urls)
}
