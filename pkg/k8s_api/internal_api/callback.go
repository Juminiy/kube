package internal_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
)

type CallBack struct {
	Latest    any
	LatestErr error
	Create    util.Func
	Update    util.Func
	Delete    util.Func
	List      util.Func
}

func LogLatestCallBack() *CallBack {
	callback := &CallBack{}
	logFunc := func() error {
		if callback.Latest != nil {
			stdlog.Info(callback.Latest)
		}
		callback.Latest = nil
		return nil
	}
	callback.Create = logFunc
	callback.Update = logFunc
	callback.Delete = logFunc
	callback.List = logFunc
	return callback
}

func SilentCallBack() *CallBack {
	return &CallBack{
		Create: util.NothingFunc(),
		Update: util.NothingFunc(),
		Delete: util.NothingFunc(),
		List:   util.NothingFunc(),
	}
}
