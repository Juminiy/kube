// global var
package zaplog

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"slices"
)

// LogLevelMap is constant immutable variable
var (
	_logLevelMap = map[string]zapcore.Level{
		"DEBUG": zap.DebugLevel, "Debug": zap.DebugLevel, "debug": zap.DebugLevel,
		"INFO": zap.InfoLevel, "Info": zap.InfoLevel, "info": zap.InfoLevel,
		"WARN": zap.WarnLevel, "Warn": zap.WarnLevel, "warn": zap.WarnLevel,
		"ERROR": zap.ErrorLevel, "Error": zap.ErrorLevel, "error": zap.ErrorLevel,
		"DPANIC": zap.DPanicLevel, "DPanic": zap.DPanicLevel, "dpanic": zap.DPanicLevel,
		"PANIC": zap.PanicLevel, "Panic": zap.PanicLevel, "panic": zap.PanicLevel,
		"FATAL": zap.FatalLevel, "Fatal": zap.FatalLevel, "fatal": zap.FatalLevel,
	}
)

// global config
var (
	// default sugar, optional: {"stdlib", "zap", "zap sugared"}
	_logEngine string

	// runtime dynamic log level, can change by RESTAPI
	_logLevel string

	// show log json caller
	_caller bool

	// show log json stacktrace
	_stacktrace bool

	// output path list
	_outputPaths      []string
	_errorOutputPaths []string
)

// global var
var (
	_zapConfig     *zap.Config
	_restoreFunc   = util.NothingFn()
	_logger        *zap.Logger
	_sugaredLogger *zap.SugaredLogger
)

func Init() {
	var _zapError error

	checkPaths := func(paths ...string) (legal int) {
		slices.Values(paths)(func(path string) bool {
			if util.OSFilePathExists(path) {
				legal++
				return true
			}
			if err := util.OSCreateAbsolutePath(path); err == nil {
				legal++
			} else {
				stdlog.Error(err)
			}
			return true
		})
		return legal
	}
	legalOutput := checkPaths(_outputPaths...)
	legalErrOutput := checkPaths(_errorOutputPaths...)
	if legalOutput == 0 &&
		legalErrOutput == 0 {
		stdlog.Error("output path and error output path must have one available absolute path at least")
		return
	} else if legalOutput == 0 {
		_outputPaths = _errorOutputPaths
	} else if legalErrOutput == 0 {
		_errorOutputPaths = _outputPaths
	}

	_zapConfig = util.New(zap.NewProductionConfig())
	func(lvl string) {
		logLevelInt8, ok := _logLevelMap[lvl]
		if !ok {
			stdlog.WarnF("in config file log.level: %s is not supported, please use: Debug, Info, Warn, Error, DPanic, Panic, Fatal", lvl)
			logLevelInt8 = zap.InfoLevel
		}
		_zapConfig.Level = zap.NewAtomicLevelAt(logLevelInt8)
	}(_logLevel)
	_zapConfig.DisableCaller = !_caller
	_zapConfig.DisableStacktrace = !_stacktrace
	_zapConfig.OutputPaths = _outputPaths
	_zapConfig.ErrorOutputPaths = _errorOutputPaths
	_zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	_zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	_logger, _zapError = _zapConfig.Build()
	if _zapError != nil {
		stdlog.ErrorF("zap config init error: %s", _zapError.Error())
		return
	}
	_sugaredLogger = _logger.Sugar()

	_restoreFunc = zap.ReplaceGlobals(_logger)
}

func Get() *zap.Logger {
	return _logger
}

func GetSugar() *zap.SugaredLogger {
	return _sugaredLogger
}
