package zero_reflect

import "testing"

type t0 struct {
	I32 int32
	I64 int64
	Str string
	Byt byte
}

func TestTypeOf(t *testing.T) {
	for i := range 16 {
		t0Typ := TypeOf(t0{I32: int32(i)})
		t.Log(t0Typ)
	}

	for i := range 16 {
		t0slTyp := TypeOf([]t0{{I64: int64(i)}})
		t.Log(t0slTyp)
	}
}
