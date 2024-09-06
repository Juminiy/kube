package minio_inst

type ConfigOption struct {
	_ struct{}
}

func (o *ConfigOption) WithEndpoint(endpoint string) *ConfigOption {
	_minioEndpoint = endpoint
	return o
}

func (o *ConfigOption) WithAccessKeyID(accessKeyID string) *ConfigOption {
	_minioAccessKeyID = accessKeyID
	return o
}

func (o *ConfigOption) WithSecretAccessKey(secretAccessKey string) *ConfigOption {
	_minioSecretAccessKey = secretAccessKey
	return o
}

func (o *ConfigOption) WithSessionToken(sessionToken string) *ConfigOption {
	_minioSessionToken = sessionToken
	return o
}

func (o *ConfigOption) WithSecure() *ConfigOption {
	_minioSecure = true
	return o
}
