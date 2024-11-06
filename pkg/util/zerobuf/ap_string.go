package zerobuf

import "unsafe"

type String interface {
	Stringer
	Allocator
	Pooler
}

type apString struct {
	b []byte
	apBufferPool
}

func (ap *apString) WriteString(s string) (int, error) {
	ap.b = append(ap.b, s...)
	return len(ap.b), nil
}

func (ap *apString) WriteByte(b byte) error {
	ap.b = append(ap.b, b)
	return nil
}

func (ap *apString) Write(bs []byte) (int, error) {
	ap.b = append(ap.b, bs...)
	return len(ap.b), nil
}

func (ap *apString) Len() int {
	return len(ap.b)
}

func (ap *apString) Cap() int {
	return cap(ap.b)
}

// String safe
func (ap *apString) String() string {
	return string(ap.b)
}

// UnsafeString unsafe
// be cautious when use it, always BUGs
// Deprecated
func (ap *apString) UnsafeString() string { return unsafe.String(unsafe.SliceData(ap.b), len(ap.b)) }

func (ap *apString) Bytes() []byte {
	return ap.b
}

func (ap *apString) Reset() {
	ap.b = ap.b[:0]
}

func (ap *apString) Grow(n int) {
	if n < 0 {
		return
	} else if cap(ap.b)-len(ap.b) < n {
		ap.grow(n)
	}
}

// inlined
// copied from: strings.Builder.grow(n)
func (ap *apString) grow(n int) {
	buf := makeNoZero(2*cap(ap.b) + n)[:len(ap.b)]
	copy(buf, ap.b)
	ap.b = buf
}

func (ap *apString) Free() {
	ap.apBufferPool.put(ap)
}

//go:linkname makeNoZero internal/bytealg.MakeNoZero
func makeNoZero(int) []byte
