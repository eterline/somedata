package somedata

import "unsafe"

//go:noescape
func prefetch(addr unsafe.Pointer)
