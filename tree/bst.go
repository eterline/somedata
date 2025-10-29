package somedata

import (
	"github.com/eterline/somedata"
	"golang.org/x/exp/constraints"
)

type NodeBST[T constraints.Ordered, D any] interface {
	Value() T
	Data() D
}

type nodeBST[T constraints.Ordered, D any] struct {
	value T
	data  D
	low   *nodeBST[T, D]
	hight *nodeBST[T, D]
}

func (n *nodeBST[T, D]) Value() T {
	if n == nil {
		panic(somedata.ErrNilBstNode)
	}
	return n.value
}

func (n *nodeBST[T, D]) Data() D {
	if n == nil {
		panic(somedata.ErrNilBstNode)
	}
	return n.data
}

func (n *nodeBST[T, D]) rm(value T) (*nodeBST[T, D], bool) {
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

func (t *nodeBST[T, D]) min() T {
	if t.low == nil {
		return t.value
	}
	return t.low.min()
}

func (n *nodeBST[T, D]) minNode() *nodeBST[T, D] {
	cur := n
	for cur.low != nil {
		cur = cur.low
	}
	return cur
}

func (t *nodeBST[T, D]) max() T {
	if t.hight == nil {
		return t.value
	}
	return t.hight.max()
}

func (t *nodeBST[T, D]) inOrder(fn func(T)) {
	if t == nil {
		return
	}

	fn(t.value)
	t.low.inOrder(fn)
	t.hight.inOrder(fn)
}

func insertNode[T constraints.Ordered, D any](node *nodeBST[T, D], value T, data D) *nodeBST[T, D] {
	switch {
	case node == nil:
		return &nodeBST[T, D]{value: value, data: data}

	case value < node.value:
		node.low = insertNode(node.low, value, data)

	case value > node.value:
		node.hight = insertNode(node.hight, value, data)
	}

	return node
}

type threeBST[T constraints.Ordered, D any] struct {
	size int
	root *nodeBST[T, D]
}

func NewThreeBST[T constraints.Ordered, D any]() *threeBST[T, D] {
	return &threeBST[T, D]{}
}

func (t *threeBST[T, D]) Size() int {
	return t.size
}

func (t *threeBST[T, D]) Insert(value T, data D) {
	t.root = insertNode(t.root, value, data)
	t.size++
}

func (t *threeBST[T, D]) Min() (T, bool) {
	if t.size == 0 {
		var zero T
		return zero, false
	}
	return t.root.min(), true
}

func (t *threeBST[T, D]) Max() (T, bool) {
	if t.size == 0 {
		var zero T
		return zero, false
	}
	return t.root.max(), true
}

func (t *threeBST[T, D]) InOrder(fn func(T)) {
	if t.size == 0 {
		return
	}
	t.root.inOrder(fn)
}

func (t *threeBST[T, D]) Delete(value T) (ok bool) {
	if t.size == 0 {
		return false
	}

	t.root, ok = t.root.rm(value)
	if ok {
		t.size--
	}

	return ok
}
