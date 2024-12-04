package fiberserver

import (
	"context"
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
		return _defaultLogger{}
	}
	return fiberlog.DefaultLogger()
}

type _defaultLogger struct{}

func (l _defaultLogger) Trace(v ...any)                          {}
func (l _defaultLogger) Debug(v ...any)                          {}
func (l _defaultLogger) Info(v ...any)                           {}
func (l _defaultLogger) Warn(v ...any)                           {}
func (l _defaultLogger) Error(v ...any)                          {}
func (l _defaultLogger) Fatal(v ...any)                          {}
func (l _defaultLogger) Panic(v ...any)                          {}
func (l _defaultLogger) Tracef(format string, v ...any)          {}
func (l _defaultLogger) Debugf(format string, v ...any)          {}
func (l _defaultLogger) Infof(format string, v ...any)           {}
func (l _defaultLogger) Warnf(format string, v ...any)           {}
func (l _defaultLogger) Errorf(format string, v ...any)          {}
func (l _defaultLogger) Fatalf(format string, v ...any)          {}
func (l _defaultLogger) Panicf(format string, v ...any)          {}
func (l _defaultLogger) Tracew(msg string, keysAndValues ...any) {}
func (l _defaultLogger) Debugw(msg string, keysAndValues ...any) {}
func (l _defaultLogger) Infow(msg string, keysAndValues ...any)  {}
func (l _defaultLogger) Warnw(msg string, keysAndValues ...any)  {}
func (l _defaultLogger) Errorw(msg string, keysAndValues ...any) {}
func (l _defaultLogger) Fatalw(msg string, keysAndValues ...any) {}
func (l _defaultLogger) Panicw(msg string, keysAndValues ...any) {}
func (l _defaultLogger) SetLevel(level fiberlog.Level)           {}
func (l _defaultLogger) SetOutput(w io.Writer)                   {}
func (l _defaultLogger) WithContext(ctx context.Context) fiberlog.CommonLogger {
	return _defaultLogger{}
}
