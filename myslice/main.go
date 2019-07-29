package main

import (
	"fmt"
)

var slice1 []int = []int{1, 2, 3, 4, 5, 6, 7}
var slice1_ []int = []int{10: 21, 22, 23, 12, 25, 26} // specify index rather than default to 0.  Go will automatically increment from initial value.

func main() {
	fmt.Println("slice1: ", slice1)
	//
	// . slice operator s1:=X[x:y]
	//    following slice operator assigns following metadata to s1
	//
	//    s1.len = y-x              e.g.  [2:5] len=3,  [:5] len=5 === [0:5], indices 0,1,2,3,4
	//    s1.cap = cap(X)-x
	//    s1.data = &X.data[x]         s1 points to start of its slice in underlying array
	//
	s1 := slice1[2:3] // s1 is infered to be a slice of int of len 2
	fmt.Printf("\n 10 s1=slice1[2:3] len %d cap  %d  s1 %v\n", len(s1), cap(s1), s1)
	//
	// reslice s1 - extend beyond current len - increases s1 len from 1 to 3 but cap remains the same
	//
	s1 = s1[:3]
	//
	fmt.Printf("\n 11 s1[:3]] len %d cap  %d  s1 %v\n", len(s1), cap(s1), s1)
	fmt.Printf("\n 20  s1[:] len %d cap  %d  s1 %v\n", len(s1), cap(s1), s1[:])
	fmt.Printf("\n 30 s1[:len(s1)] len %d cap  %d  s1 %v\n", len(s1), cap(s1), s1[:len(s1)])
	fmt.Println("========= another reslice example ===============")
	// cannot reference slice element beyond len()
	// .      fmt.Printf(" 40 s1[len(s)+2] = %d\n", s1[len(s1)+2]) // panic: runtime error: index out of range
	// however can reslice by specifying slice operator s[x:len(s1)+2]
	fmt.Printf("\n 41 s1[:len(s1)+2] len %d cap  %d  s1 %d \n", len(s1[:len(s1)+2]), cap(s1[:len(s1)+2]), s1[:len(s1)+2])
	// reslice in last statement is not assigned  to s1 so orignal slice still holds
	fmt.Printf("\n 42  s1[:] len %d cap  %d  s1 %v  s1 resliced %v \n", len(s1), cap(s1), s1[:], s1[:len(s1)+1])
	//
	// assign reslice  to original slice to preserve change to len
	//
	s1 = s1[:len(s1)+2]
	fmt.Printf("\n 50 s1=s1[:len(s1)+2]  len %d cap  %d  s1 %v\n", len(s1), cap(s1), s1[:len(s1)])
	//        s1[len(s1)]=55     // causes panic - index out of range, must extend then popluate
	fmt.Printf("\n 60 s1 %d %d %d - %v %v last element is always len-1 because slices starts at element 0, also  s1[:] === s1[0:len(s1)]\n", s1[0], s1[1], s1[len(s1)-1], s1[:], s1[0:len(s1)]) // s1 3 4 [3 4]
	//
	// cannot reslice beyond cap - so following generates slice bounds out of range.
	//  s1 = s1[:len(s1)+1]                        // extend len byt 1
	// use append as this will realloc a new underlying array with a larger size
	//
	s1 = append(s1, 80, 90)
	fmt.Printf("\n 70  s1[:] len %d cap  %d  s1 %v - so len increased 2 and cap by 5. Note append appends to len but may also increase cap.\n", len(s1), cap(s1), s1[:])
	// original slice1 unchanged (of course)
	fmt.Printf("\n 71  slice1[:] len %d cap  %d  slice1 %v\n", len(slice1), cap(slice1), slice1[:])
	fmt.Println("========================")
	fmt.Printf("\n 2:1 s1 len %d %v", len(s1), s1) // s1 len 3 [3 4 5]
	s1[len(s1)-1] = 55
	fmt.Printf("\n 2:2 s1 len %d cap %d %v", len(s1), cap(s1), s1) // s1 len 3 [3 4 55]
	//
	// append extentds len and may optionaly extend cap - in this case it only extends len.
	//
	s1 = append(s1, 66)
	fmt.Printf("\n 2:3 s1 len %d cap %d %v", len(s1), cap(s1), s1) //s1 len 4 [3 4 55 66]
	//
	// append now allocates new array - increasing len by 1 and cap
	//
	s1 = append(s1, 99, 101, 102, 103)
	fmt.Printf("\n 2:4 s1 len %d cap %d %v", len(s1), cap(s1), s1) //s1 len 4 [3 4 55 66]

	fmt.Printf("\n\n 3:1 s1:  %+v", s1)
	fmt.Printf("\n 3:2 s1:  %#v", s1)
	fmt.Printf("\n 3:3 s1:  %v", s1)
	//
	// reslice to smaller slice - cap remains the same but len changes.
	//
	s1 = s1[:3]
	fmt.Printf("\n 3:4 s1 = s1[:3]  len %d cap %d  s1:  %v", len(s1), cap(s1), s1)
	//
	// reslice to larger slice
	//
	s1 = s1[:15]
	fmt.Printf("\n 3:5 s1 = s1[:3]  len %d cap %d  s1:  %v", len(s1), cap(s1), s1)
	//
	fmt.Println("\n========================\n")
	fmt.Printf("\n slice1_[13] d = %d", s1[10])                                                                   // prints 4
	fmt.Printf("\n slice1_[13] b = %08b", s1[10])                                                                 // prints 4
	fmt.Printf("\n slice1_[13] q = %q  a single-quoted character literal safely escaped with Go syntax.", s1[10]) // prints 4
	fmt.Printf("\n slice1_[13] o = %o     octal 0..80..80..8", s1[10])                                            // prints 4
	fmt.Printf("\n slice1_[13] x = %x     hexidecimal  0..F0..F  ", s1[10])                                       // prints 4
	fmt.Printf("\n slice1_[13] X = %X     hexidecimal  0..F0..F", s1[10])                                         // prints 4
	fmt.Printf("\n slice1_[13] u = %u    -  u not a fmt operator", s1[10])
	fmt.Printf("\n slice1_[13] U = %U    -  Unicode format", s1[10]) // prints 4

	fmt.Println("\n======= empty slice =================\n")
	var s []int // s is a pointer to something, but its storage component is nil hence s==nil
	fmt.Printf("\nvar s []int  ")

	if s == nil {
		fmt.Printf("\n  s==nil    is true. So slice zero value is nil\n\n   Testing for nil in slice tests for underlying array. s is infact not-nil as s points to slice metadata &s=%p, %#v", &s, s)
	} else {
		fmt.Printf(" %p %#v", &s, s)
	}
	fmt.Printf("\n .   s metadata:  len %d  cap %d \n", len(s), cap(s))

	s = append(s, 1) //  adding first elment creates storage to s and consequently s != nil

	if s != nil {
		fmt.Printf("\n  Append to s. This makes s != nil true.  Underlying storage has been allocated to accomodate append %p %#v", &s, s) //  prints.. s is not nil
	} else {
		fmt.Printf("\n     s is nil...")
	}
	fmt.Printf("\n      s metadata:  len %d  cap %d \n", len(s), cap(s))
	//
	// slice literal is not nil
	//
	fmt.Println("\n=======  slice literal =================\n")
	//
	slice20 := []int{} // slice20 is !=nill, so this allocates storage, AS PER DOCO.
	fmt.Printf("\n\n slice20 := []int{}")

	if slice20 == nil {
		fmt.Printf("\n             .... slice20 IS NIL..n")
	} else {
		fmt.Printf(" \n  slice20 != nil     - where as a non-literal slice is nil\n    A slice literal allocates underlying storage even if its empty so its not nil  type: %T  #v: %#[1]v\n", slice20) //  slice20 is a []int  []int{}
	}
	fmt.Printf("slice 20 len %d   cap %d", len(slice20), cap(slice20))
	slice20 = append(slice20, 66)
	fmt.Printf("\n Append 66 to slice20.  type: %T #v:  %#[1]v\n", slice20) //  slice20 is a []int  []int{}
	fmt.Printf("slice 20 len %d   cap %d", len(slice20), cap(slice20))

	slice30 := []int(nil) // NOTE:  T(?) is a conversion function - take ? and convert to type Tn
	fmt.Printf("\n\n slice30 := []int(nil)  ")

	if slice30 == nil {
		fmt.Printf(" \n slice30 == nil     - as no underlying storage is allocated because it is a nil converson.  %T  %#[1]v", slice30)
	}
	fmt.Printf("\ntype:  %T   slice30 len %d   cap %d", slice30, len(slice30), cap(slice30))

	slice40 := make([]int, 3)
	fmt.Printf("\n\n slice40:=make([]int,3)   len = %d,  cap = %d   , %#v", len(slice40), cap(slice40), slice40)
	if slice40 == nil {
		fmt.Printf("\nslice40 IS NIL")
	} else {
		fmt.Print("\nSlice40 IS NOT NIL  - as underlying storage is allocated.")
	}

	slice41 := make([]int, 1, 3)
	fmt.Printf("\n\n slice41:=make([]int,1, 3)   len = %d,  cap = %d   , %#v", len(slice41), cap(slice41), slice41)
	if slice41 == nil {
		fmt.Printf("\nslice41 IS NIL")
	} else {
		fmt.Print("\nSlice41 IS NOT NIL  - as underlying storage is allocated.")
	}

	slice50 := make([]int, 3)[3:]
	fmt.Printf("\n\n slice50:= make([]int,3)[3:]   len = %d,  cap = %d   , %#v", len(slice50), cap(slice50), slice50)
	if slice50 == nil {
		fmt.Printf("\nslice50 IS NIL")
	} else {
		fmt.Print("\nSlice50 IS NOT NIL - as underlying storage is allocated.")
	}

	fmt.Println("\n===============================")
	fmt.Println("\n=======  array =================\n")
	array1 := [...]int{1, 2, 3, 4, 5, 6}
	//array1 = append(array1,99)                                     // fails: first argument to append must be slice; have [6]int - as array1 is an array which is fixed in size

	fmt.Printf("\n\n slice1 is a %T  %#[1]v         - not slice has no dimension size", slice1)
	fmt.Printf("\n\n array1 is a %T  %#[1]v         - not array always has a dimension size", array1)

	slice20 = array1[:] // does this create a new underlying array or merely point to array1 storage
	array1[0] = -99
	fmt.Printf("\n\n slice20 is a %T  %#[1]v", slice20) // this confirms slice points to array storage.  Lets append to slice
	fmt.Printf("\n\n array1 is a %T  %#[1]v", array1)

	slice20 = append(slice20, []int{11, 12, 13}...) // this creates new storage for slice20 and copies data from array storage into it.
	array1[0] = -77                                 // now change array value to confirm this is not reflected in slice as it is using different storage now..

	fmt.Printf("\n\n slice20 is a %T  %#[1]v", slice20) // this confirms slice is using its own storage.
	fmt.Printf("\n array1 is a %T  %#[1]v", array1)

	fmt.Println("\n======= empty array =================\n")
	var s2 [2]int // s is a pointer to something, but its storage component is nil hence s==nil
	fmt.Printf("\nvar s2 [2]int ")

	fmt.Printf("\n &s=%p, %#v  - Array zero value as in this example contains the [arrya] elemennt zero values", &s2, s2)

}
