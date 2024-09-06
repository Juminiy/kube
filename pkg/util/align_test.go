package util

import (
	"testing"
	"unsafe"
)

func TestInternalKeywordSizeOf(t *testing.T) {
	println(unsafe.Sizeof("viviviviviviviviviviviviviviviviviviviviviv")) // 16B
	println(unsafe.Sizeof("v"))                                           // 16B
	println(unsafe.Sizeof(map[int]struct{}{}))                            // 8B
	println(unsafe.Sizeof(func() {}))                                     // 8B
	println(unsafe.Sizeof([]int{}))                                       // 24B
	println(unsafe.Sizeof([5]int{}))                                      // len*type_size B
}
