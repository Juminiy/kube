// global var
package api_inst

import (
	kubeapi "github.com/Juminiy/kube/pkg/api"
)

// global config
var ()

// global var
var (
	_logger    kubeapi.LevelLogger
	_iLogger   kubeapi.InternalLogger
	_iLoggerV2 kubeapi.InternalLoggerV2
)

func Init() {}
