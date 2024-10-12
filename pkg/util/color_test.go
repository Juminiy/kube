package util

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

func TestYN(t *testing.T) {
	fmt.Println(YN(false), YN(true))
	color.Green("%s", "YN")
}
