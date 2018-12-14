package main

import (
	"fmt"
	_ "math/rand"
	)
func main() {
    var a chan string 
    b := make(chan string)

    go func() { b <- "b" }()

    fmt.Printf("\na  %#v %[1]T" ,a)
    fmt.Printf("\nb  %#v %[1]T" ,b)

    select {
    case s := <-a:		// will block if a is nil. handy as you can switch a on and off.  a will always be a chan string, but nil means unallocated channel.
        fmt.Println("\ngot", s)
    case s := <-b:
        fmt.Println("\ngot", s)
    }

    a = make (chan string)

   go func()  { a <- "a" }()


    fmt.Printf("\na  %#v %[1]T" ,a)
    fmt.Printf("\nb  %#v %[1]T" ,b)
   select {
    case s := <-a:		
        fmt.Println("\ngot", s)
    case s := <-b:
        fmt.Println("\ngot", s)
   }
}
