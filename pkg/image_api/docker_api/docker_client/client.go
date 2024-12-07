package docker_client

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/go-resty/resty/v2"
	"strings"
)

type Client struct {
	rCli     *resty.Client
	hostAddr string // ip:port, domain:port
	version  string // 1.27, v1.27

	registryAddr  string     // ip:port, domain:port
	xRegistryAuth string     // base64 of username:password
	page          *util.Page // for list pagination
	ctx           context.Context
}

func New(host, version string) *Client {
	return &Client{
		rCli:     resty.NewWithClient(util.DefaultHTTPClient()),
		hostAddr: util.TrimProto(host),
		version:  versionWithV(version),
		page:     util.DefaultPage(),
		ctx:      util.TODOContext(),
	}
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) WithPage(page util.Page) *Client {
	c.page = &page
	return c
}

func versionWithV(version string) string {
	if strings.HasPrefix(version, "v") {
		return version
	}
	return "v" + version
}
