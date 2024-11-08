package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

const (
	endpoint        = "192.168.31.131:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	sessionToken    = ""
	secure          = false
)

var (
	testMinioClient, testMinioClientError = New(
		endpoint,
		accessKeyID,
		secretAccessKey,
		sessionToken,
		secure,
	)
)

// +passed
func TestClient_AtomicWorkflow(t *testing.T) {
	resp, err := testMinioClient.AtomicWorkflow(Req{
		UserID:          16,
		UserName:        "calimimicc",
		BucketQuotaByte: 114514,
		BucketName:      "funkqqqzaaweqaadfafwq4124asf1122",
	})
	if err != nil {
		panic(err)
	}
	respJSONStr, err := util.MarshalJSONPretty(resp)
	if err != nil {
		panic(err)
	}
	stdlog.Info(respJSONStr)
}
