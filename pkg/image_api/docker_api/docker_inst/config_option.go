package docker_inst

import (
	"context"
	"sync"
)

type ConfigOption struct {
	_ struct{}
	sync.Once
}

func NewConfig() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() *ConfigOption {
	o.Do(Init)
	return o
}

func (o *ConfigOption) WithHost(host string) *ConfigOption {
	_hostURL = host
	return o
}

func (o *ConfigOption) WithVersion(version string) *ConfigOption {
	_version = version
	return o
}

func (o *ConfigOption) WithContext(ctx context.Context) *ConfigOption {
	_docketContext = ctx
	return o
}
