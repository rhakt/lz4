// cf. https://github.com/cloudflare/golz4

package lz4

/*
#cgo CFLAGS: -O3
#cgo LDFLAGS: -Wl,--allow-multiple-definition 
#include "lz4/src/lz4.h"
#include "lz4/src/lz4.c"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// p gets a char pointer to the first byte of a []byte slice
func p(in []byte) *C.char {
	if len(in) == 0 {
		return (*C.char)(unsafe.Pointer(nil))
	}
	return (*C.char)(unsafe.Pointer(&in[0]))
}

// clen gets the length of a []byte slice as a char *
func clen(s []byte) C.int {
	return C.int(len(s))
}

// Decompress : Decompress
func Decompress(in, out []byte) (outSize int, err error) {
	outSize = int(C.LZ4_decompress_safe(p(in), p(out), clen(in), clen(out)))
	if outSize < 0 {
		err = fmt.Errorf("fail to decompress")
	}
	return
}

// CompressBound : calculate max output size
func CompressBound(in []byte) int {
	return int(C.LZ4_compressBound(C.int(len(in))))
}

// Compress : compresses
func Compress(in, out []byte) (outSize int, err error) {
	outSize = int(C.LZ4_compress_limitedOutput(p(in), p(out), clen(in), clen(out)))
	if outSize <= 0 {
		err = fmt.Errorf("fail to compress")
	}
	return
}
