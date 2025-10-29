package somedata

import (
	"slices"

	"github.com/eterline/somedata"
)

// matrix3 - 3D matrix
type matrix3[T Numeric] struct {
	width  int
	height int
	deep   int
	arr    []T // flat data store
}

// NewMatrix3 - creates new 3D matrix
func NewMatrix3[T Numeric](width, height, deep int) *matrix3[T] {
	if width < 1 || height < 1 || deep < 1 {
		panic(somedata.ErrMatNegativeCoords(3))
	}

	return &matrix3[T]{
		width:  width,
		height: height,
		deep:   deep,
		arr:    make([]T, width*height*deep),
	}
}

// index = (z*height+y)*width+x
func (mt *matrix3[T]) coords2idx(width, height, deep int) int {
	if width >= mt.width || height >= mt.height || deep >= mt.deep {
		panic(somedata.ErrMatOutCoords(mt.Rank()))
	}

	return (deep*mt.height+height)*mt.width + width
}

func (mt *matrix3[T]) Rank() int {
	return 3
}

func (mt *matrix3[T]) Shape() []int {
	return []int{mt.width, mt.height, mt.deep}
}

func (mt *matrix3[T]) Size() int {
	return len(mt.arr)
}

func (mt *matrix3[T]) ShapeEquals(m Matrix[T]) bool {
	return shapeEq(mt, m)
}

func (mt *matrix3[T]) Get(coords ...int) T {
	if len(coords) != mt.Rank() {
		panic(somedata.ErrMatDimCoordMismatch(mt.Rank()))
	}

	idx := mt.coords2idx(coords[0], coords[1], coords[2])
	return mt.arr[idx]
}

func (mt *matrix3[T]) Set(value T, coords ...int) {
	if len(coords) != mt.Rank() {
		panic(somedata.ErrMatDimCoordMismatch(mt.Rank()))
	}

	idx := mt.coords2idx(coords[0], coords[1], coords[2])
	mt.arr[idx] = value
}

func (mt *matrix3[T]) Flatten() []T {
	return mt.arr
}

// clone - create new matrix with the same data values and sizes
func (mt *matrix3[T]) clone() *matrix3[T] {
	new := &matrix3[T]{
		width:  mt.width,
		height: mt.height,
		deep:   mt.deep,
		arr:    make([]T, mt.width*mt.height*mt.deep),
	}

	copy(new.arr, mt.arr)
	return new
}

func (mt *matrix3[T]) Clone() Matrix[T] {
	return mt.clone()
}

func (mt *matrix3[T]) Scale(k T) Matrix[T] {
	cloned := mt.clone()
	for i := range mt.arr {
		cloned.arr[i] *= k
	}
	return cloned
}

func (mt *matrix3[T]) Equals(m Matrix[T]) bool {
	flt := m.Flatten()
	return slices.Equal(mt.arr, flt)
}

func (mt *matrix3[T]) Zero() {
	var dflt T
	for i := range mt.arr {
		mt.arr[i] = dflt
	}
}

func (mt *matrix3[T]) Add(m Matrix[T]) (Matrix[T], error) {
	if !mt.ShapeEquals(m) {
		return nil, somedata.ErrMatUnequalShapes(mt.Rank())
	}

	newMt := mt.clone()
	addMt := m.Flatten()

	for i := range newMt.arr {
		newMt.arr[i] += addMt[i]
	}

	return newMt, nil
}

func (mt *matrix3[T]) Sub(m Matrix[T]) (Matrix[T], error) {
	if !mt.ShapeEquals(m) {
		return nil, somedata.ErrMatUnequalShapes(mt.Rank())
	}

	newMt := mt.clone()
	mFlat := m.Flatten()

	for i := range newMt.arr {
		newMt.arr[i] -= mFlat[i]
	}

	return newMt, nil
}

func (mt *matrix3[T]) MulHadamard(m Matrix[T]) (Matrix[T], error) {
	if !mt.ShapeEquals(m) {
		return nil, somedata.ErrMatUnequalShapes(mt.Rank())
	}

	newMt := mt.clone()
	mFlat := m.Flatten()

	for i := range newMt.arr {
		newMt.arr[i] *= mFlat[i]
	}

	return newMt, nil
}
