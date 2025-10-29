package somedata

import (
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Matrix[T Numeric] interface {
	// Rank - matrix xD fomation
	Rank() int
	// Shape - matrix sizes from start to end
	Shape() []int
	// Size - fulll elements count
	Size() int
	// ShapeEquals - compares two matrix sizes
	ShapeEquals(m Matrix[T]) bool

	// Get - element by coordinates
	Get(coords ...int) T
	// Set - element by coordinates
	Set(value T, coords ...int)

	// Flatten - give flatten slice
	Flatten() []T
	// Scale - multipile on scalar value
	Scale(k T) Matrix[T]

	// Equals - compares two matrix by sizes and values
	Equals(m Matrix[T]) bool
	// Zero - set all matrix values by default value in type
	Zero()

	Add(m Matrix[T]) (Matrix[T], error)
	Sub(m Matrix[T]) (Matrix[T], error)
	MulHadamard(m Matrix[T]) (Matrix[T], error)
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

func shapeEq[T Numeric](a, b Matrix[T]) bool {
	aShape := a.Shape()
	bShape := b.Shape()

	if len(aShape) != len(bShape) {
		return false
	}

	for i := 0; i < len(aShape); i++ {
		if aShape[i] != bShape[i] {
			return false
		}
	}

	return true
}
