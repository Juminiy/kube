// global config
package api_inst

import (
	kubeapi "github.com/Juminiy/kube/pkg/api"
	"github.com/Juminiy/kube/pkg/internal"
	"sync"
)

type ConfigOption struct {
	_      internal.NoCmp
	noCopy internal.NoCopy
	sync.Once
}

func (o *ConfigOption) Load() {
	o.Do(Init)
}

func (o *ConfigOption) WithLogger(levelLogger kubeapi.LevelLogger) *ConfigOption {
	_logger = levelLogger
	return o
}

func (o *ConfigOption) WithInternalLogger(iLogger kubeapi.InternalLogger) *ConfigOption {
	_iLogger = iLogger
	return o
}

func (o *ConfigOption) WithInternalLoggerV2(iLoggerV2 kubeapi.InternalLoggerV2) *ConfigOption {
	_iLoggerV2 = iLoggerV2
	return o
}
