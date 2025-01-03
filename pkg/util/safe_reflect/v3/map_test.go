package safe_reflectv3

import "testing"

func TestV_MapSetField(t *testing.T) {
	setv := map[string]any{
		"k1": "v2",
		"k2": 666,
		"1":  "v2",
		"2":  666,
	}
	for _, mapv := range []any{
		map[string]any{
			"k1": "v1",
			"k2": 2,
		},
		map[string]int{
			"k1": 1,
			"k2": 2,
		},
		map[int]any{
			1: "v1",
			2: 2,
		},
	} {
		Indirect(mapv).MapSetField(setv)
		t.Log(mapv)
	}
}
