package safe_go

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"strconv"
	"testing"
	"time"
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

	fns := make([]util.Func, 1024)
	//fakedErr := errors.New("faked error")
	for i := range fns {
		fns[i] = func() error {
			if i > 0 && i%7 == 0 {
				return errors.New("faked error index: " + strconv.Itoa(i) + "")
			}
			return nil
		}
	}

	err := runner(fns...)
	if err != nil {
		t.Log(err.Error())
		return
	}

	time.Sleep(util.TimeSecond(10))
}
