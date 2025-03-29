package huffman

type nodeHeap []*node

func (l nodeHeap) Len() int {
	return len(l)
}

func (l nodeHeap) Less(i, j int) bool {
	if l[i].weight == l[j].weight {
		if l[i].leaf && l[j].leaf {
			return true
		} else if l[i].leaf {
			return true
		} else if l[j].leaf {
			return false
		} else {
			return l[i].leafs < l[j].leafs
		}
	}
	return l[i].weight < l[j].weight
}

func (l nodeHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *nodeHeap) Push(x any) {
	*l = append(*l, x.(*node))
}

func (l *nodeHeap) Pop() any {
	old, n := *l, len(*l)
	if n > 0 {
		item := old[n-1]
		old[n-1] = nil
		*l = old[0 : n-1]
		return item
	}
	return (*node)(nil)
}
