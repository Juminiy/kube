package util

import "unsafe"

type Func func() error

func String2BytesNoCopy(s string) []byte {
	sh := (*[2]uintptr)(unsafe.Pointer(&s))
	bh := [3]uintptr{sh[0], sh[1], sh[1]}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2StringNoCopy(bs []byte) string {
	bh := (*[3]uintptr)(unsafe.Pointer(&bs))
	sh := [2]uintptr{bh[0], bh[1]}
	return *(*string)(unsafe.Pointer(&sh))
}
