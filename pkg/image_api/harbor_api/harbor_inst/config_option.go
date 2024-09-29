package harbor_inst

import (
	"github.com/Juminiy/kube/pkg/util"
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

func (o *ConfigOption) WithRegistry(harborRegistry string) *ConfigOption {
	_harborRegistry = harborRegistry
	_harborInsecure = util.IsURLWithHTTP(harborRegistry)
	return o
}

func (o *ConfigOption) WithInsecure() *ConfigOption {
	_harborInsecure = true
	return o
}

func (o *ConfigOption) WithSecure() *ConfigOption {
	_harborInsecure = false
	return o
}

func (o *ConfigOption) WithUsername(username string) *ConfigOption {
	_harborUsername = username
	return o
}

func (o *ConfigOption) WithPassword(password string) *ConfigOption {
	_harborPassword = password
	return o
}
