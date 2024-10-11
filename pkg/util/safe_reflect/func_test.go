package safe_reflect

import (
	"strconv"
	"testing"
)

// code section static address can not change and address any case
func testFn(a int, b int) string {
	return strconv.Itoa(a) + " [fn] " + strconv.Itoa(b)
}
func testFn1(a int, b int) string {
	return strconv.Itoa(a) + " [fn1] " + strconv.Itoa(b)
}
func testFn2(a int, b string) string {
	return strconv.Itoa(a) + " [fn2] " + b
}

// +passed
func TestTypVal_FuncSet(t *testing.T) {
	fn := func() { t.Logf("1: %s", "from fn") }
	fn1 := func(int) { t.Logf("2: %s", "to fn1") }
	fn2 := func() { t.Logf("2: %s", "to fn2") }

	Of(fn).FuncSet(fn1) // not set
	fn()

	Of(fn).FuncSet(fn2) // not set
	fn()
}

// +passed
func TestTypVal_FuncSet2(t *testing.T) {
	fn := func() { t.Logf("1: %s", "from fn") }
	fn1 := func(int) { t.Logf("2: %s", "to fn1") }
	fn2 := func() { t.Logf("2: %s", "to fn2") }

	Of(&fn).FuncSet(fn1) // not set
	fn()

	Of(&fn).FuncSet(fn2) // set
	fn()
}

// +passed
func TestTypVal_FuncSet3(t *testing.T) {
	testFnV := testFn

	Of(testFn).FuncSet(testFn1) // not set
	t.Log(testFnV(66, 99))
	t.Log(testFn(666, 999))

	Of(testFn).FuncSet(testFn2) // not set
	t.Log(testFnV(22, 33))
	t.Log(testFn(222, 333))

	Of(&testFnV).FuncSet(testFn1) // not set
	t.Log(testFnV(666, 999))
	t.Log(testFn(666, 999))

	Of(&testFnV).FuncSet(testFn2) // not set
	t.Log(testFnV(666, 999))
	t.Log(testFn(222, 333))
}
