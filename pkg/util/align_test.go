package util

import (
	"bytes"
	"io"
	"os"
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
	println(unsafe.Sizeof(bool(true)))                                    // 1B
}

func fa(*bytes.Buffer) {}

func fb(io.Writer) {}

func dofa(func(*bytes.Buffer)) {}

func dofb(func(io.Writer)) {}

func fe(io.Writer) error {
	return nil
}

func ff(buffer *bytes.Buffer) *os.PathError {
	return nil
}

func dofe(func(io.Writer) error) {}

func TestF(t *testing.T) {
	dofa(fa)
	//dofa(fb) //error

	//dofb(fa) //error
	dofb(fb)

	dofe(fe)
	//dofe(ff) //error
}
