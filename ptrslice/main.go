// Package ptrslice is a utility for converting pointers into
// single-element slices referring to the same underlying memory.
//
// This package uses unsafe to do its work, so it may be broken on later
// versions of Go.
//
// This if you are using this package then you are probably doing something
// wrong. The possibly-legitimate use-case is passing a single instance
// of a huge value (that would be expensive to copy) to a function that
// deals in slices, but if the function can be modified to take a slice
// of pointers instead then that is a superior solution.
package ptrslice

import (
	"reflect"
	"unsafe"
)

// PointerToSlice takes a value that must be a pointer and returns a value
// that is a slice of the given value's pointer type with length 1 and
// capacity 1.
//
// For example:
//
//     i := 1
//     sl := ptrslice.PointerToSlice(&i).([]int)
//     // sl[i] is the same 1 as i .
func PointerToSlice(ptr interface{}) interface{} {
	ptrVal := reflect.ValueOf(ptr)
	ptrTy := reflect.TypeOf(ptr)
	if ptrTy.Kind() != reflect.Ptr {
		panic("given value is not a pointer")
	}

	elemTy := ptrTy.Elem()
	addr := ptrVal.Elem().UnsafeAddr()
	hdr := reflect.SliceHeader{
		Data: addr,
		Len:  1,
		Cap:  1,
	}
	slicePtrVal := reflect.NewAt(reflect.SliceOf(elemTy), unsafe.Pointer(&hdr))

	// Dummy extra reference to ptrVal after we've built slicePtrVal
	// to ensure that our original value doesn't get collected before
	// we get a chance to return.
	if ptrVal.Interface() == nil {
		// should never happen.
		panic("ptrVal.Interface() returned nil")
	}

	return slicePtrVal.Elem().Interface()
}
