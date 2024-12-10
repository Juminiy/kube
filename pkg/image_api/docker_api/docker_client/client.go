package docker_client

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_registry"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	"github.com/go-resty/resty/v2"
	"strings"
)

type Client struct {
	rCli *resty.Client   // HTTP API client
	page *util.Page      // for list pagination
	ctx  context.Context // for task cancellation

	hostAddr string                   // docker host address, ex. ip:port, domain:port
	version  string                   // docker client version, ex. 1.27, v1.27
	reg      docker_registry.Registry // docker registry, for pull, push, tag, etc
}

func New(host, version string) *Client {
	hostAddr := util.TrimProto(host)
	restyCli := resty.NewWithClient(util.DefaultHTTPClient()).
		SetAllowGetMethodPayload(true).
		SetBaseURL(util.URLWithHTTP(hostAddr)).
		SetScheme("http").
		SetTimeout(util.TimeSecond(60))
	return &Client{
		rCli:     restyCli,
		hostAddr: hostAddr,
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

func (c *Client) WithRegistryAuth(authConfig registry.AuthConfig) *Client {
	authResp, err := c.RegistryLogin(authConfig)
	if err != nil {
		stdlog.ErrorF("docker registry login error: %s", err.Error())
	} else {
		authConfig.IdentityToken = authResp.IdentityToken
	}
	c.reg = docker_registry.FromAuthConfig(authConfig)
	if err != nil {
		stdlog.ErrorF("docker registry encode authConfig error: %s", err.Error())
	}
	return c
}

func (c *Client) GetRegistry() docker_registry.Registry {
	return c.reg
}

func versionWithV(version string) string {
	if strings.HasPrefix(version, "v") {
		return version
	}
	return "v" + version
}
