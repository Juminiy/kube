package safe_reflect

import "testing"

func TestTypVal_unpack(t *testing.T) {
	var i32 int32 = 10
	unpackEqual(i32, pointerCast(i32, 2)) // dead-loop found
}
