// global config
package minio_inst

import (
	"github.com/Juminiy/kube/pkg/internal"
	"sync"
)

type ConfigOption struct {
	_      internal.NoCmp
	noCopy internal.NoCopy
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
