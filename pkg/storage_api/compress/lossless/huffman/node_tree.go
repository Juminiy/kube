package huffman

func (t *Tree) dfsPreWalk(fn visNode) {
	nd, stk := t.root, make(nodeStack, 0, t.n)
	for nd != nil || !stk.Empty() {
		for nd != nil {
			stk.Push(nd)
			fn(nd)
			nd = nd.left
		}
		if !stk.Empty() {
			nd = stk.Top()
			stk.Pop()
			nd = nd.right
		}
	}
}

func (t *Tree) dfsMidWalk(fn visNode) {
	nd, stk := t.root, make(nodeStack, 0, t.n)
	for nd != nil || !stk.Empty() {
		for nd != nil {
			stk.Push(nd)
			nd = nd.left
		}
		if !stk.Empty() {
			nd = stk.Top()
			stk.Pop()
			fn(nd)
			nd = nd.right
		}
	}
}

func (t *Tree) dfsPostWalk(fn visNode) {

}

func (t *Tree) bfsWalk(fn visNode) {
	que := make(nodeQueue, 0, t.n)
	que.Push(t.root)
	for !que.Empty() {
		nd := que.Front()
		que.Pop()
		fn(nd)
		if nd.left != nil {
			que.Push(nd.left)
		}
		if nd.right != nil {
			que.Push(nd.right)
		}
	}
}

func (t *Tree) bfsLevelWalk(fn visNode) {

}
