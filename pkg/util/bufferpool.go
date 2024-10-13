package util

import (
	valyalabuffer "github.com/valyala/bytebufferpool"
	uberbuffer "go.uber.org/zap/buffer"
)

var (
	_uberPool = uberbuffer.NewPool()
)

//func GetUberBuffer() *uberbuffer.Buffer {
//	return _uberPool.Get()
//}

const (
	MagicBufferCap = 16 * Ki
)

// GetBuffer
// when GetBuffer from ByteBuffer, must call PutBuffer after GetBuffer
func GetBuffer() *valyalabuffer.ByteBuffer {
	return valyalabuffer.Get()
}

// PutBuffer
// return memory to BufferPool
func PutBuffer(bb *valyalabuffer.ByteBuffer) {
	valyalabuffer.Put(bb)
}
