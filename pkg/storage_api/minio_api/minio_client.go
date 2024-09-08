package minio_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	AccessKeyIDMaxLen     = 50
	SecretAccessKeyMaxLen = 128
)

type (
	Client struct {
		Endpoint string
		miniocred.Value
		Secure bool

		mc  *minio.Client
		ma  *madmin.AdminClient
		ctx context.Context
	}
)

func New(
	endpoint string,
	accessKeyID string,
	secretAccessKey string,
	sessionToken string,
	secure bool,
) (*Client, error) {
	mOpts := &minio.Options{
		Creds:  miniocred.NewStaticV4(accessKeyID, secretAccessKey, sessionToken),
		Secure: secure,
	}
	mc, err := minio.New(endpoint, mOpts)
	if err != nil {
		stdlog.ErrorF("minio client error: %s", err.Error())
		return nil, err
	}

	ma, err := madmin.NewWithOptions(endpoint, &madmin.Options{Creds: mOpts.Creds, Secure: mOpts.Secure})
	if err != nil {
		stdlog.ErrorF("minio admin client error: %s", err.Error())
		return nil, err
	}

	return &Client{
		mc:  mc,
		ma:  ma,
		ctx: context.TODO(),
	}, nil
}
