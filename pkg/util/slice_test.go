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
