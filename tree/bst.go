package somedata

import "golang.org/x/exp/constraints"

type nodeBST[T constraints.Ordered] struct {
	value T
	low   *nodeBST[T]
	hight *nodeBST[T]
}

func (n *nodeBST[T]) rm(value T) (*nodeBST[T], bool) {
	if n == nil {
		return nil, false
	}

	if value < n.value {
		var deleted bool
		n.low, deleted = n.low.rm(value)
		return n, deleted
	}
	if value > n.value {
		var deleted bool
		n.hight, deleted = n.hight.rm(value)
		return n, deleted
	}

	if n.low == nil && n.hight == nil {
		return nil, true
	}
	if n.low == nil {
		return n.hight, true
	}
	if n.hight == nil {
		return n.low, true
	}

	successor := n.hight.minNode()
	n.value = successor.value
	var deleted bool
	n.hight, deleted = n.hight.rm(successor.value)
	return n, deleted
}

func (t *nodeBST[T]) min() T {
	if t.low == nil {
		return t.value
	}
	return t.low.min()
}

func (n *nodeBST[T]) minNode() *nodeBST[T] {
	cur := n
	for cur.low != nil {
		cur = cur.low
	}
	return cur
}

func (t *nodeBST[T]) max() T {
	if t.hight == nil {
		return t.value
	}
	return t.hight.max()
}

func (t *nodeBST[T]) inOrder(fn func(T)) {
	if t.low != nil {
		t.low.inOrder(fn)
	}

	fn(t.value)

	if t.hight != nil {
		t.hight.inOrder(fn)
	}
}

func insertNode[T constraints.Ordered](node *nodeBST[T], value T) *nodeBST[T] {
	switch {
	case node == nil:
		return &nodeBST[T]{value: value}

	case value < node.value:
		node.low = insertNode(node.low, value)

	case value > node.value:
		node.hight = insertNode(node.hight, value)
	}

	return node
}

type threeBST[T constraints.Ordered] struct {
	size int
	root *nodeBST[T]
}

func NewThreeBST[T constraints.Ordered]() *threeBST[T] {
	return &threeBST[T]{}
}

func (t *threeBST[T]) Size() int {
	return t.size
}

func (t *threeBST[T]) Insert(value T) {
	t.root = insertNode(t.root, value)
	t.size++
}

func (t *threeBST[T]) Min() (T, bool) {
	if t.size == 0 {
		var zero T
		return zero, false
	}
	return t.root.min(), true
}

func (t *threeBST[T]) Max() (T, bool) {
	if t.size == 0 {
		var zero T
		return zero, false
	}
	return t.root.max(), true
}

func (t *threeBST[T]) InOrder(fn func(T)) {
	if t.size == 0 {
		return
	}
	t.root.inOrder(fn)
}

func (t *threeBST[T]) Delete(value T) (ok bool) {
	if t.size == 0 {
		return false
	}

	t.root, ok = t.root.rm(value)
	if ok {
		t.size--
	}
	return ok
}
