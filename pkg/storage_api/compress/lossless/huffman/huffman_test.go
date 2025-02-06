package huffman

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	huf := New([]rune{'F', 'O', 'R', 'G', 'E', 'T'},
		[]int{2, 3, 4, 4, 5, 7})
	//for _, sym := range []rune{
	//	'F', 'O', 'R', 'G', 'E', 'T',
	//	'f', 'o', 'r', 'g', 'e', 't'} {
	//	bin, ok := huf.Get(sym)
	//	if ok {
	//		t.Logf("%c -> %s", sym, bin)
	//	} else {
	//		t.Logf("%c not found", sym)
	//	}
	//}

	tVis := func(nd *node) {
		fmt.Printf("%d ", nd.weight)
	}
	tsVis := func(nds []*node) {
		fmt.Print("[")
		slices.All(nds)(func(_ int, nd *node) bool {
			tVis(nd)
			return true
		})
		fmt.Print("]")
	}
	huf.dfsPreWalk(tVis)
	util.TestLongHorizontalLine(t)
	huf.dfsMidWalk(tVis)
	util.TestLongHorizontalLine(t)
	huf.dfsPostWalk(tVis)
	util.TestLongHorizontalLine(t)
	huf.bfsWalk(tVis)
	util.TestLongHorizontalLine(t)
	huf.bfsLevelWalk(tsVis)
	util.TestLongHorizontalLine(t)
}
