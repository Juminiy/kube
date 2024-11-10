//go:build !go1.20

package util

import (
	"reflect"
	"unsafe"
)

func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
