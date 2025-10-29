package somedata

import (
	"slices"
	"unsafe"

	"github.com/eterline/somedata"
	"github.com/eterline/somedata/utils"
)

/*
matrix2 - default 2D matrix

for example:

	[1 2 3 4 5]
	[2 5 6 7 2]
	[3 3 4 6 5]
*/
type matrix2[T Numeric] struct {
	width  int
	height int
	arr    []T // flat data store
}

// NewMatrix2 - creates new 2D matrix
func NewMatrix2[T Numeric](width, height int) Matrix[T] {
	if width < 1 || height < 1 {
		panic(somedata.ErrMatNegativeCoords(2))
	}

	return &matrix2[T]{
		width:  width,
		height: height,
		arr:    make([]T, width*height),
	}
}

// index=y⋅width+x
func (mt *matrix2[T]) coords2idx(w, h int) int {
	if w >= mt.width || h >= mt.height {
		panic(somedata.ErrMatOutCoords(mt.Rank()))
	}

	return w*mt.width + h
}

func (mt *matrix2[T]) Rank() int {
	return 2
}

func (mt *matrix2[T]) Shape() []int {
	return []int{mt.width, mt.height}
}

func (mt *matrix2[T]) Size() int {
	return len(mt.arr)
}

func (mt *matrix2[T]) ShapeEquals(m Matrix[T]) bool {
	return shapeEq(mt, m)
}

func (mt *matrix2[T]) Get(coords ...int) T {
	if len(coords) != mt.Rank() {
		panic(somedata.ErrMatDimCoordMismatch(mt.Rank()))
	}

	idx := mt.coords2idx(coords[0], coords[1])
	return mt.arr[idx]
}

func (mt *matrix2[T]) Set(value T, coords ...int) {
	if len(coords) != mt.Rank() {
		panic(somedata.ErrMatDimCoordMismatch(mt.Rank()))
	}

	idx := mt.coords2idx(coords[0], coords[1])
	mt.arr[idx] = value
}

func (mt *matrix2[T]) Flatten() []T {
	return mt.arr
}

// clone - create new matrix with the same data values and sizes
func (mt *matrix2[T]) clone() *matrix2[T] {
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
	n := len(mt.arr)

	for i := 0; i < n; i++ {
		if i+8 < n {
			utils.Prefetch(unsafe.Pointer(&mt.arr[i+8]))
		}
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

func (mt *matrix2[T]) Add(m Matrix[T]) (Matrix[T], error) {
	if !mt.ShapeEquals(m) {
		return nil, somedata.ErrMatUnequalShapes(mt.Rank())
	}

	var (
		newMt = mt.clone()
		addMt = m.Flatten()
		n     = len(mt.arr)
	)

	for i := 0; i < n; i++ {
		if i+8 < n {
			utils.Prefetch(unsafe.Pointer(&mt.arr[i+8]))
		}
		newMt.arr[i] += addMt[i]
	}

	return newMt, nil
}

func (mt *matrix2[T]) Sub(m Matrix[T]) (Matrix[T], error) {
	if !mt.ShapeEquals(m) {
		return nil, somedata.ErrMatUnequalShapes(mt.Rank())
	}

	var (
		newMt = mt.clone()
		mFlat = m.Flatten()
		n     = len(mt.arr)
	)

	for i := 0; i < n; i++ {
		if i+8 < n {
			utils.Prefetch(unsafe.Pointer(&mt.arr[i+8]))
		}
		newMt.arr[i] -= mFlat[i]
	}

	return newMt, nil
}

func (mt *matrix2[T]) MulHadamard(m Matrix[T]) (Matrix[T], error) {
	if !mt.ShapeEquals(m) {
		return nil, somedata.ErrMatUnequalShapes(mt.Rank())
	}

	var (
		newMt = mt.clone()
		mFlat = m.Flatten()
		n     = len(mt.arr)
	)

	for i := 0; i < n; i++ {
		if i+8 < n {
			utils.Prefetch(unsafe.Pointer(&mt.arr[i+8]))
		}
		newMt.arr[i] *= mFlat[i]
	}

	return newMt, nil
}
