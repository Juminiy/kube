// global config
package stdserver

import (
	"github.com/Juminiy/kube/pkg/internal"
	"log"
	"net/http"
	"sync"
)

type ConfigOption struct {
	_      internal.NoCmp
	noCopy internal.NoCopy
	sync.Once
}

func New() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() {
	o.Do(Init)
}

// WithHost
// referred to net.Dial Examples
func (o *ConfigOption) WithHost(host string) *ConfigOption {
	_host = host
	return o
}

func (o *ConfigOption) WithPort(port int) *ConfigOption {
	_port = port
	return o
}

func (o *ConfigOption) WithTimeoutSec(read, write int) *ConfigOption {
	_readTimeoutSec = read
	_writeTimeoutSec = write
	return o
}

func (o *ConfigOption) WithMaxHeaderBytes(b int) *ConfigOption {
	_maxHeaderBytes = b
	return o
}

func (o *ConfigOption) WithTLS(certFilePath, keyFilePath string) *ConfigOption {
	_tls = true
	_tlsCertFilePath = certFilePath
	_tlsKeyFilePath = keyFilePath
	return o
}

func (o *ConfigOption) WithLogger(logger *log.Logger) *ConfigOption {
	_logger = logger
	return o
}

func (o *ConfigOption) WithHandler(handler http.Handler) *ConfigOption {
	_handler = handler
	return o
}
