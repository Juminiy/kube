package safe_go

import (
	"errors"
	"fmt"
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
		t.Log("TestRun end")
	}()
	t.Log("TestRun start")

	err := runner(getTestTasks()...)
	if err != nil {
		t.Logf("error: %s", err.Error())
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
				util.Must(fmt.Errorf("faked panic index: %d", i))
			}
			return nil
		}
	}
	return fns
}
