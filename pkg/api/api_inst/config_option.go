package api_inst

import (
	kubeapi "github.com/Juminiy/kube/pkg/api"
	"sync"
)

type ConfigOption struct {
	_ struct{}
	sync.Once
}

func (o *ConfigOption) Load() {
	o.Do(Init)
}

func (o *ConfigOption) WithLogger(levelLogger kubeapi.LevelLogger) *ConfigOption {
	_logger = levelLogger
	return o
}
