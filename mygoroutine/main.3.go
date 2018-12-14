package main

import (
	"fmt"
	"runtime"
	"math/rand"
)


func main() {
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))
	fmt.Printf("%d\n",0x7FFFFFFF)
	fmt.Printf("%d\n",0xFFFFFFFFFFFF)
	fmt.Printf("GOMAXPROCS: %b %[1]d",runtime.GOMAXPROCS(0));
}

