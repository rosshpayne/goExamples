package main

import (
	rt "runtime"
	"fmt"
)

func  main () {
	fmt.Printf("NumCPU: %d \n",rt.NumCPU())
	fmt.Printf("NumGoroutine: %d \n",rt.NumGoroutine())
	fmt.Printf("GOMAXPROCS: %d \n",rt.GOMAXPROCS(4))
}
