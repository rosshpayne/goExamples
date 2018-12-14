package main

import (
	"fmt"
)

type Buffer struct {
  	buf       []byte   // contents are the bytes buf[off : len(buf)]
  	off       int      // read at &buf[off], write at &buf[len(buf)]
  	bootstrap [64]byte // memory to hold first slice; helps small buffers avoid allocation.
  }

func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }

func main() {
         
    buf := NewBuffer(nil)
    fmt.Println("len,cap: ",len(buf.buf),cap(buf. buf))
    buf.buf=buf.bootstrap[0: ]
    fmt.Println("len,cap: ",len(buf.buf),cap(buf. buf))
}
