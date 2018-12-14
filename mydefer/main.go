package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a")) // trace() executes immediately, un() gets deferred
	fmt.Println("in a")
}

func b() {
	defer un(trace("b")) // trace() executes immediately, un() gets deferred.
	fmt.Println("in b")
	a()
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	b()
}
