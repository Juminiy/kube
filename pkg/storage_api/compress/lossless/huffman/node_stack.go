package huffman

type nodeStack []*node

func (l nodeStack) Len() int {
	return len(l)
}

func (l nodeStack) Empty() bool {
	return len(l) == 0
}

func (l *nodeStack) Push(nd *node) {
	*l = append(*l, nd)
}

func (l *nodeStack) Pop() {
	if n := len(*l); n > 0 {
		*l = (*l)[:n-1]
	}
}

func (l nodeStack) Top() *node {
	if n := len(l); n > 0 {
		return l[n-1]
	}
	return nil
}
