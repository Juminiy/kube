package zaplog

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
)

const (
	logEngineStdlib   = "stdlib"
	logEngineZap      = "zap"
	logEngineZapSugar = "zap sugared"
)

type ConfigOption struct {
	_ struct{}
}

func NewConfig() *ConfigOption {
	return &ConfigOption{}
}

func (*ConfigOption) Load() {
	stdlog.Debug(_sugaredLogger)
	util.NothingFn()
}

// WithLogEngineStd
// +optional
func (o *ConfigOption) WithLogEngineStd() *ConfigOption {
	_logEngine = logEngineStdlib
	return o
}

// WithLogEngineZap
// +optional
func (o *ConfigOption) WithLogEngineZap() *ConfigOption {
	_logEngine = logEngineZap
	return o
}

// WithLogEngineSugar
// +optional
func (o *ConfigOption) WithLogEngineSugar() *ConfigOption {
	_logEngine = logEngineZapSugar
	return o
}

// WithLogLevel
// +optional
func (o *ConfigOption) WithLogLevel(level string) *ConfigOption {
	_logLevel = level
	return o
}

// WithLogCaller
// +optional
func (o *ConfigOption) WithLogCaller() *ConfigOption {
	_caller = true
	return o
}

// WithLogStackTrace
// +optional
func (o *ConfigOption) WithLogStackTrace() *ConfigOption {
	_stacktrace = true
	return o
}

// WithOutputPaths
// +optional
func (o *ConfigOption) WithOutputPaths(s ...string) *ConfigOption {
	_outputPaths = s
	_errorOutputPaths = s
	return o
}

// WithErrorOutputPaths
// +optional
func (o *ConfigOption) WithErrorOutputPaths(s ...string) *ConfigOption {
	_errorOutputPaths = s
	return o
}
