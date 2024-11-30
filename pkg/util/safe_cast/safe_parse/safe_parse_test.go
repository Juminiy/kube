package safe_parse

import (
	"testing"
)

func TestParse(t *testing.T) {
	var readV Type
	readStringV := func(s ...string) {
		for _, stringV := range s {
			readV = Parse(stringV)
			t.Log(readV)
		}
	}

	// test bool
	readStringV("1", "0", "True", "False", "true", "false")

	// test number
	readStringV("666", "-666", "666.666", "-666.666")

	// test time.Time
	readStringV("2024-11-30 14:49:01", "2021-11-21 18:32:02")

	// test invalid
	readStringV(
		"tRue",                     // invalid bool
		"+-666",                    // invalid number
		"2024-11-30$14:49:01",      // invalid time.Time
		"v2ray", "k8s.io", "$666#") // invalid others
}
