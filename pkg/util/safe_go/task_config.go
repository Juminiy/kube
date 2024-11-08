package safe_go

import (
	"context"
	"github.com/Juminiy/kube/pkg/internal"
	"github.com/Juminiy/kube/pkg/util"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"time"
)

func WithContext(context context.Context) Option {
	return func(r *Runner) {
		r.context = context
	}
}

func WithLimit(limit int) Option {
	return func(r *Runner) {
		r.limit = true
		r.goLimit = limit
		r.limiter = &limiter{tokens: make(chan struct{}, limit)}
	}
}

func WithErrCancel() Option {
	return func(r *Runner) {
		r.errCancel = true
		r.errDryRun = false
		r.panicRecover = false
	}
}

func WithErrDryRun() Option {
	return func(r *Runner) {
		r.errCancel = false
		r.errDryRun = true
		r.panicRecover = false
		r.errs = make([]error, r.tasksz)
	}
}

func WithPanicRecover() Option {
	return func(r *Runner) {
		r.errCancel = false
		r.errDryRun = false
		r.panicRecover = true
		r.panicstack = make([][]byte, r.tasksz)
	}
}

func WithTimeoutCancel(timeout time.Duration) Option {
	return func(r *Runner) {
		r.timeoutCancel = true
		r.timeout = timeout
	}
}

func WithDeadlineCancel(deadline time.Time) Option {
	return func(r *Runner) {
		r.deadlineCancel = true
		r.deadline = deadline
	}
}

func WithProgressReport(bar chan<- int) Option {
	return func(r *Runner) {
		r.progressReport = true
		r.progress = &progress{
			bar:     bar,
			tot:     atomic.Int64{},
			success: atomic.Int64{},
			failure: atomic.Int64{},
		}
	}
}

// 8B
type config struct {
	limit          bool
	errCancel      bool // mutual-exclusion with errDryRun, panicRecover
	errDryRun      bool // mutual-exclusion with errCancel, panicRecover
	panicRecover   bool // mutual-exclusion with errCancel, errDryRun
	timeoutCancel  bool
	progressReport bool
	deadlineCancel bool
	_              bool // bool_align
}

// 64B
type option struct {
	context  context.Context
	goLimit  int
	timeout  time.Duration
	deadline time.Time
	tasks    []util.Func
	tasksz   int
}

const noGoLimit = -1

// 64B
type runner struct {
	wgroup *sync.WaitGroup
	egroup *errgroup.Group // for align, always nil

	golaunchtime time.Time
	gofinishtime time.Time
}

// 8B
type limiter struct {
	tokens chan struct{}
}

// 32B
type errhandler struct {
	// err cancel: when config.errCancel = true
	err     error
	errOnce *sync.Once

	// errs collect: when config.errDryRun = true
	errs []error

	// panicstack collect: when config.panicRecover = true
	panicstack [][]byte
}

// 32B
type progress struct {
	bar     chan<- int
	tot     atomic.Int64
	success atomic.Int64
	failure atomic.Int64
}

type volatile struct {
	_      internal.NoCmp
	noCopy internal.NoCopy
}
