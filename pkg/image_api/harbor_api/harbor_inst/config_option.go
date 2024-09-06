package harbor_inst

type ConfigOption struct {
	_ struct{}
}

func NewConfig() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() *ConfigOption {
	Init()
	return o
}

func (o *ConfigOption) WithURL(harborURL string) *ConfigOption {
	_harborRegistry = harborURL
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
