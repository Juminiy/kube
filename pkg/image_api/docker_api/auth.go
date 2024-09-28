package docker_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"github.com/docker/docker/api/types/registry"
)

func (c *Client) RegistryAuth(config *registry.AuthConfig) (string, error) {
	return c.registryAuth(config)
}

// it's funny that:https://github.com/moby/moby/issues/10983
// if registry do not return token,
// give an arbitrary value,
// make docker daemon happy,
// the registry Manufacturer like: harbor, aws-image-handler may not response you a token,
// but the docker daemon check it.
func (c *Client) registryAuth(config *registry.AuthConfig) (string, error) {
	if config == nil {
		stdlog.Error("registry auth config is nil")
		return "", nil
	}
	authResp, err := c.cli.RegistryLogin(c.ctx, *config)
	if err != nil {
		return "", err
	}
	return authResp.IdentityToken, nil
}

// MethodWithAuth
// try at least 3 times
// do not need anymore
func (c *Client) methodWithAuth(fn util.Fn) {

}

const (
	atLeastTryCount = 3
)

// try at least 3 times
func (c *Client) internalRegistryAuth(config *registry.AuthConfig) (cacheToken string) {
	if config == nil ||
		len(config.ServerAddress) == 0 ||
		len(config.Username) == 0 ||
		len(config.Password) == 0 {
		stdlog.ErrorF("registry auth config is not ok")
		return
	}

	cacheToken = c.cache.getAuthIdentityToken(config)
	if len(cacheToken) > 0 {
		return
	}

	tryFunc(
		func() bool {
			var err error
			cacheToken, err = c.registryAuth(config)
			if err != nil {
				stdlog.ErrorF("registry auth error: %s", err.Error())
				return false
			} else if len(cacheToken) == 0 {
				stdlog.Warn("registry auth is ok, but cacheToken is nil, set an arbitrary token value")
				cacheToken = random.Password()
			}
			return true
		},
		atLeastTryCount,
	)

	c.cache.setAuthIdentityToken(config, cacheToken)
	return
}

type breakFn func() bool

func tryFunc(fn breakFn, count int) {
	if fn == nil {
		return
	}
	for cnt := 0; cnt < count; cnt++ {
		if breakIf := fn(); breakIf {
			break
		}
	}
}
