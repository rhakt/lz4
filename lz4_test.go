package lz4

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) {
	src, err := ioutil.ReadFile("testdata/jabberwocky.txt")
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	dst := make([]byte, CompressBound(src))
	dstSize, err := Compress(src, dst)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}
	if dstSize <= 0 {
		t.Fatal("Output buffer is empty.")
	}
	dst = dst[:dstSize]

	res := make([]byte, dstSize*10)
	resSize, err := Decompress(dst, res)
	if err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}
    if resSize <= 0 {
		t.Fatal("Output buffer is empty.")
	}
	res = res[:resSize]
    
    err = ioutil.WriteFile("testdata/jabberwocky_test.txt", res, os.ModePerm)
    if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	if string(src) != string(res) {
		t.Fatalf("Decompressed unmatch")
	}
}
