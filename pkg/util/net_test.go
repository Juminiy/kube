package util

import "testing"

// +passed
func TestLookupIP(t *testing.T) {
	/*t.Log(LookupIP("harbor.local"))
	t.Log(LookupIP("harbor.local:18111"))*/
	t.Log(LookupIP("host.local"))
	t.Log(LookupIP("host.local:18111"))
}

func TestTrimProto(t *testing.T) {
	t.Log(TrimProto("host.local"))
	t.Log(TrimProto("host.local"))
	t.Log(TrimProto("host.local"))
	t.Log(TrimProto("http://host.local"))
	t.Log(TrimProto("tcp://host.local"))
	t.Log(TrimProto("udp://host.local"))
}
