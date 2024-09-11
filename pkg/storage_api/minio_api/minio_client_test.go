package minio_api

import (
	"testing"
)

const (
	endpoint        = "192.168.31.110:9000"
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

func TestClient_AtomicWorkflow(t *testing.T) {

}
