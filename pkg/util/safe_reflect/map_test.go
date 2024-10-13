package safe_reflect

import (
	"strconv"
	"testing"
)

// +passed
func TestMapAssign(t *testing.T) {
	var m1 map[string]string
	// nil map
	Of(m1).mapKeyExistAssign("k1", "v1")
	t.Logf("nil map key exist assign %v", m1)

	// nil map
	Of(m1).mapDryAssign("k1", "v1")
	t.Logf("nil map dry assign %v", m1)

	// len0 map assign key exist
	m1 = make(map[string]string)
	Of(m1).mapKeyExistAssign("k1", "v1")
	t.Logf("len=0 map key exist assign %v", m1)

	// len0 map dry assign
	Of(m1).mapDryAssign("k1", "v1")
	t.Logf("len=0 map dry assign %v", m1)

	// map dry delete
	Of(m1).mapDryDelete("k2")
	t.Logf("map dry delete key not exist %v", m1)

	// map dry delete
	Of(m1).mapDryDelete("k1")
	t.Logf("map dry delete key exist %v", m1)

	// map key_type mismatch
	Of(m1).mapDryAssign(1, "v2")
	t.Logf("map assign key_type mismatch %v", m1)

	// map value_type mismatch
	Of(m1).mapDryAssign("k2", 1)
	t.Logf("map assign value_type mismatch %v", m1)

	// map key_type mismatch and value_type mismatch
	Of(m1).mapDryAssign(1, 1)
	t.Logf("map assign key_type and value_type mismatch %v", m1)
}

func TestTypVal_MapAssign2(t *testing.T) {
	var m map[string]string
	Of(m).mapNilDryAssign("k1", "v1")
	t.Log(m)

	for idx := range 16 {
		Of(&m).mapNilDryAssign("k"+strconv.Itoa(idx), "v"+strconv.Itoa(idx))
		t.Log(m)
	}

	var sl []int
	Of(sl).mapNilDryAssign("1", 2)
	Of(&sl).mapDryAssign(2, "1")
}
