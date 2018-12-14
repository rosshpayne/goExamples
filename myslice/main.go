package main    

import (
	"fmt"
)
var slice1  []int = []int{1,2,3,4,5,6}
var slice1_ []int = []int{10:21,22,23,12,25,26}			// specify index rather than default to 0.  Go will automatically increment from initial value.

func main() {
        fmt.Println("slice1: ",slice1)
	s1:=slice1[2:4]    // s1 is infered to be a slice of int of len 2

        fmt.Printf("\n slice1[2:4] len %d %v",len(s1),s1)
//        s1[len(s1)]=55     // causes panic - index out of range, must extend then popluate
        fmt.Printf("\n s1 %d %d %v  ",s1[0], s1[1],s1[:len(s1)]) i				// s1 3 4 [3 4] 

        s1 = s1[:len(s1)+1]     // extend len byt 1
        fmt.Printf("\n s1 len %d %v",len(s1),s1)						// s1 len 3 [3 4 5]
        s1[len(s1)-1]=55
        fmt.Printf("\n s1 len %d %v",len(s1),s1)						// s1 len 3 [3 4 55]
	s1=append(s1,66)
        fmt.Printf("\n s1 len %d %v",len(s1),s1)						//s1 len 4 [3 4 55 66]
	//s1[len(s1)]=77;     // extend by 1 in assignment - panic's index out of range, so [x] wher 0<x<len(s1)-1  



        fmt.Printf("\n\n\n\ns1:  %#v", s1) 
	s1=s1[:3]         // s1 is extended to assigned back to itself

        fmt.Printf("\ns1:  %#v", s1) 

	fmt.Printf("\n slice1_[13] d = %d",slice1_[13])			// prints 4
	fmt.Printf("\n slice1_[13] b = %08b",slice1_[13])			// prints 4
	fmt.Printf("\n slice1_[13] q = %q",slice1_[13])			// prints 4
	fmt.Printf("\n slice1_[13] o = %o",slice1_[13])			// prints 4
	fmt.Printf("\n slice1_[13] x = %x",slice1_[13])			// prints 4
	fmt.Printf("\n slice1_[13] X = %X",slice1_[13])			// prints 4
	fmt.Printf("\n slice1_[13] U = %U",slice1_[13])			// prints 4

	fmt.Printf("\n slice1_[15] = %d",slice1_[15])			// prints 6
	fmt.Printf("\n slice1_[3] = %d",slice1_[3])			// this should error as there is no index 3, start at 10,11,12,13,14,15

	var s []int;							// s is a pointer to something, but its storage component is nil hence s==nil
	fmt.Printf("\nvar s []int") 

	if s == nil {
		fmt.Printf("              ....  s is nil...")	       // s is nil
	} else {
		fmt.Printf(" %p %#v",&s,s)
	}
	fmt.Printf("\n %p %#v",&s,s)					//   but it also has an address. Why??  Maybe nil is checking on pointer component to storage

	s=append(s,1)							//  adding first elment creates storage to s and consequently s != nil 

	if s != nil {
		fmt.Printf("\n %p %#v",&s,s)				//  prints.. s is not nil
	} else {
		fmt.Printf("\n s is nil...")				
	}
	
	array1 := [...]int{1,2,3,4,5,6}					
	//array1 = append(array1,99)                                     // fails: first argument to append must be slice; have [6]int - as array1 is an array which is fixed in size

	fmt.Printf("\n\n slice1 is a %T  %#[1]v",slice1)
	fmt.Printf("\n\n array1 is a %T  %#[1]v",array1)

	slice20 := []int{}						// slice20 is !=nill, so this allocates storage, AS PER DOCO.
	fmt.Printf("\n\n slice20 := []int{}");

	if slice20 == nil {
	        fmt.Printf("             .... slice20 IS NIL..");
	} else {
		fmt.Printf("             .... slice20 is NOT NILL:   %T  %#[1]v",slice20)	//  slice20 is a []int  []int{}
	}
	slice20=append(slice20,66)
	fmt.Printf("\n slice20 is a %T  %#[1]v",slice20)	//  slice20 is a []int  []int{}

	slice30 := []int(nil)						// NOTE:  T(?) is a conversion function - take ? and convert to type Tn
	fmt.Printf("\n\n slice30 := []int(nil)");

	if slice30 == nil {
	        fmt.Printf("             ...slice30 IS NIL..");			// slice30 is nill AS PER DOCO
	} else {
		fmt.Printf("             ...slice30 IS NOT NIL  is a %T  %#[1]v",slice30)
	}

	slice40:=make([]int,3)
	fmt.Printf("\n\n slice40:=make([]int,3    len = %d,  cap = %d   , %#v", len(slice40), cap(slice40), slice40)
	if slice40 == nil { 
		fmt.Printf("\nslice40 IS NIL")
	} else {
		fmt.Print("\nSlice40 IS NOT NIL")
	}

	slice50:=make([]int,3)[3:]
	fmt.Printf("\n\n slice50:= make([]int,3)[3:]   len = %d,  cap = %d   , %#v", len(slice50), cap(slice50), slice50)
	if slice50 == nil { 
		fmt.Printf("\nslice50 IS NIL")
	} else {
		fmt.Print("\nSlice50 IS NOT NIL")
	}

	fmt.Println("\n===============================");

	slice20=array1[:];						// does this create a new underlying array or merely point to array1 storage
        array1[0]=-99
	fmt.Printf("\n\n slice20 is a %T  %#[1]v",slice20)		// this confirms slice points to array storage.  Lets append to slice
	fmt.Printf("\n\n array1 is a %T  %#[1]v",array1)

	slice20=append(slice20,[]int{11,12,13}...)			// this creates new storage for slice20 and copies data from array storage into it.
        array1[0]=-77							// now change array value to confirm this is not reflected in slice as it is using different storage now..

	fmt.Printf("\n\n slice20 is a %T  %#[1]v",slice20)		// this confirms slice is using its own storage.
	fmt.Printf("\n array1 is a %T  %#[1]v",array1)
}
