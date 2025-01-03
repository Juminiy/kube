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
			"k1": "v1", // k1: v2
			"k2": 2,    // k2: 666
			// 1: v2
			// 2: 666
		},
		map[string]int{
			"k1": 1, // k1: 1
			"k2": 2, // k2: 666
			// 2: 666
		},
		map[int]any{
			1: "v1",
			2: 2,
		},
	} {
		Direct(mapv).MapSetField(setv)
		t.Log(mapv)
	}
}

func TestV_MapSetField2(t *testing.T) {
	setv := map[string]any{
		"k1": nil,
		"k2": nil,
		"1":  nil,
		"2":  nil,
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
		Direct(mapv).MapSetField(setv)
		t.Log(mapv)
	}
}
