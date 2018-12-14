package testvar

import (
	"fmt"
)

func Func1 (d ... int) {
        if d == nil {
	    fmt.Println("d is nil")
	} else {
	    fmt.Println("d is not nil")
	}
	fmt.Println("len,cap: ",len(d), cap(d))
	for i,v := range d {
		fmt.Println(i,v)
	}
}

func Func2 (d ... interface{} ) {
	fmt.Println("*****************")
        if d == nil {
	    fmt.Println("d is nil")
	} else {
	    fmt.Println("d is not nil")
	}
	fmt.Println("len,cap: ",len(d), cap(d))

	for _,v := range d {
	   if _, ok := v.(chan struct{}); ok  {
		fmt.Println("done is a chan struct{}")
		continue
	   }
	   if _, ok := v.(chan int); ok  {
		fmt.Println("done is a chan int")
		continue
	   }
	   fmt.Printf("\nType: %T",v)	  
	}
}

