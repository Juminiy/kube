package safe_cast

import (
	"math"
	"testing"
)

func TestSafeCast(t *testing.T) {
	castLog := func(fromTyp, toTyp string, fromVal, goVal any, myVal any) {
		t.Logf("go   cast %s(%v) -> %s(%v)", fromTyp, fromVal, toTyp, goVal)
		t.Logf("safe cast %s(%v) -> %s(%v)", fromTyp, fromVal, toTyp, myVal)
	}

	var i int = -10
	// negative to unsigned
	castLog("int", "uint", i, uint(i), ItoU(i))

	var i0 int = math.MaxInt
	// size reach overflow: positive overflow
	castLog("int", "int32", i0, int32(i0), ItoI32(i0))

	var i1 int = math.MinInt
	// size reach overflow: negative overflow
	castLog("int", "int32", i1, int32(i1), ItoI32(i1))

	var u uint = math.MaxUint
	// bound reach overflow: wide-bound to thin-bound abi.Type.Size_
	castLog("uint", "int32", u, int32(u), UtoI32(u))

	// bound reach overflow: equal abi.Type.Size_ but unsigned to signed bound lost
	castLog("uint", "int", u, int(u), UtoI(u))
}
