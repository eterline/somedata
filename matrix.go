package somedata

import (
	"cmp"
	"slices"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Matrix[T cmp.Ordered] interface {
	Rank() int    // Matrix xD fomation
	Shape() []int // Matrix sizes from start to end
	Size() int    // Fullly elements count

	Get(coords ...int) T        // Get element by Coordinates
	Set(value T, coords ...int) // Set element by Coordinates

	Flatten() []T        // Give flatten slice
	Scale(k T) Matrix[T] // Multipile on scalar value

	// Add(m Matrix[T]) (Matrix[T], error)
	// Sub(m Matrix[T]) (Matrix[T], error)
	// Mul(m Matrix[T]) (Matrix[T], error) // Элементное умножение (Hadamard)

	Equals(m Matrix[T]) bool
	Zero()
}

func coords2Idx(shape, coords []int) int {
	if len(shape) != len(coords) {
		panic("matrix: dimension mismatch")
	}

	index := 0
	stride := 1

	for i := len(shape) - 1; i >= 0; i-- {
		if coords[i] < 0 || coords[i] >= shape[i] {
			panic("matrix: index out of range")
		}
		index += coords[i] * stride
		stride *= shape[i]
	}
	return index
}

func idx2Coords(shape []int, index int) []int {
	coords := make([]int, len(shape))
	for i := len(shape) - 1; i >= 0; i-- {
		coords[i] = index % shape[i]
		index /= shape[i]
	}
	return coords
}

// ==============================

// matrix2 - default 2D matrix
type matrix2[T Numeric] struct {
	width  int
	height int
	arr    []T // flat data store
}

func NewMatrix2[T Numeric](width, height int) Matrix[T] {
	return &matrix2[T]{
		width:  width,
		height: height,
		arr:    make([]T, width*height),
	}
}

func (mt *matrix2[T]) nilPanic() {
	if mt == nil || mt.arr == nil {
		panic("2D matrix: is nil pointer")
	}
}

func (mt *matrix2[T]) coords2idx(w, h int) int {
	if w == 0 {
		w = 1
	}
	if h == 0 {
		h = 1
	}
	return (w - 1) + (h-1)*w
}

func (mt *matrix2[T]) Rank() int {
	return 2
}

func (mt *matrix2[T]) Shape() []int {
	mt.nilPanic()
	return []int{mt.width, mt.height}
}

func (mt *matrix2[T]) Size() int {
	mt.nilPanic()
	return len(mt.arr)
}

func (mt *matrix2[T]) Get(coords ...int) T {
	mt.nilPanic()

	if len(coords) > mt.Rank() {
		panic("2D matrix: dimension coordinates mismatch")
	}

	idx := mt.coords2idx(coords[0], coords[1])
	return mt.arr[idx]
}

func (mt *matrix2[T]) Set(value T, coords ...int) {
	mt.nilPanic()

	if len(coords) > mt.Rank() {
		panic("2D matrix: dimension coordinates mismatch")
	}

	idx := mt.coords2idx(coords[0], coords[1])
	mt.arr[idx] = value
}

func (mt *matrix2[T]) Flatten() []T {
	mt.nilPanic()
	return mt.arr
}

func (mt *matrix2[T]) clone() *matrix2[T] {
	mt.nilPanic()

	new := &matrix2[T]{
		width:  mt.width,
		height: mt.height,
		arr:    make([]T, mt.width*mt.height),
	}

	copy(new.arr, mt.arr)
	return new
}

func (mt *matrix2[T]) Clone() Matrix[T] {
	return mt.clone()
}

func (mt *matrix2[T]) Scale(k T) Matrix[T] {
	cloned := mt.clone()
	for i := range mt.arr {
		cloned.arr[i] *= k
	}
	return cloned
}

func (mt *matrix2[T]) Equals(m Matrix[T]) bool {
	flt := m.Flatten()
	return slices.Equal(mt.arr, flt)
}

func (mt *matrix2[T]) Zero() {
	var dflt T
	for i := range mt.arr {
		mt.arr[i] = dflt
	}
}
