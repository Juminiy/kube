package util

import (
	"fmt"
	"testing"
)

func TestVar1(t *testing.T) {
	a, b := 1, 2
	b, c := 3, 4
	{
		c, d := 5, 6
		_, _, _, _ = a, b, c, d
	}
	t.Log(c) // 4
}

func TestVar2(t *testing.T) {
	last := []byte{2, 0, 2, 4}
	this := append(last[:3], 5)
	t.Log(last, this) // [2 0 2 5] [2 0 2 5]
}

type sFile string

func sOpen(s string) sFile {
	return sFile(s)
}

func (s sFile) Close() {
	fmt.Println(s)
}

func TestSFile(t *testing.T) {
	f := sOpen("1")
	defer f.Close()
	f.Close() // 1
	f = sOpen("2")
	// 1
}
