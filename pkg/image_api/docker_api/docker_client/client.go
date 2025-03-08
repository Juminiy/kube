package docker_client

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_registry"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
	"github.com/dubbogo/gost/log/logger"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
)

type Client struct {
	rCli *resty.Client   // HTTP API client
	page *util.Page      // for list pagination
	ctx  context.Context // for task cancellation

	// referred from: dockercli.Client
	rawHost    string
	rawVersion string
	rest       *restConfig

	// docker registry, for pull, push, tag, etc
	reg docker_registry.Registry
}

func New(host, version string) *Client {
	return (&Client{
		page: util.DefaultPage(),
		ctx:  util.TODOContext(),

		rawHost:    host,
		rawVersion: version,
	}).WithHTTPClient(util.DefaultHTTPClient())
}

func (c *Client) WithHTTPClient(client *http.Client) *Client {
	rCli, rest, err := restyClientWithHTTPClient(c.rawHost, client)
	if err != nil {
		logger.Errorf("set resty with http.Client error: %s", err.Error())
		return c
	}
	rest.version = versionWithV(c.rawVersion)
	c.rCli = rCli
	c.rest = rest
	return c
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

type restConfig struct {
	scheme   string
	host     string
	proto    string
	addr     string
	basePath string
	version  string
}

func restyClientWithHTTPClient(host string, client *http.Client) (*resty.Client, *restConfig, error) {
	hostURL, err := dockercli.ParseHostURL(host)
	if err != nil {
		return nil, nil, err
	}
	if hostURL.Scheme == "tcp" {
		hostURL.Scheme = "http"
	}
	return resty.NewWithClient(client).
			SetAllowGetMethodPayload(true).
			SetBaseURL(hostURL.String()).
			SetScheme(hostURL.Scheme).
			SetTimeout(client.Timeout),
		&restConfig{
			scheme:   hostURL.Scheme,
			host:     host,
			proto:    hostURL.Scheme,
			addr:     hostURL.Host,
			basePath: hostURL.Path,
		}, nil
}

func versionWithV(version string) string {
	if strings.HasPrefix(version, "v") {
		return version
	}
	return "v" + version
}
