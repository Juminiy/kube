package util

import (
	"maps"
	"testing"
)

func TestMapOk(t *testing.T) {
	m1 := map[string]map[string]struct{}{
		"Field1": {"K1": {}, "K2": {}},
	}
	m2 := map[string]map[string]struct{}{
		"Field1": {"K3": {}},
	}
	t.Log(m1)
	t.Log(m2)

	// wanna result -> Field1: K1, K2, K3

	for field := range m1 {
		maps.Copy(m1[field], m2[field])
	}

	t.Log(m1)
	t.Log(m2)
}
func TestMapMerge(t *testing.T) {
	m1 := map[string]map[string]struct{}{
		"Field1": {"K1": {}, "K2": {}},
	}
	m2 := map[string]map[string]struct{}{
		"Field1": {"K3": {}},
	}
	t.Log(m1)
	t.Log(m2)

	MapMerge(m1, m2)

	t.Log(m1)
	t.Log(m2)

}

func TestMapEvict(t *testing.T) {
	m1 := map[string]int{
		"K1": 1,
		"K2": 2,
		"K3": 3,
	}

	m2 := map[string]uint{
		"K1": 99,
		"K2": 88,
		"K4": 77,
	}

	t.Log(m1)
	t.Log(m2)

	MapEvict(m1, m2)
	t.Log(m1)
	t.Log(m2)
}
