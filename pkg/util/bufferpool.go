package util

import (
	uberbuffer "go.uber.org/zap/buffer"
)

var (
	_uberPool = uberbuffer.NewPool()
)

func GetUberBuffer() *uberbuffer.Buffer {
	return _uberPool.Get()
}

const (
	MagicBufferCap = 16 * Ki
)

func GetBuffer() []byte {
	return make([]byte, 0, MagicBufferCap)
}

func PutBuffer(buf []byte) {

}
