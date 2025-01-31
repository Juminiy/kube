package safe_validator

import (
	"testing"
)

func TestStrict_Map(t *testing.T) {
	t.Log(Strict().MapE(map[string]t0{
		"kv1": correctT0Elem[0],
		"kv2": correctT0Elem[1],
		"kv3": correctT0Elem[2],
	}))
	t.Log(Strict().MapE(map[string]t0{
		"kv1": errT0Elem[0],
		"kv2": errT0Elem[1],
		"kv3": errT0Elem[2],
	}))
}
