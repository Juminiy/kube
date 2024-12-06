package safe_validator

import "testing"

func TestStrict_Array(t *testing.T) {
	t.Log(Strict().ArrayE(correctT0Elem))
	t.Log(Strict().ArrayE(errT0Elem))
}
