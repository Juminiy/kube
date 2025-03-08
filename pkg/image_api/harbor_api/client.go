// Package harbor_api
package harbor_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/harbor"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"net/http"
	"time"
)

//go:generate go run codegen/codegen.go
type Client struct {
	// global variant
	v2Cli       *v2client.HarborAPI
	ctx         context.Context
	httpCli     *http.Client
	httpTimeout time.Duration
	pageConfig  *util.Page
}

func New(
	harborRegistry string,
	harborInsecure bool,
	harborUsername string,
	harborPassword string) (*Client, error) {
	if harborInsecure {
		harborRegistry = util.URLWithHTTP(harborRegistry)
	} else {
		harborRegistry = util.URLWithHTTPS(harborRegistry)
	}
	hCli, err := harbor.NewClientSet(&harbor.ClientSetConfig{
		URL:      harborRegistry,
		Insecure: harborInsecure,
		Username: harborUsername,
		Password: harborPassword,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		v2Cli:       hCli.V2(),
		ctx:         util.TODOContext(),
		httpCli:     util.DefaultHTTPClient(),
		httpTimeout: util.TimeSecond(600),
		pageConfig:  util.DefaultPage(),
	}, nil
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) WithHttpClient(httpCli *http.Client) *Client {
	c.httpCli = httpCli
	return c
}

func (c *Client) WithTimeout(httpTimeout time.Duration) *Client {
	c.httpTimeout = httpTimeout
	return c
}

func (c *Client) WithPageConfig(pCfg *util.Page) *Client {
	c.pageConfig = pCfg
	return c
}
