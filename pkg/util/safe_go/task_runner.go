package safe_go

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"sync"
	"time"
)

// Runner
/*
 * Runner
 * 1. context injection
 * 2. goroutine limit control
 * 3. error catch mode:
 * 	(1). one error cancel all: return error
 * 	(2). any error dry run: collect errors, return mergedError
 *  (3). any panic recover: collect errors, return mergedError, collect stacktrace,
 * 4. timeout control
 * 5. progress report
 */
type Runner struct {
	*config
	*option
	*runner
	*limiter
	*errhandler
	*progress

	volatile
}

type Option func(r *Runner)

func NewRunner(tasks []util.Func, options ...Option) *Runner {
	r := newDefaultRunner(tasks)

	for i := range options {
		options[i](r)
	}

	r.flagBitSet()
	return r
}

func newDefaultRunner(tasks []util.Func) *Runner {
	return &Runner{
		config: &config{
			errCancel: true,
		},
		option: &option{
			context: util.TODOContext(),
			goLimit: noGoLimit,
			tasks:   tasks,
			tasksz:  len(tasks),
		},
		runner: &runner{
			wgroup: &sync.WaitGroup{},
			egroup: nil,
		},
		limiter: nil,
		errhandler: &errhandler{
			err:        nil,
			errOnce:    &sync.Once{},
			errs:       nil,
			panicstack: nil,
		},
		progress: nil,
		volatile: volatile{},
	}
}

func (r *Runner) Error() string {
	switch {
	case r.errCancel:
		return errString(r.err)

	case r.errDryRun, r.panicRecover:
		return errString(util.MergeError(r.errs...))

	default:
		return ""
	}
}

func errString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

var _configDoNotSupport = errors.New("config do not support")

func (r *Runner) Go() *Runner {
	switch r.flag {
	case flagErrCancel:
		r.run2()

	case flagLimit | flagPanicRecover | flagProgressBar:
		r.run73()

	default:
		r.err = _configDoNotSupport
	}

	return r
}

func (r *Runner) Report() string {
	taskReport := &report{
		TotalTask:       0,
		SuccessTask:     0,
		FailureTask:     0,
		ErrorList:       nil,
		PanicStacktrace: nil,
		FirstTaskLaunch: time.Time{},
		LastTaskLaunch:  time.Time{},
		LastTaskFinish:  time.Time{},
	}

	if r.progressReport {
		taskReport.TotalTask = r.tot.Load()
		taskReport.SuccessTask = r.success.Load()
		taskReport.FailureTask = r.failure.Load()
	} else {
		taskReport.TotalTask = safe_cast.ItoI64(r.tasksz)
		if len(r.Error()) == 0 {
			taskReport.SuccessTask = taskReport.TotalTask
		} else {
			taskReport.FailureTask = taskReport.TotalTask
		}
	}

	if r.errCancel && r.err != nil {
		taskReport.ErrorList = append(taskReport.ErrorList, r.err.Error())
	} else if r.errDryRun || r.panicRecover {
		for i := range r.errs {
			if r.errs[i] != nil {
				taskReport.ErrorList = append(taskReport.ErrorList, r.errs[i].Error())
			}
		}
	}

	if r.panicRecover {
		var panicsz int64
		for i := range r.panicstack {
			if len(r.panicstack[i]) != 0 {
				panicsz++
				taskReport.PanicStacktrace = append(taskReport.PanicStacktrace, string(r.panicstack[i]))
			}
		}
		taskReport.SuccessTask -= panicsz
		taskReport.FailureTask += panicsz
	}

	taskReport.FirstTaskLaunch = r.golaunchtimestart
	taskReport.LastTaskLaunch = r.golaunchtimeend
	taskReport.LastTaskFinish = r.gofinishtime

	return safe_json.Pretty(taskReport)
}

func TaskConverter(fns ...util.Fn) []util.Func {
	tasks := make([]util.Func, 0, len(fns))
	for i := range fns {
		if fns[i] != nil {
			tasks[i] = func() error {
				fns[i]()
				return nil
			}
		}
	}
	return tasks
}
