package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
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

func TestFuncMake(t *testing.T) {
	t.Log(
		FuncMake(
			[]any{int32(1), uint(1)}, []any{uint(0), int32(0)}, false,
			MetaFunc(func(values []reflect.Value) []reflect.Value {
				return []reflect.Value{values[1], values[0]}
			}),
		).(func(int32, uint) (uint, int32))(10, 11),
	)

	// let it panic is ok
	//util.Recover(func() {
	//	t.Log(
	//		FuncMake(
	//			[]any{int32(1), uint(1)}, []any{uint(0), int32(0)}, false,
	//			MetaFunc(func(values []reflect.Value) []reflect.Value {
	//				return []reflect.Value{values[1], values[0]}
	//			}),
	//		).(func(int32, uint) (int32, int32))(10, 11),
	//	)
	//})

}

func TestTypVal_FuncCall(t *testing.T) {
	tfn := func() { t.Log("test func") }
	t.Log(Of(&tfn).FuncCall(nil))

	t.Log(Of(&tfn).FuncCall([]any{[]int{10}}))

	t2fn := func(a int, b int) int { return a + b }
	t.Log(Of(&t2fn).FuncCall([]any{1, 2}))

}

func TestTypVal_HasMethod(t *testing.T) {
	t.Log(Of(&t0{}).HasMethod("v", []any{}, []any{})) // unexported value-receiver
	t.Log(Of(&t0{}).HasMethod("k", []any{}, []any{})) // unexported pointer-receiver
	t.Log(Of(&t0{}).HasMethod("V", []any{}, []any{})) // exported value-receiver
	t.Log(Of(&t0{}).HasMethod("K", []any{}, []any{})) // exported pointer-receiver

	util.TestLongHorizontalLine(t)
	t.Log(Of(&t0{}).HasMethod("v", []any{t0{}}, []any{})) // unexported value-receiver
	t.Log(Of(&t0{}).HasMethod("k", []any{t0{}}, []any{})) // unexported pointer-receiver
	t.Log(Of(&t0{}).HasMethod("V", []any{t0{}}, []any{})) // exported value-receiver
	t.Log(Of(&t0{}).HasMethod("K", []any{t0{}}, []any{})) // exported pointer-receiver

	util.TestLongHorizontalLine(t)
	t.Log(Of(&t0{}).HasMethod("v", []any{&t0{}}, []any{})) // unexported value-receiver
	t.Log(Of(&t0{}).HasMethod("k", []any{&t0{}}, []any{})) // unexported pointer-receiver
	t.Log(Of(&t0{}).HasMethod("V", []any{&t0{}}, []any{})) // exported value-receiver
	t.Log(Of(&t0{}).HasMethod("K", []any{&t0{}}, []any{})) // exported pointer-receiver
}
