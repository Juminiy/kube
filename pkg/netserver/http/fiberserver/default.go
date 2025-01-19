package fiberserver

import (
	"context"
	"encoding/xml"
	"github.com/Juminiy/kube/pkg/log_api/zaplog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/gofiber/fiber/v3"
	fiberlog "github.com/gofiber/fiber/v3/log"
	"io"
	"net/http"
)

func DefaultRESTAPI(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("OK")
	})
}

func DefaultLogger() fiberlog.AllLogger {
	if len(_simpleCfg.logFilePath) > 0 {
		zaplog.New().
			WithLogEngineSugar().
			WithErrorOutputPaths(_simpleCfg.logFilePath).
			WithOutputPaths(_simpleCfg.logFilePath).
			WithLogCaller(true).
			WithLogStackTrace(true).
			WithLogLevel("info").Load()
		return _defaultLogger{}
	}
	return fiberlog.DefaultLogger()
}

type _defaultLogger struct{}

func (l _defaultLogger) Trace(v ...any)                 { zaplog.Debug(v...) }
func (l _defaultLogger) Debug(v ...any)                 { zaplog.Debug(v...) }
func (l _defaultLogger) Info(v ...any)                  { zaplog.Info(v...) }
func (l _defaultLogger) Warn(v ...any)                  { zaplog.Warn(v...) }
func (l _defaultLogger) Error(v ...any)                 { zaplog.Error(v...) }
func (l _defaultLogger) Fatal(v ...any)                 { zaplog.Fatal(v...) }
func (l _defaultLogger) Panic(v ...any)                 { zaplog.Panic(v...) }
func (l _defaultLogger) Tracef(format string, v ...any) { zaplog.DebugF(format, v...) }
func (l _defaultLogger) Debugf(format string, v ...any) { zaplog.DebugF(format, v...) }
func (l _defaultLogger) Infof(format string, v ...any)  { zaplog.InfoF(format, v...) }
func (l _defaultLogger) Warnf(format string, v ...any)  { zaplog.WarnF(format, v...) }
func (l _defaultLogger) Errorf(format string, v ...any) { zaplog.ErrorF(format, v...) }
func (l _defaultLogger) Fatalf(format string, v ...any) { zaplog.FatalF(format, v...) }
func (l _defaultLogger) Panicf(format string, v ...any) { zaplog.PanicF(format, v...) }
func (l _defaultLogger) Tracew(msg string, v ...any)    { zaplog.DebugW(msg, v...) }
func (l _defaultLogger) Debugw(msg string, v ...any)    { zaplog.DebugW(msg, v...) }
func (l _defaultLogger) Infow(msg string, v ...any)     { zaplog.InfoW(msg, v...) }
func (l _defaultLogger) Warnw(msg string, v ...any)     { zaplog.WarnW(msg, v...) }
func (l _defaultLogger) Errorw(msg string, v ...any)    { zaplog.ErrorW(msg, v...) }
func (l _defaultLogger) Fatalw(msg string, v ...any)    { zaplog.FatalW(msg, v...) }
func (l _defaultLogger) Panicw(msg string, v ...any)    { zaplog.PanicW(msg, v...) }
func (l _defaultLogger) SetLevel(level fiberlog.Level)  {}
func (l _defaultLogger) SetOutput(w io.Writer)          {}
func (l _defaultLogger) WithContext(ctx context.Context) fiberlog.CommonLogger {
	return _defaultLogger{}
}

func Minimal(restapi ...func(*fiber.App)) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:                 "",
		StrictRouting:                true,
		CaseSensitive:                true,
		Immutable:                    true,
		UnescapePath:                 true,
		BodyLimit:                    util.Mi,
		Concurrency:                  util.M,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  util.TimeSecond(5),
		WriteTimeout:                 util.TimeSecond(5),
		IdleTimeout:                  util.TimeSecond(5),
		ReadBufferSize:               util.Mi,
		WriteBufferSize:              util.Mi,
		CompressedFileSuffixes:       nil,
		ProxyHeader:                  "",
		GETOnly:                      true,
		ErrorHandler:                 nil,
		DisableKeepalive:             true,
		DisableDefaultDate:           true,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		AppName:                      "minimal-app",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: true,
		ReduceMemoryUsage:            true,
		JSONEncoder:                  safe_json.GoCCY().Marshal,
		JSONDecoder:                  safe_json.GoCCY().Unmarshal,
		XMLEncoder:                   xml.Marshal,
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		EnableIPValidation:           false,
		ColorScheme:                  fiber.Colors{},
		StructValidator:              nil,
		RequestMethods:               nil,
		EnableSplittingOnParsers:     false,
	})
	if len(restapi) > 0 && restapi[0] != nil {
		restapi[0](app)
	} else {
		DefaultRESTAPI(app)
	}
	return app
}
