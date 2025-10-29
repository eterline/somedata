package somedata

import "fmt"

type sErr string

func (e sErr) Error() string { return string(e) }

func newSErr(s string, p ...any) sErr {
	return sErr(fmt.Sprintf(s, p...))
}

// =============== Ring buffer errors ===============
const (
	ErrRingBufferInvalidSize sErr = "ring buffer: size must be above 0"
	ErrRingBufferOverflow    sErr = "ring buffer: overflow"
)

// =============== Deque buffer errors ===============
func ErrDequeEmptyQueue(fun string) sErr {
	return newSErr("deque: %s() called on empty queue", fun)
}

func ErrDequeCalledNegCount(fun string) sErr {
	return newSErr("deque: %s() called with negative count", fun)
}

func ErrDequeOutOfRange(r, l int) sErr {
	return newSErr("deque: index out of range %d with length %d", r, l)
}

const (
	ErrDequeDuringIter sErr = "deque: modified during iteration"
)

// =============== Tree BST errors ===============
const (
	ErrNilBstNode sErr = "tree BST: node is nil"
)

// =============== Matrix errors ===============
func ErrMatUnequalShapes(rank int) sErr {
	return newSErr("%dd matrix: not equal matrix shapes", rank)
}

func ErrMatDimCoordMismatch(rank int) sErr {
	return newSErr("%dd matrix: dimension coordinates mismatch", rank)
}

func ErrMatNegativeCoords(rank int) sErr {
	return newSErr("%dd matrix: dimensions must be above zero", rank)
}

func ErrMatOutCoords(rank int) sErr {
	return newSErr("%dd matrix: coords out of matrix range", rank)
}
