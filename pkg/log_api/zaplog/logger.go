package zaplog

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger

func logFunc(level zapcore.Level, fn string, templateOrMessage string, v ...any) {
	useZap := _logEngine == logEngineZap && _logger != nil
	useSugar := _logEngine == logEngineZapSugar && _sugaredLogger != nil
	//useStdlib := _logEngine == logEngineStdlib || (!useZap && !useSugar)

	// zap sugar
	if useSugar || useZap {
		switch level {
		case zap.DebugLevel:
			switch fn {
			case "f":
				_sugaredLogger.Debugf(templateOrMessage, v...)
			case "w":
				_sugaredLogger.Debugw(templateOrMessage, v...)
			default:
				_sugaredLogger.Debug(v...)
			}
		case zap.InfoLevel:
			switch fn {
			case "f":
				_sugaredLogger.Infof(templateOrMessage, v...)
			case "w":
				_sugaredLogger.Infow(templateOrMessage, v...)
			default:
				_sugaredLogger.Info(v...)
			}
		case zap.WarnLevel:
			switch fn {
			case "f":
				_sugaredLogger.Warnf(templateOrMessage, v...)
			case "w":
				_sugaredLogger.Warnw(templateOrMessage, v...)
			default:
				_sugaredLogger.Warn(v...)
			}
		case zap.ErrorLevel:
			switch fn {
			case "f":
				_sugaredLogger.Errorf(templateOrMessage, v...)
			case "w":
				_sugaredLogger.Errorw(templateOrMessage, v...)
			default:
				_sugaredLogger.Error(v...)
			}
		case zap.DPanicLevel:
			switch fn {
			case "f":
				_sugaredLogger.DPanicf(templateOrMessage, v...)
			case "w":
				_sugaredLogger.DPanicw(templateOrMessage, v...)
			default:
				_sugaredLogger.DPanic(v...)
			}
		case zap.PanicLevel:
			switch fn {
			case "f":
				_sugaredLogger.Panicf(templateOrMessage, v...)
			case "w":
				_sugaredLogger.Panicw(templateOrMessage, v...)
			default:
				_sugaredLogger.Panic(v...)
			}
		case zap.FatalLevel:
			switch fn {
			case "f":
				_sugaredLogger.Fatalf(templateOrMessage, v...)
			case "w":
				_sugaredLogger.Fatalw(templateOrMessage, v...)
			default:
				_sugaredLogger.Fatal(v...)
			}
		}
		return
	}

	// default or set to stdlib
	if _logEngine != logEngineStdlib {
		stdlog.ErrorF("config log engine is: %s, but it seems wrong, use stdlib instead", _logEngine)
	}
	switch level {
	case zap.DebugLevel:
		switch fn {
		case "f":
			stdlog.DebugF(templateOrMessage, v...)
		case "w":
			stdlog.DebugW(templateOrMessage, v...)
		default:
			stdlog.Debug(v...)
		}
	case zap.InfoLevel:
		switch fn {
		case "f":
			stdlog.InfoF(templateOrMessage, v...)
		case "w":
			stdlog.InfoW(templateOrMessage, v...)
		default:
			stdlog.Info(v...)
		}
	case zap.WarnLevel:
		switch fn {
		case "f":
			stdlog.WarnF(templateOrMessage, v...)
		case "w":
			stdlog.WarnW(templateOrMessage, v...)
		default:
			stdlog.Warn(v...)
		}
	case zap.ErrorLevel:
		switch fn {
		case "f":
			stdlog.ErrorF(templateOrMessage, v...)
		case "w":
			stdlog.ErrorW(templateOrMessage, v...)
		default:
			stdlog.Error(v...)
		}
	case zap.DPanicLevel,
		zap.PanicLevel:
		switch fn {
		case "f":
			stdlog.PanicF(templateOrMessage, v...)
		case "w":
			stdlog.PanicW(templateOrMessage, v...)
		default:
			stdlog.Panic(v...)
		}
	case zap.FatalLevel:
		switch fn {
		case "f":
			stdlog.FatalF(templateOrMessage, v...)
		case "w":
			stdlog.FatalW(templateOrMessage, v...)
		default:
			stdlog.Fatal(v...)
		}
	}

}

func Debug(v ...any) {
	logFunc(zap.DebugLevel, "", "", v...)
}

func DebugF(format string, v ...any) {
	logFunc(zap.DebugLevel, "f", format, v...)
}

func DebugW(msg string, kv ...any) {
	logFunc(zap.DebugLevel, "w", msg, kv...)
}

func Info(v ...any) {
	logFunc(zap.InfoLevel, "", "", v...)
}

func InfoF(format string, v ...any) {
	logFunc(zap.InfoLevel, "f", format, v...)
}

func InfoW(msg string, kv ...any) {
	logFunc(zap.InfoLevel, "w", msg, kv...)
}

func Warn(v ...any) {
	logFunc(zap.WarnLevel, "", "", v...)
}

func WarnF(format string, v ...any) {
	logFunc(zap.WarnLevel, "f", format, v...)
}

func WarnW(msg string, kv ...any) {
	logFunc(zap.WarnLevel, "w", msg, kv...)
}

func Error(v ...any) {
	logFunc(zap.ErrorLevel, "", "", v...)
}

func ErrorF(format string, v ...any) {
	logFunc(zap.ErrorLevel, "f", format, v...)
}

func ErrorW(msg string, kv ...any) {
	logFunc(zap.ErrorLevel, "w", msg, kv...)
}

func Fatal(v ...any) {
	logFunc(zap.FatalLevel, "", "", v...)
}

func FatalF(format string, v ...any) {
	logFunc(zap.FatalLevel, "f", format, v...)
}

func FatalW(msg string, kv ...any) {
	logFunc(zap.FatalLevel, "w", msg, kv...)
}

func Panic(v ...any) {
	logFunc(zap.PanicLevel, "", "", v...)
}

func PanicF(format string, v ...any) {
	logFunc(zap.PanicLevel, "f", format, v...)
}

func PanicW(msg string, kv ...any) {
	logFunc(zap.PanicLevel, "w", msg, kv...)
}

func DPanic(v ...any) {
	logFunc(zap.DPanicLevel, "", "", v...)
}

func DPanicF(format string, v ...any) {
	logFunc(zap.DPanicLevel, "f", format, v...)
}

func DPanicW(msg string, kv ...any) {
	logFunc(zap.DPanicLevel, "w", msg, kv...)
}
