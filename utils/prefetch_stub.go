package utils

import "unsafe"

// Prefetch - ASM realization of CPU cache fetching ahead compute
//go:noescape
func Prefetch(addr unsafe.Pointer)
