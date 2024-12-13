package file

import (
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	for _, s := range []string{"rrr", " rrr", "     rrr"} {
		t.Log(len(strings.TrimLeft(s, " ")))
	}
}
