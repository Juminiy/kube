package safe_go

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"math/rand/v2"
	"testing"
	"time"
)

func TestNewRunner(t *testing.T) {
	t.Log(
		NewRunner(
			getTestTasks(),
			WithErrCancel(),
		).Go().Error(),
	)
}

func TestNewRunner2(t *testing.T) {
	taskBar := make(chan int)
	defer func() { close(taskBar) }()
	go func() {
		for prog := range taskBar {
			stdlog.InfoF("progress: %d%%", prog)
		}
	}()

	taskRunner := NewRunner(getTestTasks2(),
		WithPanicRecover(),
		WithLimit(16),
		WithProgressReport(taskBar),
	)
	t.Log(taskRunner.Go().Report())
}

func getTestTasks2() []util.Func {
	tasks := make([]util.Func, 32)

	for i := range 32 {
		tasks[i] = func() error {
			sleepSec := rand.IntN(10)
			stdlog.InfoW("sleep for", "time", sleepSec)
			time.Sleep(util.TimeSecond(sleepSec))
			if sleepSec == 7 {
				panic("faked panic")
			} else if sleepSec == 5 {
				return errors.New("faked error")
			}
			return nil
		}
	}

	return tasks
}
