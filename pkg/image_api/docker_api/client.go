// Package docker_api
package docker_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	cli        *dockercli.Client
	ctx        context.Context
	pageConfig *util.Page
	hostAddr   string
	version    string

	registryAddr  string
	xRegistryAuth string
	restyCli      *resty.Client
	cache         *clientCache
}

func New(hostURL, version string) (*Client, error) {
	dCli, err := dockercli.NewClientWithOpts(
		dockercli.WithHost(hostURL),
		dockercli.WithVersion(version),
	)
	if err != nil {
		stdlog.ErrorF("connect to docker host: %s error: %s", hostURL, err.Error())
		return nil, err
	}

	return &Client{
		cli:        dCli,
		ctx:        util.TODOContext(),
		pageConfig: util.DefaultPage(),
		hostAddr:   util.TrimProto(hostURL),
		version:    version,
		restyCli:   resty.NewWithClient(util.DefaultHTTPClient()),
		cache:      newClientCache(),
	}, nil
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) WithPage(page *util.Page) *Client {
	c.pageConfig = page
	return c
}

func (c *Client) WithRegistryAuth(registryAuthConfig *registry.AuthConfig) *Client {
	cacheToken := c.internalRegistryAuth(registryAuthConfig)
	c.cache.setLatestAuth(registryAuthConfig, cacheToken)

	registryAuthConfig.IdentityToken = cacheToken
	var encodeAuthConfigErr error
	c.xRegistryAuth, encodeAuthConfigErr = registry.EncodeAuthConfig(*registryAuthConfig)
	if encodeAuthConfigErr != nil {
		stdlog.WarnF("encode auth config error: %s", encodeAuthConfigErr.Error())
	}

	c.registryAddr = registryAuthConfig.ServerAddress
	c.restyCli.SetAllowGetMethodPayload(true).
		SetBaseURL(util.URLWithHTTP(c.hostAddr)).
		SetScheme("http").
		SetTimeout(util.TimeSecond(60))
	return c
}

func (c *Client) GC(gcFn ...util.Func) {}

type ClientFunc func(client *dockercli.Client) error

func (c *Client) Do(cfn ...ClientFunc) error {
	for i := range cfn {
		if err := cfn[i](c.cli); err != nil {
			return err
		}
	}
	return nil
}
