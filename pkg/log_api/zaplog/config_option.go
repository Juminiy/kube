// global config
package zaplog

import "sync"

const (
	logEngineStdlib   = "stdlib"
	logEngineZap      = "zap"
	logEngineZapSugar = "zap sugared"
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

func (o *ConfigOption) WithLogEngine(logEngine string) *ConfigOption {
	_logEngine = logEngine
	return o
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
// +required
func (o *ConfigOption) WithLogLevel(level string) *ConfigOption {
	_logLevel = level
	return o
}

// WithLogCaller
// +optional
func (o *ConfigOption) WithLogCaller(ok bool) *ConfigOption {
	_caller = ok
	return o
}

// WithLogStackTrace
// +optional
func (o *ConfigOption) WithLogStackTrace(ok bool) *ConfigOption {
	_stacktrace = ok
	return o
}

// WithOutputPaths
// +required
func (o *ConfigOption) WithOutputPaths(s ...string) *ConfigOption {
	_outputPaths = s
	return o
}

// WithErrorOutputPaths
// +optional
func (o *ConfigOption) WithErrorOutputPaths(s ...string) *ConfigOption {
	_errorOutputPaths = s
	return o
}
