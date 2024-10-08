package reflect

import (
	"reflect"
	"testing"
)

func TestMapAssign(t *testing.T) {
	var m1 map[string]string
	// nil map
	mapKeyExistAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Logf("nil map key exist assign %v", m1)

	// nil map
	mapDryAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Logf("nil map dry assign %v", m1)

	// len0 map assign key exist
	m1 = make(map[string]string)
	mapKeyExistAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Logf("len=0 map key exist assign %v", m1)

	// len0 map dry assign
	//m1 = make(map[string]string)
	mapDryAssign(reflect.ValueOf(m1), "v1", "v2")
	t.Logf("len=0 map dry assign %v", m1)

	mapDryDelete(reflect.ValueOf(m1), "v2")
	t.Logf("map dry delete key not exist %v", m1)

	// map dry delete
	mapDryDelete(reflect.ValueOf(m1), "v1")
	t.Logf("map dry delete key exist %v", m1)

	// map key_type mismatch
	mapDryAssign(reflect.ValueOf(m1), 1, "v2")
	t.Logf("map assign key_type mismatch %v", m1)

	// map value_type mismatch
	mapDryAssign(reflect.ValueOf(m1), "v2", 1)
	t.Logf("map assign value_type mismatch %v", m1)

	// map key_type mismatch and value_type mismatch
	mapDryAssign(reflect.ValueOf(m1), 1, 1)
	t.Logf("map assign key_type and value_type mismatch %v", m1)
}
