// global config
package docker_inst

import (
	"context"
	"sync"
)

type ConfigOption struct {
	_ struct{}
	sync.Once
}

// NewConfig
// Deprecated, use New instead
func NewConfig() *ConfigOption {
	return &ConfigOption{}
}

func New() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() {
	o.Do(Init)
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
