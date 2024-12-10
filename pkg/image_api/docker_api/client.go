// Package docker_api
package docker_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_registry"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
	"github.com/go-resty/resty/v2"
)

//go:generate go run codegen/codegen.go
type Client struct {
	cli      *dockercli.Client // official SDK
	ctx      context.Context
	page     *util.Page
	hostAddr string
	version  string

	apiClient *docker_client.Client // official API
	reg       docker_registry.Registry

	// Deprecated
	restyCli *resty.Client
	// Deprecated
	cache *clientCache
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

	cli := &Client{
		cli:      dCli,
		ctx:      util.TODOContext(),
		page:     util.DefaultPage(),
		hostAddr: util.TrimProto(hostURL),
		version:  version,
	}
	cli.apiClient = docker_client.New(hostURL, version).
		WithPage(*cli.page).
		WithContext(cli.ctx)
	return cli, nil
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	c.apiClient.WithContext(ctx)
	return c
}

func (c *Client) WithPage(page *util.Page) *Client {
	c.page = page
	c.apiClient.WithPage(*page)
	return c
}

func (c *Client) WithRegistryAuth(authConfig *registry.AuthConfig) *Client {
	//cacheToken := c.internalRegistryAuth(registryAuthConfig)
	//c.cache.setLatestAuth(registryAuthConfig, cacheToken)
	//
	//registryAuthConfig.IdentityToken, _ = c.registryAuth(registryAuthConfig)
	//var encodeAuthConfigErr error
	//c.xRegistryAuth, encodeAuthConfigErr = registry.EncodeAuthConfig(*registryAuthConfig)
	//if encodeAuthConfigErr != nil {
	//	stdlog.WarnF("encode auth config error: %s", encodeAuthConfigErr.Error())
	//}
	//
	//c.registryAddr = registryAuthConfig.ServerAddress
	//c.restyCli.SetAllowGetMethodPayload(true).
	//	SetBaseURL(util.URLWithHTTP(c.hostAddr)).
	//	SetScheme("http").
	//	SetTimeout(util.TimeSecond(60))
	c.reg = c.apiClient.WithRegistryAuth(*authConfig).GetRegistry()
	return c
}

func (c *Client) WithProject(project string) *Client {
	c.reg.WithProject(project)
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
