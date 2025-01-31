package util

import "unsafe"

func GetTypeByteSize(varOf any) uintptr {
	return unsafe.Sizeof(varOf)
}

func GetValueByteSize(varOf any) int {
	switch v := varOf.(type) {
	case string:
		return len(v)
	case []byte:
		return len(v)
	default:
		return -1
	}
}
