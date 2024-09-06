package zaplog

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

// describe in config file
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

// immutable global variable
var (
	_zapConfig   *zap.Config
	_restoreFunc = util.NothingFn
)

func Init() {
	var _zapError error

	/*if len(_outputPaths) == 0 &&
		len(_errorOutputPaths) == 0 {
		stdlog.Error("output path and error output path must have one available absolute path at least")
		return
	} else if len(_errorOutputPaths) == 0 {
		_errorOutputPaths = _outputPaths
	} else if len(_outputPaths) == 0 {
		_outputPaths = _errorOutputPaths
	}*/

	checkOutputPath := func(outputPaths ...string) int {
		cnt := 0
		for _, outputPath := range outputPaths {
			if !util.OSFilePathExists(outputPath) {
				_zapError = util.OSCreateAbsolutePath(outputPath)
				if _zapError != nil {
					stdlog.Error(_zapError)
					continue
				}
				cnt++
				continue
			}
			cnt++
		}
		return cnt
	}
	outputPathCnt := checkOutputPath(_outputPaths...)
	errorOutputPathCnt := checkOutputPath(_errorOutputPaths...)

	if outputPathCnt == 0 &&
		errorOutputPathCnt == 0 {
		stdlog.Error("output path and error output path must have one available absolute path at least")
		return
	} else if outputPathCnt == 0 {
		_outputPaths = _errorOutputPaths
	} else if errorOutputPathCnt == 0 {
		_errorOutputPaths = _outputPaths
	}

	_zapConfig = util.New(zap.NewProductionConfig())
	_zapConfig.DisableCaller = !_caller
	_zapConfig.DisableStacktrace = !_stacktrace
	setLogLevel(_logLevel)
	setOutputPaths(_outputPaths, _errorOutputPaths)
	_zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	_zapConfig.EncoderConfig.EncodeCaller = func(_ zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("")
	}

	_logger, _zapError = _zapConfig.Build()
	if _zapError != nil {
		stdlog.ErrorF("zap config init error: %s", _zapError.Error())
		return
	}
	_sugaredLogger = _logger.Sugar()

	_restoreFunc = zap.ReplaceGlobals(_logger)
}

// will be exported later
func setLogLevel(level string) {
	logLevelInt8, ok := _logLevelMap[level]
	if !ok {
		stdlog.ErrorF("in config file log.level: %s is not supported, please use: Debug, Info, Warn, Error, DPanic, Panic, Fatal", level)
		return
	}
	_zapConfig.Level = zap.NewAtomicLevelAt(logLevelInt8)
}

// will be exported later
func setOutputPaths(paths, errorPaths []string) {
	_zapConfig.OutputPaths = paths
	_zapConfig.ErrorOutputPaths = errorPaths
}
