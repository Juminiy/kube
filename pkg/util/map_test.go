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

func TestMapVal(t *testing.T) {
	m := map[string]string{"k1": "v1"}
	t.Log(MapElem(m, "k2"))
	t.Log(MapElem(m, "k1"))

	m = nil
	t.Log(MapElem(m, "k2"))
	t.Log(MapElem(m, "k1"))
}

func TestMapClone(t *testing.T) {
	m := map[string]string{"k1": "v1"}

	mshallowCopy := m
	var mdeepCopy map[string]string
	mdeepCopy = MapCopy(mdeepCopy, m)
	t.Log(mshallowCopy)
	t.Log(mdeepCopy)
}

func TestMapCopy(t *testing.T) {
	m := map[string]string{"k1": "v1"}
	var newM map[string]string
	newM = MapCopy(newM, m)
	t.Log(newM)

	var newM2 = map[string]string{"k2": "v2"}
	MapCopy(newM2, m)
	t.Log(newM2)

	MapDelete(m, "k1", "k2")
	t.Log(m)

	MapClear(m, newM, newM2)
	t.Log(m, newM, newM2)
}

func TestMapKeys2(t *testing.T) {
	m := map[string]string{"k1": "v1", "k2": "v2"}
	t.Log(MapKeys(m))
	t.Log(MapValues(m))
}

func TestMapInsert(t *testing.T) {
	var m map[string]struct{}
	m = MapInsert(m, "xxx", "vvv", "qqq")
	t.Log(m)

	isl := []int{10, 10, 10}
	t.Log(Slice2Map[[]int, map[int]struct{}, int](isl))
}
