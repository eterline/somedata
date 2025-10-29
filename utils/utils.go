package utils

import "unsafe"

func Sizeof[T any]() int {
	return int(unsafe.Sizeof(*new(T)))
}

func SizeofMul[T any](v int) int {
	return v * int(unsafe.Sizeof(*new(T)))
}
