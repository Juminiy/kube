package util

import "unsafe"

func GetTypeByteSize(varOf any) uintptr {
	return unsafe.Sizeof(varOf)
}

func GetValueByteSize(varOf any) int {
	switch varOf.(type) {
	case string:
		return len(varOf.(string))
	case []byte:
		return len(varOf.([]byte))
	default:
		return -1
	}
}
