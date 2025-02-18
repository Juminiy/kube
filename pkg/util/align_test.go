package util

import (
	"bytes"
	"io"
	"os"
	"testing"
	"unsafe"
)

func TestInternalKeywordSizeOf(t *testing.T) {
	t.Log(unsafe.Sizeof("viviviviviviviviviviviviviviviviviviviviviv")) // 16B
	t.Log(unsafe.Sizeof("v"))                                           // 16B
	t.Log(unsafe.Sizeof(map[int]struct{}{}))                            // 8B
	t.Log(unsafe.Sizeof(func() {}))                                     // 8B
	t.Log(unsafe.Sizeof([]int{}))                                       // 24B
	t.Log(unsafe.Sizeof([5]int{}))                                      // len*type_size B
	t.Log(unsafe.Sizeof(bool(true)))                                    // 1B
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

func TestG(t *testing.T) {
	var logFn = func(lv, rv string, eq bool) {
		eqStr := "=="
		if !eq {
			eqStr = "!="
		}
		t.Logf("key[T] value: %8v %s %8v", lv, eqStr, rv)
	}
	type key[T any] struct{}
	var kInt any = key[int]{}
	var kStr any = key[string]{}
	logFn("key[int]", "key[int]", kInt == kInt)
	logFn("key[str]", "key[str]", kStr == kStr)
	logFn("key[int]", "key[str]", kInt == kStr)
}

func TestSize(t *testing.T) {
	t.Log(unsafe.Sizeof(struct {
		B1 bool
		B2 bool
		I  int
	}{}))
}
