package zerobuf

import "testing"

func TestApString_String(t *testing.T) {
	buf := Get()
	for range 1 << 20 {
		_, _ = buf.WriteString("ssr")
		t.Log(buf.UnsafeString())
	}
	buf.Free()
}
