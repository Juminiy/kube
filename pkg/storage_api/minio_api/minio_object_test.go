package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed
func TestClient_TempGetObject(t *testing.T) {
	url, err := _cli.TempGetObject(&ObjectConfig{BucketName: "kube-env", ObjectPath: "shell", ObjectName: "info.sh"}, util.DurationMinute*10)
	if err != nil {
		panic(err)
	}
	t.Log(url)
}

// +passed
func TestClient_TempPutObject(t *testing.T) {
	url, err := _cli.TempPutObject(&ObjectConfig{BucketName: "kube-env", ObjectPath: "/shell/folder", ObjectName: "rr.sh"}, util.DurationMinute*10)
	if err != nil {
		panic(err)
	}
	t.Log(url)
}

// +passed
func TestClient_TempGetObjectList(t *testing.T) {
	urls, err := _cli.TempGetObjectList([]ObjectConfig{
		{BucketName: "kube-env", ObjectPath: "/shell/folder", ObjectName: "cc.sh"},
		{BucketName: "kube-env", ObjectPath: "/shell/folder", ObjectName: "rr.sh"}},
		util.DurationMinute*10)
	if err != nil {
		t.Log(err)
	}
	t.Log(urls)
}
