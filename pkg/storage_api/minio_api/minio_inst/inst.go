package minio_inst

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api"
)

var (
	_minioEndpoint        string
	_minioAccessKeyID     string
	_minioSecretAccessKey string
	_minioSessionToken    string
	_minioSecure          bool
)

var (
	_minioClient *minio_api.Client
)

func Init() {
	var minioClientError error
	_minioClient, minioClientError = minio_api.New(
		_minioEndpoint,
		_minioAccessKeyID,
		_minioSecretAccessKey,
		_minioSessionToken,
		_minioSecure,
	)
	if minioClientError != nil {
		stdlog.ErrorF("minio client connect to endpoint: %s, AccessKeyID: %s, SecretAccessKey: ******, SessionToken: %s, Secure: %v, error: %s",
			_minioEndpoint, _minioAccessKeyID, _minioSessionToken, _minioSecure, minioClientError.Error())
		return
	}

}
