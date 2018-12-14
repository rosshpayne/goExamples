package main

import (
	"fmt"
        "sync"
)


func main() {

	var x,y int
	var wg sync.WaitGroup

        wg.Add(2)
	go func() {
		x=1
		fmt.Print("y: ",y, " ")
		wg.Done()
	}()

	go func() {
		y=1
		fmt.Print("x: ",x, " ")
		wg.Done()
	}()
        wg.Wait() 
}
