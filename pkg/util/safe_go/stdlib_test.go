package safe_go

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"strconv"
	"testing"
)

func TestRun(t *testing.T) {
	testRunT(t, Run)
}

func TestDryRun(t *testing.T) {
	testRunT(t, DryRun)
}

func testRunT(t *testing.T, runner func(...util.Func) error) {
	defer func() {
		stdlog.Info("TestRun end")
	}()
	stdlog.Info("TestRun start")

	err := runner(getTestTasks()...)
	if err != nil {
		stdlog.Error(err.Error())
		return
	}

	//time.Sleep(util.TimeSecond(10))
}

func getTestTasks() []util.Func {
	fns := make([]util.Func, 1024)
	//fakedErr := errors.New("faked error")
	for i := range fns {
		fns[i] = func() error {
			if i == 0 {
				return nil
			} else if i%7 == 0 {
				return errors.New("faked error index: " + strconv.Itoa(i) + "")
			} else if i%555 == 0 {
				panic("faked panic index: " + strconv.Itoa(i) + "")
			}
			return nil
		}
	}
	return fns
}
