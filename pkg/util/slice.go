package util

const (
	MagicSliceCap = (iota + 1) << 3
)

func ElemIn[E comparable](e E, s ...E) bool {
	for i := range s {
		if s[i] == e {
			return true
		}
	}
	return false
}

// ElemsIn
// src all in dst return true
// ,or not return false
func ElemsIn[E comparable](src, dst []E) bool {
	eMap := make(map[E]struct{}, len(dst))
	for dstI := range dst {
		eMap[dst[dstI]] = struct{}{}
	}
	for srcI := range src {
		if _, ok := eMap[src[srcI]]; !ok {
			return false
		}
	}
	return true
}

const (
	LoNotFound = -1
)
