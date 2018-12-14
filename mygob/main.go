package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "log"
    //"os"
)

type P struct {
    X, Y, Z int
    D  *int
    Name    string
}

type Q struct {         // desitination fields types do not need to match
    //X,Y int           // does not need to match source e.g  X, Y *int
    X,Y *int32           // does not need to match source e.g  X, Y *int
    //D    *int
    D    int
    Name string
}

func main() {
    // Initialize the encoder and decoder.  Normally enc and dec would be
    // bound to network connections and the encoder and decoder would
    // run in different processes.
    var network bytes.Buffer        // Stand-in for a network connection
    enc := gob.NewEncoder(&network) // Will write to network.
    dec := gob.NewDecoder(&network) // Will read from network.
    // Encode (send) the value.
    var inputint int = 44
    err := enc.Encode(P{3, 4, 5, &inputint  , "Pythagoras"})
    //np:=&network
    //fmt.Println("len: ",np.Len())
    //fmt.Println("cap: ",np.Cap())
    //np.WriteTo(os.Stdout)
    if err != nil {
        log.Fatal("encode error:", err)
    }
    // Decode (receive) the value.
    var q Q
    err = dec.Decode(&q)
    if err != nil {
        log.Fatal("decode error:", err)
    }
    //var str string
    //fmt.Printf("%q: {%d,%d} %d\n", q.Name, *(q.X), *q.Y,*q.D)
    fmt.Printf("%q: {%d,%d} %d\n", q.Name, *(q.X), *q.Y,q.D)
    //fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)
}
