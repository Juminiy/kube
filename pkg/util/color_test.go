package util

import (
	"github.com/fatih/color"
	"testing"
)

func TestYN(t *testing.T) {
	t.Log(YN(false), YN(true))
	color.Green("%s", "YN")
}
