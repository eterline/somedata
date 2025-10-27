package somedata

/* dict - dictionary is a data structure that stores key–value pairs.
It uses a hash function to map each key to an index in an internal array,
enabling average O(1) time complexity for lookup, insertion, and deletion.
*/
type dict[T comparable] struct {
	data map[T]struct{} // empty struct usage
}

// NewDict - creates dict
func NewDict[T comparable]() *dict[T] {
	return &dict[T]{
		data: make(map[T]struct{}),
	}
}

// Set - add element to dict
func (v *dict[T]) Set(value T) {
	v.data[value] = struct{}{}
}

// Delete - delete element from dict
func (v *dict[T]) Delete(value T) {
	delete(v.data, value)
}

// Exists - check value existing in dict
func (v *dict[T]) Exists(value T) bool {
	_, ok := v.data[value]
	return ok
}

// Intersects - returns intersection slice within two vocabularies
func (v *dict[T]) Intersects(voc *dict[T]) (intersected []T, ok bool) {
	slc := make([]T, 0)

	for value := range v.data {
		if voc.Exists(value) {
			slc = append(slc, value)
		}
	}

	if len(slc) > 0 {
		return slc, true
	}

	return nil, false
}

// Slice - voc as a Slice
func (v *dict[T]) Slice() []T {
	slc := make([]T, 0, len(v.data))
	for value := range v.data {
		slc = append(slc, value)
	}
	return slc
}

// Size - dict elements count
func (v *dict[T]) Size() int {
	return len(v.data)
}

// Clear - fully clears all elements
func (v *dict[T]) Clear() {
	for value := range v.data {
		delete(v.data, value)
	}
}
