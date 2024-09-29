package stdlog

import (
	"strings"
)

// Debug color: purple
func Debug(v ...any) {
	outputLine(logLevelDebug, v...)
}

func DebugF(format string, v ...any) {
	outputFormatLine(logLevelDebug, format, v...)
}

func DebugW(message string, kv ...any) {
	outputKeyValueLine(logLevelDebug, message, kv...)
}

// Info color: green
func Info(v ...any) {
	outputLine(logLevelInfo, v...)
}

func InfoF(format string, v ...any) {
	outputFormatLine(logLevelInfo, format, v...)
}

func InfoW(message string, kv ...any) {
	outputKeyValueLine(logLevelInfo, message, kv...)
}

// Warn color: yellow
func Warn(v ...any) {
	outputLine(logLevelWarn, v...)
}

func WarnF(format string, v ...any) {
	outputFormatLine(logLevelWarn, format, v...)
}

func WarnW(message string, kv ...any) {
	outputKeyValueLine(logLevelWarn, message, kv...)
}

// Error color: red
func Error(v ...any) {
	outputLine(logLevelError, v...)
}

func ErrorF(format string, v ...any) {
	outputFormatLine(logLevelError, format, v...)
}

func ErrorW(message string, kv ...any) {
	outputKeyValueLine(logLevelError, message, kv...)
}

// Fatal color: red
func Fatal(v ...any) {
	_logger.SetPrefix(logLevelFatal)
	_logger.Fatalln(v...)
}

func FatalF(format string, v ...any) {
	_logger.Fatalf(formatLine(logLevelFatal, format), v...)
}

func FatalW(message string, kv ...any) {
	_logger.SetPrefix(logLevelFatal)
	_logger.Fatalln(message, kv)
}

// Panic color: red
func Panic(v ...any) {
	_logger.SetPrefix(logLevelPanic)
	_logger.Panicln(v...)
}

func PanicF(format string, v ...any) {
	_logger.Panicf(formatLine(logLevelPanic, format), v...)
}

func PanicW(message string, kv ...any) {
	_logger.SetPrefix(logLevelPanic)
	_logger.Panicln(message, kv)
}

func outputLine(level string, v ...any) {
	_logger.SetPrefix(level)
	_logger.Println(v...)
}

func outputFormatLine(level, format string, v ...any) {
	_logger.Printf(formatLine(level, format), v...)
}

func formatLine(level, format string) string {
	_logger.SetPrefix(level)
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	return format
}

func outputKeyValueLine(level, message string, kv ...any) {
	_logger.SetPrefix(level)
	_logger.Println(message, kv)
}
