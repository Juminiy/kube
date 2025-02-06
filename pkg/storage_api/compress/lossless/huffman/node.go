package huffman

type node struct {
	left   *node
	right  *node
	parent *node

	symbol rune
	weight int
	code   string

	leafs int
	leaf  bool
	root  bool
}

func newLeaf(symbol rune, weight int) *node {
	return &node{
		symbol: symbol,
		weight: weight,
		leafs:  1,
		leaf:   true,
		root:   false,
	}
}

func newInter(left, right *node) *node {
	return &node{
		left:   left,
		right:  right,
		weight: left.weight + right.weight,
		leafs:  left.leafs + right.leafs,
		leaf:   false,
		root:   false,
	}
}

type visNode func(cur *node)

func setCode(cur *node) {
	if !cur.root {
		if cur.parent.left == cur {
			cur.code = cur.parent.code + "0"
		} else {
			cur.code = cur.parent.code + "1"
		}
	}
}

type visNodeList func(cur []*node)
