package util

import "testing"

// +passed
func TestLookupIP(t *testing.T) {
	t.Log(LookupIP("harbor.local"))
	t.Log(LookupIP("harbor.local:18111"))
	t.Log(LookupIP("192.168.31.66"))
	t.Log(LookupIP("192.168.31.66:18111"))
}

func TestTrimProto(t *testing.T) {
	t.Log(TrimProto("192.168.31.66"))
	t.Log(TrimProto("192.168.31.66"))
	t.Log(TrimProto("192.168.31.66"))
	t.Log(TrimProto("http://192.168.31.66"))
	t.Log(TrimProto("tcp://192.168.31.66"))
	t.Log(TrimProto("udp://192.168.31.66"))
}
