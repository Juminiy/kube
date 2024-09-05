package docker_inst

import "context"

type ConfigOption struct {
	_ struct{}
}

func NewConfig() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) WithHost(host string) *ConfigOption {
	return o
}

func (o *ConfigOption) WithVersion(version string) *ConfigOption {
	return o
}

func (o *ConfigOption) WithContext(ctx context.Context) *ConfigOption {
	return o
}
