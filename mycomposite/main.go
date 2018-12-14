package main

import (
       "fmt"
)

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

const (
      Enone int =iota
      Eio
      Einval
)

func main () {

fmt.Println(Enone)
fmt.Println(Eio)
fmt.Println(Einval)
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

for i,v:=range a {
	fmt.Printf("%d %s\n",i,v)
}

for i,v:= range s {
	fmt.Printf("%v %s\n",i,v)
}

for i,v:= range m {
	fmt.Printf("%v %s\n",i,v)
}
fmt.Printf("%032b  %f\n",int32(KB),KB)
fmt.Printf("%032b  %f\n",int32(MB),MB)
fmt.Printf("%032b  %f\n",int32(GB),GB)
fmt.Printf("%064b  %f\n",int64(EB),EB)
fmt.Printf("%064b  %f\n",int64(EB),YB)
}
