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
		r.errOnce = &sync.Once{}
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
		r.errs = make([]error, r.tasksz)
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

// 16B
type config struct {
	limit          bool
	errCancel      bool // mutual-exclusion with errDryRun, panicRecover
	errDryRun      bool // mutual-exclusion with errCancel, panicRecover
	panicRecover   bool // mutual-exclusion with errCancel, errDryRun
	timeoutCancel  bool
	progressReport bool
	deadlineCancel bool
	_              bool // bool_align
	flag           int64
}

const (
	flagInvalid = 0
	flagLimit   = 1 << (iota - 1)
	flagErrCancel
	flagErrDryRun
	flagPanicRecover
	flagTimeoutCancel
	flagDeadlineCancel
	flagProgressBar
)

func (c *config) flagBitSet() {
	if c.limit {
		c.flag |= flagLimit
	}

	if c.errCancel {
		c.flag |= flagErrCancel
	} else if c.errDryRun {
		c.flag |= flagErrDryRun
	} else if c.panicRecover {
		c.flag |= flagPanicRecover
	}

	if c.timeoutCancel {
		c.flag |= flagTimeoutCancel
	}

	if c.deadlineCancel {
		c.flag |= flagDeadlineCancel
	}

	if c.progressReport {
		c.flag |= flagProgressBar
	}
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

// 96B
type runner struct {
	wgroup *sync.WaitGroup // for errorDryRun, panicRecover
	egroup *errgroup.Group // for errorCancel

	golaunchtimestart time.Time // the first task start launch
	golaunchtimeend   time.Time // the last task finish launch
	gofinishtime      time.Time // the last task finish
	_                 time.Time // time align
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

type report struct {
	TotalTask       int64     `json:"total_task"`
	SuccessTask     int64     `json:"success_task"`
	FailureTask     int64     `json:"failure_task"`
	ErrorList       []string  `json:"error_list"`
	PanicStacktrace []string  `json:"panic_stacktrace"`
	FirstTaskLaunch time.Time `json:"first_task_launch"`
	LastTaskLaunch  time.Time `json:"last_task_launch"`
	LastTaskFinish  time.Time `json:"last_task_finish"`
}

type volatile struct {
	_      internal.NoCmp
	noCopy internal.NoCopy
}
