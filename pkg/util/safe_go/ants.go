package safe_go

import (
	"github.com/Juminiy/kube/pkg/log_api/ilog"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/panjf2000/ants/v2"
)

var (
	_antsPool *ants.Pool
)

const (
	_defaultPoolSize = 16
	_defaultBlocking = 1024
)

func AntsInit() {
	var initPoolErr error
	_antsPool, initPoolErr = ants.NewPool(_defaultPoolSize,
		ants.WithOptions(ants.Options{
			ExpiryDuration:   util.DurationDay,
			PreAlloc:         true,
			MaxBlockingTasks: _defaultBlocking,
			Nonblocking:      false,
			PanicHandler: func(v interface{}) {
				util.PanicHandler(v)
			},
			Logger:       ilog.AntsLogger{},
			DisablePurge: false,
		}))
	if initPoolErr != nil {
		stdlog.ErrorF("ants pool init error: %s", initPoolErr.Error())
		return
	}
}

func AntsRunner(fns ...util.Func) error {
	var err error
	for _, fn := range fns {
		if fn != nil {
			err = _antsPool.Submit(func() {
				if err != nil {
					return
				}
				err = fn()
				//util.SilentErrorf("submit to ants task, but catch error itself, error", fn())
			})
			//util.SilentErrorf("task submit to ants pool by Submit(task) return error", err)
		}
	}
	return err
}
