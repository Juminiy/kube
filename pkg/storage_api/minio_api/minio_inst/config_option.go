package minio_inst

import "sync"

type ConfigOption struct {
	_ struct{}
	sync.Once
}

func NewConfig() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() *ConfigOption {
	o.Do(Init)
	return o
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

func (o *ConfigOption) WithPublicBucket(publicBucket string) *ConfigOption {
	_minioPublicBucket = publicBucket
	return o
}
