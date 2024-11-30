package safe_parse

import "testing"

func TestParseNumber(t *testing.T) {
	t.Logf("%+v", ParseNumber(""))
	t.Logf("%+v", ParseNumber("v"))
	t.Logf("%+v", ParseNumber("-"))
	t.Logf("%+v", ParseNumber("1"))
	t.Logf("%+v", ParseNumber("-1"))
}
