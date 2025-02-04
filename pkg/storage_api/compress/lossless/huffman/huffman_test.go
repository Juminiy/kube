package huffman

import "testing"

func TestNew(t *testing.T) {
	huf := New([]rune{'F', 'O', 'R', 'G', 'E', 'T'},
		[]int{2, 3, 4, 4, 5, 7})
	for _, sym := range []rune{
		'F', 'O', 'R', 'G', 'E', 'T',
		'f', 'o', 'r', 'g', 'e', 't'} {
		t.Log(huf.Get(sym))
	}
}
