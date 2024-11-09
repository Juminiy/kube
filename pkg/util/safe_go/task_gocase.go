package safe_go

import (
	"golang.org/x/sync/errgroup"
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

// errGryRun
// panicRecover
// progressBar
func (r *Runner) run73() {

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
