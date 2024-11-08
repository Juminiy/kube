package safe_go

import (
	"github.com/Juminiy/kube/pkg/util"
	gostruntime "github.com/dubbogo/gost/runtime"
	"golang.org/x/sync/errgroup"
	"sync"
)

func Run(fns ...util.Func) error {
	grp, ctx := errgroup.WithContext(util.TODOContext())
	defer ctx.Done()
	for _, fn := range fns {
		if fn != nil {
			grp.Go(fn)
		}
	}
	return grp.Wait()
}

func DryRun(fns ...util.Func) error {
	wg := &sync.WaitGroup{}
	errs := make([]error, len(fns))

	for i, fn := range fns {
		safeGoWithWaitGroup(wg, fn, &errs[i])
	}
	wg.Wait()
	return util.MergeError(errs...)
}

func safeGoWithWaitGroup(wg *sync.WaitGroup, fn util.Func, err *error) {
	if fn != nil {
		gostruntime.GoSafely(
			wg,
			false,
			func() {
				fnErr := fn()
				if fnErr != nil {
					*err = fnErr
				}
			}, func(r interface{}) {
				util.PanicHandler(r)
			},
		)
	}
}
