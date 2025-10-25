package somedata

type vocabulary[T comparable] struct {
	data map[T]struct{}
}

func NewVocabulary[T comparable]() *vocabulary[T] {
	return &vocabulary[T]{
		data: make(map[T]struct{}),
	}
}

func (v *vocabulary[T]) nilPanic() {
	if v == nil || v.data == nil {
		panic("vocabulary: is nil pointer")
	}
}

func (v *vocabulary[T]) Set(value T) {
	v.nilPanic()
	v.data[value] = struct{}{}
}

func (v *vocabulary[T]) Delete(value T) {
	v.nilPanic()
	delete(v.data, value)
}

func (v *vocabulary[T]) Exists(value T) bool {
	v.nilPanic()
	_, ok := v.data[value]
	return ok
}

func (v *vocabulary[T]) Intersects(voc *vocabulary[T]) (intersected []T, ok bool) {
	v.nilPanic()
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

func (v *vocabulary[T]) Slice() []T {
	v.nilPanic()
	slc := make([]T, 0, len(v.data))
	for value := range v.data {
		slc = append(slc, value)
	}
	return slc
}

func (v *vocabulary[T]) Size() int {
	v.nilPanic()
	return len(v.data)
}

func (v *vocabulary[T]) Clear() {
	v.nilPanic()
	for value := range v.data {
		delete(v.data, value)
	}
}
