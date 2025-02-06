package huffman

import (
	"container/heap"
)

type Tree struct {
	root *node
	n    int

	encoder map[rune]string
	decoder map[string]rune
}

// New
// TODO: BUG fix
func New(symbol []rune, weight []int) *Tree {
	n := len(symbol)
	if n != len(weight) {
		return nil
	}
	nodes := make(nodeHeap, len(symbol))
	for idx := range symbol {
		nodes[idx] = newLeaf(symbol[idx], weight[idx])
	}
	heap.Init(&nodes)

	for nodes.Len() > 1 {
		left, right := heap.Pop(&nodes).(*node), heap.Pop(&nodes).(*node)
		inter := newInter(left, right)
		left.parent, right.parent = inter, inter
		heap.Push(&nodes, inter)
	}

	nodes[0].root = true
	t := &Tree{
		root:    nodes[0],
		n:       n,
		encoder: make(map[rune]string, n),
		decoder: make(map[string]rune, n),
	}
	t.dfsPreWalk(setCode)
	t.dfsPreWalk(func(cur *node) {
		if cur.leaf {
			t.encoder[cur.symbol] = cur.code
			t.decoder[cur.code] = cur.symbol
		}
	})
	return t
}

// Lookup
// binary representation to symbol
func (t *Tree) Lookup(bin string) (rune, bool) {
	sym, ok := t.decoder[bin]
	return sym, ok
}

// Get
// symbol to binary representation
func (t *Tree) Get(sym rune) (string, bool) {
	bin, ok := t.encoder[sym]
	return bin, ok
}

// String
// print tree ASCII TEXT
func (t *Tree) String() string {
	return ``
}
