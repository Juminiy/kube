package badger_api

import (
	"github.com/Juminiy/kube/pkg/log_api/zaplog"
	"github.com/dgraph-io/badger/v4"
)

func New(fd string) (*badger.DB, error) {
	bOpts := badger.DefaultOptions(fd)
	bOpts.Logger = _ZapLogger{}
	return badger.Open(bOpts)
}

type _ZapLogger struct{}

func (_ZapLogger) Errorf(f string, v ...interface{}) {
	zaplog.ErrorF(f, v...)
}
func (_ZapLogger) Warningf(f string, v ...interface{}) {
	zaplog.WarnF(f, v...)
}
func (_ZapLogger) Infof(f string, v ...interface{}) {
	zaplog.InfoF(f, v...)
}
func (_ZapLogger) Debugf(f string, v ...interface{}) {
	zaplog.DebugF(f, v...)
}
