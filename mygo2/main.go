package main

import (
       "fmt"
       "bytes"
       "sync"
       )

type ByteSize float64

const (
    KB ByteSize = 1 << (10 * 1)
    MB ByteSize = 1 << (10 * 2)
    GB ByteSize = 1 << (10 * 3)
)

type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
//
// internal package variables
//
//var pv &SyncedBuffer     // doesn't work &
var v SyncedBuffer      // type  SyncedBuffer - allocates memory & initialises with zero values.
var  ip = new(SyncedBuffer)  // type *SyncedBuffer

func main () {
 pv := new(SyncedBuffer)  // type *SyncedBuffer
 pv2 := &SyncedBuffer{}   // example of a composite literal: see mygo/main.go -- must initialise the struct e.g T{<initialise fields>} 
//
//a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
//s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
//m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
//
 fmt.Printf("%T\n", ip);
 fmt.Printf("%T\n", v);
 fmt.Printf("%T\n", pv);
 fmt.Printf("%T\n", pv2);

 fmt.Println(KB);
 fmt.Println(MB);
 fmt.Println(GB);
}
