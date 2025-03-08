// Package docker_api
package docker_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/api_provider"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_registry"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
	"github.com/go-resty/resty/v2"
	"net/http"
)

//go:generate go run codegen/codegen.go
type Client struct {
	cli     *dockercli.Client // official SDK
	ctx     context.Context
	page    *util.Page
	host    string
	version string

	apiClient *docker_client.Client // official API
	reg       docker_registry.Registry

	// Deprecated
	restyCli *resty.Client
	// Deprecated
	cache *clientCache
}

func New(host, version string) (*Client, error) {
	dCli, err := dockercli.NewClientWithOpts(
		dockercli.WithHost(host),
		dockercli.WithVersion(version),
	)
	if err != nil {
		stdlog.ErrorF("connect to docker host: %s error: %s", host, err.Error())
		return nil, err
	}

	cli := &Client{
		cli:     dCli,
		ctx:     util.TODOContext(),
		page:    util.DefaultPage(),
		host:    host,
		version: version,
	}
	cli.apiClient = docker_client.New(host, version).
		WithPage(*cli.page).
		WithContext(cli.ctx)
	return cli, nil
}

func NewWithOpts(
	host,
	version string,
	client *http.Client,
	opt ...dockercli.Opt,
) (*Client, error) {
	dCli, err := dockercli.NewClientWithOpts(
		append(opt,
			dockercli.WithHost(host),
			dockercli.WithVersion(version),
			dockercli.WithHTTPClient(client),
		)...,
	)
	if err != nil {
		return nil, err
	}

	cli := &Client{
		cli:     dCli,
		ctx:     util.TODOContext(),
		page:    util.DefaultPage(),
		host:    host,
		version: version,
	}
	cli.apiClient = docker_client.New(host, version).
		WithPage(*cli.page).
		WithContext(cli.ctx).
		WithHTTPClient(client)
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

func (c *Client) GetAPIProvider() api_provider.APIProvider {
	return api_provider.APIProvider{
		SDK: c.cli,
		API: c.apiClient,
	}
}

func (c *Client) Do(fc func(cli *dockercli.Client) error) error {
	return fc(c.cli)
}
