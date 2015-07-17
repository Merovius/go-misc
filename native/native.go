// Package native provides a host-native binary.ByteOrder.
//
// This package can be used to communicate with badly written programs running
// on the same host which use the native byte order of the architecture it's
// running on. It uses unsafe and probably doesn't work on architectures, that
// don't have a constant byte order.
//
// See http://commandcenter.blogspot.fr/2012/04/byte-order-fallacy.html for why
// you shouldn't use this package unless you deal with badly written programs.
package native

import (
	"reflect"
	"unsafe"
)

// Native is a host-native binary.ByteOrder.
var Native native

type native struct{}

func (native) Uint16(b []byte) uint16 {
	_, _ = b[0], b[1]
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*uint16)(unsafe.Pointer(h.Data))
}

func (native) Uint32(b []byte) uint32 {
	_, _ = b[0], b[3]
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*uint32)(unsafe.Pointer(h.Data))
}

func (native) Uint64(b []byte) uint64 {
	_, _ = b[0], b[7]
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*uint64)(unsafe.Pointer(h.Data))
}

func (native) PutUint16(b []byte, i uint16) {
	h := reflect.SliceHeader{
		Len:  2,
		Cap:  2,
		Data: uintptr(unsafe.Pointer(&i)),
	}
	copy(b, *(*[]byte)(unsafe.Pointer(&h)))
}

func (native) PutUint32(b []byte, i uint32) {
	h := reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(&i)),
	}
	copy(b, *(*[]byte)(unsafe.Pointer(&h)))
}

func (native) PutUint64(b []byte, i uint64) {
	h := reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(&i)),
	}
	copy(b, *(*[]byte)(unsafe.Pointer(&h)))
}

func (native) String() string {
	return "Native"
}

func (native) GoString() string {
	return "native.Native"
}
