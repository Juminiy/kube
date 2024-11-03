package runtimeutil

import "unsafe"

// *T <----> unsafe.Pointer ----> uintptr

// Tp2Up
// Type Pointer to unsafe.Pointer
func Tp2Up[T any](tp *T) unsafe.Pointer {
	return unsafe.Pointer(tp)
}

// Up2Tp
// unsafe.Pointer to Type Pointer
func Up2Tp[T any](up unsafe.Pointer) *T {
	return (*T)(up)
}

// Up2Ui
// unsafe.Pointer to uintptr
func Up2Ui(up unsafe.Pointer) uintptr {
	return uintptr(up)
}
