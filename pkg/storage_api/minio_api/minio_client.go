// Package minio_api
package minio_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	AccessKeyIDMaxLen     = 50
	SecretAccessKeyMaxLen = 128
)

//go:generate go run codegen/codegen.go
type Client struct {
	mc   *minio.Client
	ma   *madmin.AdminClient
	ctx  context.Context
	page *util.Page
}

func New(
	endpoint string,
	accessKeyID string,
	secretAccessKey string,
	sessionToken string,
	secure bool,
) (*Client, error) {
	accessCred := miniocred.NewStaticV4(accessKeyID, secretAccessKey, sessionToken)

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  accessCred,
		Secure: secure,
	})
	if err != nil {
		stdlog.ErrorF("minio client error: %s", err.Error())
		return nil, err
	}

	ma, err := madmin.NewWithOptions(endpoint, &madmin.Options{
		Creds:  accessCred,
		Secure: secure,
	})
	if err != nil {
		stdlog.ErrorF("minio admin client error: %s", err.Error())
		return nil, err
	}

	return &Client{
		mc:   mc,
		ma:   ma,
		ctx:  util.TODOContext(),
		page: util.DefaultPage(),
	}, nil
}

func NewWithOpts(
	endpoint string,
	cliOpts *minio.Options,
	adminOpts *madmin.Options,
) (*Client, error) {
	mc, err := minio.New(endpoint, cliOpts)
	if err != nil {
		return nil, err
	}
	ma, err := madmin.NewWithOptions(endpoint, adminOpts)
	if err != nil {
		return nil, err
	}
	return &Client{
		mc:   mc,
		ma:   ma,
		ctx:  util.TODOContext(),
		page: util.DefaultPage(),
	}, nil
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) WithPage(page *util.Page) *Client {
	c.page = page
	return c
}

// GC
// garbage collection of resource that allocated:
// 1. Object
// 2. Bucket
// 3. AccessKey
// 4. IBAPolicy
func (c *Client) GC(...util.Func) {}
