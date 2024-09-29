package stdlog

import "sync"

type ConfigOption struct {
	_ struct{}
	sync.Once
}

func New() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() {
	o.Do(Init)
}

func (o *ConfigOption) WithLogPath(logPath string) *ConfigOption {
	_logPath = logPath
	return o
}

func (o *ConfigOption) WithTimeMicroSeconds() *ConfigOption {
	_logTimeMicroSeconds = true
	return o
}

func (o *ConfigOption) WithCallerLongFile() *ConfigOption {
	_logCallerLongFile = true
	return o
}

func (o *ConfigOption) WithCallerShortFile() *ConfigOption {
	_logCallerShortFile = true
	return o
}

func (o *ConfigOption) WithTimeUTC() *ConfigOption {
	_logTimeUTC = true
	return o
}
