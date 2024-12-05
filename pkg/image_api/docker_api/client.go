// Package docker_api
package docker_api

import (
	"context"
	"encoding/base64"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
)

type Client struct {
	cli        *dockercli.Client
	ctx        context.Context
	pageConfig *util.Page

	registryAddr string
	base64Auth   string
	cache        *clientCache
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
	return c.WithRegistryAuth2(registryAuthConfig.Username, registryAuthConfig.Password, registryAuthConfig.ServerAddress)
}

func (c *Client) WithRegistryAuth2(username, password, registryAddr string) *Client {
	c.registryAddr = registryAddr
	c.base64Auth = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
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
