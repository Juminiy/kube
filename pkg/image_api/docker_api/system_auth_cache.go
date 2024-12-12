// Deprecated
package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	"sync"
)

const (
	initCacheSize = 8

	registryAuthIdentityTokenKey = "RegistryAuthIdentityToken"
	registryAuthConfigKey        = "RegistryAuthConfig"
	registryLatestAuthConfigKey  = "RegistryAuthConfig:latest"
)

type clientCache struct {
	m sync.Map
}

func newClientCache() *clientCache {
	return &clientCache{m: sync.Map{}}
}

func (c *clientCache) setAuthIdentityToken(config *registry.AuthConfig, registryAuthIdentityToken string) {
	c.m.Store(getAuthConfigKey(config), config)
	c.m.Store(getAuthTokenKey(config), registryAuthIdentityToken)
}

func (c *clientCache) getAuthIdentityToken(config *registry.AuthConfig) string {
	c.m.Store(getAuthConfigKey(config), config)
	val, ok := c.m.Load(getAuthTokenKey(config))
	if ok {
		return val.(string)
	}
	return ""
}

func (c *clientCache) setLatestAuth(config *registry.AuthConfig, identityToken string) {
	c.m.Store(registryLatestAuthConfigKey, config)
	c.setAuthIdentityToken(config, identityToken)
}

func (c *clientCache) getLatestAuthConfig() *registry.AuthConfig {
	val, ok := c.m.Load(registryLatestAuthConfigKey)
	if ok {
		return val.(*registry.AuthConfig)
	}
	return nil
}

func (c *clientCache) getLatestAuthIdentityToken() string {
	latestAuthConfig := c.getLatestAuthConfig()
	if latestAuthConfig != nil {
		return c.getAuthIdentityToken(latestAuthConfig)
	}
	return ""
}

func getAuthTokenKey(config *registry.AuthConfig) string {
	if config == nil {
		return ""
	}
	return util.StringJoin("-",
		registryAuthIdentityTokenKey,
		config.ServerAddress,
		config.Username,
		config.Password)
}

func getAuthConfigKey(config *registry.AuthConfig) string {
	if config == nil {
		return ""
	}
	return util.StringJoin("-",
		registryAuthConfigKey,
		config.ServerAddress,
		config.Username,
		config.Password,
	)
}
