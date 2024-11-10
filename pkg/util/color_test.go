package util

import (
	"github.com/fatih/color"
	"testing"
)

func TestYN(t *testing.T) {
	t.Log(YN(false), YN(true))
	color.Green("%s", "YN")
}

func TestColorf(t *testing.T) {
	Colorf(
		ColorValue{Color: color.FgGreen, Value: "OK"},
		ColorValue{Color: color.FgCyan, Value: "OO"},
		ColorValue{Color: color.FgRed, Value: "NO"},
	)
}
