package huffman

type nodeQueue []*node

func (l nodeQueue) Len() int {
	return len(l)
}

func (l nodeQueue) Empty() bool {
	return len(l) == 0
}

func (l *nodeQueue) Push(nd *node) {
	*l = append(*l, nd)
}

func (l *nodeQueue) Pop() {
	if n := len(*l); n > 0 {
		*l = (*l)[1:]
	}
}

func (l nodeQueue) Front() *node {
	if n := len(l); n > 0 {
		return l[0]
	}
	return nil
}

func (l nodeQueue) Back() *node {
	if n := len(l); n > 0 {
		return l[n-1]
	}
	return nil
}
