package main

import (
	"fmt"
)
// this exercise is to prove that len(<slice>) is a property of the slice not the underly array
//   whereas cap(<slice>) is a property of the underlying array

// append(..) will append to the end of the slice so two slices with the same underlying array may append to different parts of the array

var anotherbit = []int{-1,-2,-3,-4,-5}

func main() {
        //
	var aa []int 
        aa = make([]int,0,100)
	fmt.Printf("aa = %d %d %#v\n",len(aa),cap(aa),aa)
	for i:=0;i<412;i++ {
	     aa=append(aa,i)
	}	
	fmt.Printf("aa2 = %d %d %#v\n",len(aa),cap(aa),aa)
        extabit(aa[100:122])
	fmt.Printf("aa2 = %d %d %#v\n",len(aa),cap(aa),aa)
	
}

func extabit ( aslice []int ) {
	fmt.Printf("aslice a= %d %d %#v\n",len(aslice),cap(aslice),aslice)
        fmt.Printf("alsice[22]=%d\n",aslice[len(aslice)-1])

	bb:=append(aslice,anotherbit...)

	fmt.Printf("aslice a= %d %d %#v\n",len(bb),cap(bb),bb)
	fmt.Printf("aslice a= %d %d %#v\n",len(aslice),cap(aslice),aslice)

	aslice=append(bb,anotherbit...)

	fmt.Printf("aslice a= %d %d %#v\n",len(aslice),cap(aslice),aslice)
}

