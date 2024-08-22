package util

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInternalKeywordSizeOf(t *testing.T) {
	println(unsafe.Sizeof("viviviviviviviviviviviviviviviviviviviviviv"))
	println(unsafe.Sizeof("v"))                // 16B
	println(unsafe.Sizeof(map[int]struct{}{})) // 8B
	println(unsafe.Sizeof(func() {}))          // 8B
	println(unsafe.Sizeof([]int{}))            // 24B
	println(unsafe.Sizeof([5]int{}))           // len*type_size B
}

func TestBytes2StringNoCopy(t *testing.T) {
	bytesOf := []byte{104, 98, 111}
	fmt.Println(string(bytesOf))
	fmt.Println(Bytes2StringNoCopy(bytesOf))
}

func TestString2BytesNoCopy(t *testing.T) {
	strOf := "Alan"
	fmt.Println([]byte(strOf))
	fmt.Println(String2BytesNoCopy(strOf))
}
