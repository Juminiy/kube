package util

import (
	"maps"
	"slices"
	"testing"
)

func TestToElemPtrSlice(t *testing.T) {
	s := []string{"sq", "rw", "cs"}
	slices.All(ToElemPtrSlice(s))(func(i int, s *string) bool {
		t.Log(*s)
		return true
	})
}

func TestToElemPtrMap(t *testing.T) {
	m := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	maps.All(ToElemPtrMap(m))(func(k string, v *string) bool {
		t.Log(k, *v)
		return true
	})
}
