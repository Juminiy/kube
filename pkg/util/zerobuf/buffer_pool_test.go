package zerobuf

import "testing"

func TestApString_String(t *testing.T) {
	buf := Get()
	defer buf.Free()
	for range 1 << 5 {
		_, _ = buf.WriteString("ssr")
		t.Log(buf.UnsafeString())
	}
}
