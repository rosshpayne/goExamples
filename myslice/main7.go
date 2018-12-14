package main

import (
	"fmt"
)

// this give an example of using a pointer to slice as an argument to a func.  Any change the func makes to the len of the slice will be 
// reflected in the caller.
// this proves that len is an attribute of the slice.  Append will append to the end of len of the slice.
//   extend a slice simply by adding increasing the b in slice[a:b]

var anotherbit = []int{-1,-2,-3,-4,-5}

func main() {
        //
	var aa []int 
        var parseState []int
	fmt.Printf("parseState = %d %d %#v\n",len(parseState),cap(parseState),parseState)
        parseState = []int{1,2,3,4,5}
	fmt.Printf("parseState = %d %d %#v\n",len(parseState),cap(parseState),parseState)
        xparse := parseState[0:0]
	fmt.Printf("xparse = %d %d %#v\n",len(xparse),cap(xparse),xparse)

        aa = make([]int,0,100)
	fmt.Printf("aa = %d %d %#v\n",len(aa),cap(aa),aa)
	for i:=0;i<412;i++ {
	     aa=append(aa,i)
	}	
	fmt.Printf("aa2 = %d %d %#v\n",len(aa),cap(aa),aa)

	for i,v := range aa[100:104] {
		fmt.Println(i,v)
	}

        bb:=aa[410:]								// bb is a portion of aa, note len
	fmt.Printf("\nbb2 before call = %d %d %#v\n",len(bb),cap(bb),bb)	
        extabit(&bb)								// pass pointer to bb to func
	fmt.Printf("bb2 after call = %d %d %#v\n\n",len(bb),cap(bb),bb)		// len of bb has changed, as we used pointer to bb and len is part of slice.
										//   cap has not changed as its an attribute of underlying array not slice.

        aa =aa[:len(aa)+2]							// extending slice to view more of underlying array
	fmt.Printf("aa2 = %d %d %#v\n",len(aa),cap(aa),aa)
        aa = aa[:len(aa) -5] 							// reduce len
	fmt.Printf("aa2 = %d %d %#v\n",len(aa),cap(aa),aa)
	cc:=aa[len(aa):cap(aa)]							// force def beyond len(aa) [:] means o:len(aa)
	fmt.Printf("cc  = %d %d %#v\n",len(cc),cap(cc),cc)
	dd:=aa[:]
	fmt.Printf("dd  = %d %d %#v\n",len(dd),cap(dd),dd)
	
}

func extabit ( aslice *[]int ) {

	fmt.Printf("aslice a= %d %d %#v\n",len(*aslice),cap(*aslice),*aslice)
        fmt.Println((*aslice)[1])
//        fmt.Println(aslice[1])					// invalid operation: aslice[1] (type *[]int does not support indexing)

	*aslice=append(*aslice,anotherbit...)

	fmt.Printf("aslice a= %d %d %#v\n",len(*aslice),cap(*aslice),*aslice)
}

