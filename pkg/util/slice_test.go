package util

import (
	"slices"
	"testing"
)

func TestSliceSum(t *testing.T) {
	sl := []int{1, 2, 3, -3, 5, 6}
	sum := 0
	sumFn := func(index int, elem int) bool {
		if elem < 0 || index == len(sl)-1 {
			return false
		}
		sum += elem
		return true
	}

	slices.All(sl)(sumFn)
	t.Log(sum)

	sum2 := 0
	slices.Backward(sl)(sumFn)
	t.Log(sum2)
}

func TestSliceCopyByFunc(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Log(arr) // 1,2,3
	callMain(t, arr)
	t.Log(arr) // 1,2,3
}

func callMain(t *testing.T, arr []int) {
	t.Log(arr) // 1,2,3
	callOther(t, arr)
	t.Log(arr) // 1,2,3
}

func callOther(t *testing.T, arr []int) {
	t.Log(arr) // 1,2,3
	arr = append(arr, 4, 5, 6)
	arr[0] = 666
	t.Log(arr) // 666,2,3,4,5,6
}
