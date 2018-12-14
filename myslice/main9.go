package main

import (
	"fmt"
)
var slice1  []int = []int{1,2,3,4,5,6}

func main() {
	slice2:=slice1[:2];
	slice3:=slice1[2:4];
	slice4:=slice1[4:6];

	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice1),cap(slice1), slice1)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice2),cap(slice2), slice2)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice3),cap(slice3), slice3)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice4),cap(slice4), slice4)

	for i:=0;i<len(slice1); i++ {
		slice1[i]*=-1				// this will mutate data in all slices  as they share the same storage (underlying array)
	}

	fmt.Printf("\n\n len= %d,  cap= %d, contents: %v",len(slice1),cap(slice1), slice1)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice2),cap(slice2), slice2)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice3),cap(slice3), slice3)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice4),cap(slice4), slice4)

	slice3=append(slice3,7);			// this overwrites element 4 in slice1, as they have same storageo

	slice4=append(slice4,8);			// appending exceeds capacity of slice4 & allocates completely new underlying storage 
							//  and populates it withe contents of original slice.
							//  slice 4 is now disassociated from the other slices.
							//  changing slice4 will not change other slices and visa-versa.

	fmt.Printf("\n\n len= %d,  cap= %d, contents: %v",len(slice1),cap(slice1), slice1)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice2),cap(slice2), slice2)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice3),cap(slice3), slice3)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v  ** now has new storage ",len(slice4),cap(slice4), slice4)

	for i:=0;i<len(slice1); i++ {
		slice1[i]*=10				// this will mutate data in all slices apart from slice4 which has its own array.
	}


	fmt.Printf("\n\n len= %d,  cap= %d, contents: %v",len(slice1),cap(slice1), slice1)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice2),cap(slice2), slice2)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice3),cap(slice3), slice3)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice4),cap(slice4), slice4)

	slice4[0]=-1;

	fmt.Printf("\n\n len= %d,  cap= %d, contents: %v",len(slice1),cap(slice1), slice1)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice2),cap(slice2), slice2)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice3),cap(slice3), slice3)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice4),cap(slice4), slice4)

	slice3=append(slice3,100,101,102)		// appends to len() but will exceed capacity of underlying array & create new storage array.

	for i:=0;i<len(slice1); i++ {
		slice1[i]+=1				// this will mutate data in slice 1 & 2 only
	}

	fmt.Printf("\n\n len= %d,  cap= %d, contents: %v",len(slice1),cap(slice1), slice1)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice2),cap(slice2), slice2)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v ** now has new storage - no 9 in element 4 as slice 1 has",len(slice3),cap(slice3), slice3)
	fmt.Printf("\n len= %d,  cap= %d, contents: %v",len(slice4),cap(slice4), slice4)

}
