//go:build go1.20

package util

import "unsafe"

func s2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func b2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
