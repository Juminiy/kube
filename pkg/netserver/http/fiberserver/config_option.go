package fiberserver

import (
	"github.com/Juminiy/kube/pkg/internal"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"github.com/gofiber/fiber/v3"
	fiberlog "github.com/gofiber/fiber/v3/log"
	"k8s.io/apimachinery/pkg/api/resource"
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

func (o *ConfigOption) WithName(appName string) *ConfigOption {
	_simpleCfg.webAppName = appName
	return o
}

func (o *ConfigOption) WithTimeout(readSec, writeSec, idleSec int) *ConfigOption {
	_simpleCfg.readTimeoutSec = readSec
	_simpleCfg.writeTimeoutSec = writeSec
	_simpleCfg.idleTimeoutSec = idleSec
	return o
}

func (o *ConfigOption) WithBufferSize(bodySz, readBufSz, writeBufSz string) *ConfigOption {
	_simpleCfg.bodyLimitSize = parseSize(bodySz)
	_simpleCfg.readBufferSize = parseSize(readBufSz)
	_simpleCfg.writeBufferSize = parseSize(writeBufSz)
	return o
}

func parseSize(sz string) int {
	parsed := resource.MustParse(sz)
	return safe_cast.I64toI(parsed.Value())
}

func (o *ConfigOption) WithPort(port int) *ConfigOption {
	_simpleCfg.listenPort = port
	return o
}

func (o *ConfigOption) WithListenIPv6() *ConfigOption {
	_simpleCfg.listenNetwork = fiber.NetworkTCP
	return o
}

func (o *ConfigOption) WithRESTAPI(RESTAPI func(*fiber.App)) *ConfigOption {
	_simpleCfg.RESTAPIRouter = RESTAPI
	return o
}

func (o *ConfigOption) WithLogger(logger fiberlog.AllLogger) *ConfigOption {
	_simpleCfg.logger = logger
	return o
}
