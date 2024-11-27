package safe_parse

import "testing"

func TestParse(t *testing.T) {
	t.Logf("%+v", Parse(""))
	t.Logf("%+v", Parse("v"))
	t.Logf("%+v", Parse("-"))
	t.Logf("%+v", Parse("1"))
	t.Logf("%+v", Parse("-1"))
}
