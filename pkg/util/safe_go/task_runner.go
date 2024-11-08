package safe_go

import (
	"github.com/Juminiy/kube/pkg/util"
	"sync"
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
	return ""
}

func (r *Runner) Go() *Runner {

	return r
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
