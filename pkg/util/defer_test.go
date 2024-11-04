package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"testing"
)

func TestDefer(t *testing.T) {
	t.Log(testDefer(1)) // 111

	t.Log(testDefer2(2)) // 2

	t.Log(testDefer3(3)) // 3, 444

	testDefer4()

	testDefer5()
}

func testDefer(v int) int {
	v = 111
	defer func() {
		v = 222
	}()
	return v // return value is 111, v is no-escape, called: v is assigned to return-value
}

func testDefer2(v int) int {
	defer func() {
		v = 333
	}()
	return v // return value is v, v is no-escape, called: v is assigned to return-value
}

func testDefer3(v0 int) (v2 int, v int) {
	defer func() {
		v = 444
		v0 = 555
	}()
	return v0, v // return value is v2, v, called: v2 and v is escaped
}

func testDefer4() {
	stdlog.Info("func start") // 0.
	defer func() {
		stdlog.Info("defer func in stack0") // 3.
	}()
	defer func() {
		stdlog.Info("defer func in stack1") // 2.
	}()
	stdlog.Info("return func") // 1.
}

func testDefer5() int {
	defer func() {
		stdlog.Info("value destroy")
	}()
	return func() int {
		stdlog.Info("value return")
		return 1
	}()
}

func TestDefer6(t *testing.T) {
	for i := range 16 {
		defer deferfn0(i)
	}
}

func deferfn0(i int) {
	stdlog.InfoF("defer fn %d", i)
}
