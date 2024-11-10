package safe_go

import (
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	gostruntime "github.com/dubbogo/gost/runtime"
	"golang.org/x/sync/errgroup"
	"runtime/debug"
	"time"
)

// errCancel
// noLimit
// noProgress
func (r *Runner) run2() {
	r.timeRecorder(
		func() { r.egroup, r.context = errgroup.WithContext(r.context) },
		func() {
			for _, task := range r.tasks {
				r.egroup.Go(task)
			}
		},
		func() { r.err = r.egroup.Wait() },
		func() {},
	)
}

// goLimit
// errGryRun, panicRecover
// progressBar
func (r *Runner) run73() {
	r.timeRecorder(
		func() {},
		func() {
			for i := range r.tasks {
				if task := r.tasks[i]; task != nil {
					r.tokens <- struct{}{}
					gostruntime.GoSafely(
						r.wgroup,
						false,
						func() {
							defer func() {
								<-r.tokens
								if r.errs[i] != nil {
									r.failure.Add(1)
								} else {
									r.success.Add(1)
								}
								r.tot.Add(1)
								r.bar <- safe_cast.I64toI(r.tot.Load()) * 100 / r.tasksz
							}()
							r.errs[i] = task()
						},
						func(v interface{}) {
							r.panicstack[i] = debug.Stack()
						},
					)
				}
			}
		},
		func() { r.wgroup.Wait() },
		func() {},
	)
}

func (r *Runner) timeRecorder(before, taskLaunch, groupWait, after func()) {
	before()

	r.golaunchtimestart = time.Now()
	taskLaunch()
	r.golaunchtimeend = time.Now()

	groupWait()
	r.gofinishtime = time.Now()

	after()
}
