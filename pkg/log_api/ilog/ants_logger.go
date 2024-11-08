package ilog

import "github.com/Juminiy/kube/pkg/log_api/stdlog"

type AntsLogger struct{}

func (AntsLogger) Printf(format string, args ...interface{}) {
	stdlog.InfoF(format, args...)
}

var _antsLogger AntsLogger

func Infof(format string, v ...any) {
	_antsLogger.Printf(format, v...)
}

func InfoF(format string, v ...any) {
	_antsLogger.Printf(format, v...)
}
