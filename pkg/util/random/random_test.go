package random

import "testing"

func TestInteger(t *testing.T) {
	for i := range 100 {
		t.Log(Integer(i))
	}
}
