package fiberserver

import (
	"encoding/xml"
	"github.com/Juminiy/kube/pkg/netserver/http/stdserver"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/psutil"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/Juminiy/kube/pkg/util/safe_validator"
	"github.com/gofiber/fiber/v3"
	fiberlog "github.com/gofiber/fiber/v3/log"
)

var (
	_simpleCfg = struct {
		webAppName         string
		readTimeoutSec     int
		writeTimeoutSec    int
		idleTimeoutSec     int
		bodyLimitSize      int
		readBufferSize     int
		writeBufferSize    int
		listenPort         int
		listenNetwork      string
		certFilePath       string
		certKeyFilePath    string
		certClientFilePath string
		RESTAPIRouter      func(*fiber.App)
		logFilePath        string
		logger             fiberlog.AllLogger
	}{}
)

func Init() {
	app := fiber.New(fiber.Config{
		ServerHeader:                 "",
		StrictRouting:                true,
		CaseSensitive:                true,
		Immutable:                    false,
		UnescapePath:                 false,
		BodyLimit:                    _simpleCfg.bodyLimitSize,
		Concurrency:                  1 * util.M,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  util.TimeSecond(_simpleCfg.readTimeoutSec),
		WriteTimeout:                 util.TimeSecond(_simpleCfg.writeTimeoutSec),
		IdleTimeout:                  util.TimeSecond(_simpleCfg.idleTimeoutSec),
		ReadBufferSize:               _simpleCfg.readBufferSize,
		WriteBufferSize:              _simpleCfg.writeBufferSize,
		CompressedFileSuffixes:       nil,
		ProxyHeader:                  "",
		GETOnly:                      false,
		ErrorHandler:                 nil,
		DisableKeepalive:             false,
		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		AppName:                      _simpleCfg.webAppName,
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: true,
		ReduceMemoryUsage:            true,
		JSONEncoder:                  safe_json.SafeConfig().Marshal,
		JSONDecoder:                  safe_json.SafeConfig().Unmarshal,
		XMLEncoder:                   xml.Marshal,
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		EnableIPValidation:           false,
		ColorScheme:                  fiber.Colors{},
		StructValidator:              safe_validator.Strict(),
		RequestMethods:               stdserver.LawHTTPMethod,
		EnableSplittingOnParsers:     false,
	})

	if _simpleCfg.RESTAPIRouter == nil {
		_simpleCfg.RESTAPIRouter = DefaultRESTAPI
	}
	if _simpleCfg.logger == nil {
		_simpleCfg.logger = DefaultLogger()
	}

	// others: "", fiber.NetworkTCP4
	hasIPv6 := util.ElemIn(_simpleCfg.listenNetwork, fiber.NetworkTCP, fiber.NetworkTCP6)
	chooseIPFunc := util.IsIPv4
	if hasIPv6 {
		chooseIPFunc = psutil.AllAddr
	}

	hasTLS := len(_simpleCfg.certKeyFilePath) > 0 &&
		len(_simpleCfg.certKeyFilePath) > 0 &&
		len(_simpleCfg.certClientFilePath) > 0

	_simpleCfg.RESTAPIRouter(app)
	stdserver.ListenAndServeInfoF(hasTLS, _simpleCfg.listenPort, chooseIPFunc)
	// listen and serve
	if hasTLS {
		_simpleCfg.logger.Fatal(app.Listen(stdserver.AllIntfs(_simpleCfg.listenPort), fiber.ListenConfig{
			ListenerNetwork:       _simpleCfg.listenNetwork,
			CertFile:              _simpleCfg.certFilePath,
			CertKeyFile:           _simpleCfg.certKeyFilePath,
			CertClientFile:        _simpleCfg.certClientFilePath,
			GracefulContext:       nil,
			TLSConfigFunc:         nil,
			ListenerAddrFunc:      nil,
			BeforeServeFunc:       nil,
			DisableStartupMessage: false,
			EnablePrefork:         false,
			EnablePrintRoutes:     false,
			OnShutdownError:       nil,
			OnShutdownSuccess:     nil,
		}))
	}
	_simpleCfg.logger.Fatal(app.Listen(stdserver.AllIntfs(_simpleCfg.listenPort)))
}
