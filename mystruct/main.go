package main


import (
	ft2 "fmt"
)

type adhoc struct {
	a int
	b uint64
	c bool
	X nested
}

type nested struct {
	A  map[string][]string
	B  [5]int
	}

//  NOTE: IN GO ALL DECLARATION HAVE AN ASSIGNED (ZERO) VALUE ****************************************************************************

var zoo adhoc = adhoc{a: 1, b: 23234234,c: 1==3 }		// package variable - allocated from heap - avoid 
var nest adhoc = adhoc{a: 1, b: 23234234,c: 1==3 }

//  NB "make" only applies to []slice, map, chan 

func main() {
	//
	//  use var or type literal to declare a variable
        //      a variable is a piece of storage containing a value
        //
	var dilbert adhoc					// "var" allocates memory to dilbert and populates fields with zero values of the field type.

/*	if dilbert == nil {					// error cannot compare struct to nil. /main.go:34: cannot convert nil to type adhoc
		ft2.Println("dilbert is nil")
	}*/
        dilbert3:=adhoc{}					// short var declaration with struct literal zero values

	dilbert4:=adhoc{a:1,b:3,c:true}				// short var declaration with struct literal non-zero values

	var dilbert5 adhoc = adhoc{a:1,b:3,c:true}              // long var declaration with struct literal non-zero values

        ft2.Printf("\nType of dilbert:  %T  %[1]+v\n",dilbert)
        ft2.Printf("\nPointer used to access dilbert.. [%p]\n",&dilbert)
	ft2.Printf("\ndilbert.c = %v\n",dilbert.c)
	ft2.Printf("\ndilbert3.c = %v\n",dilbert.c)
        dilbert3.b=453
        dilbert4.b=453
        dilbert5.b=453

	//candies := []string{"abc","def","klm"}                  // use slice literal to assign initial values.  NB always curlly brackets {} for literals
								//   alternative allocate indexes also   []string{2:"abc", 3:"def", 4:"klm"}
	candies := []string{2:"abc", 3:"def", 4:"klm"}		// short var declaration with struct literal assigning non-zero-type-values
	ft2.Printf("\n%s",candies[2])

	// note use of struct literals to populate fields in dilbert2  e.g.

				// :=  allocates memory for adhoc struct (and all its fields), then the struct literals populate the fields

				// map[string][]string{"ABC": []string{"CAT","DOG"}, }
				// [5]int{4,5,6,7,8}

        //dilbert2 := adhoc{1,2,true,nested{A: make(map[string][]string), B: [5]int{4,5,6,7,8}}}

        //var dilbert2  adhoc
        // dilbert2 = adhoc{1,2,true,nested{A: map[string][]string{"ABC": []string{"CAT","DOG"}, } , B: [5]int{4,5,6,7,8}}}  // assignment via struct literal T{}
        // or
        dilbert2 := adhoc{1,2,true,nested{A: map[string][]string{"ABC": []string{"CAT","DOG"}, } , B: [5]int{4,5,6,7,8}}}    // short var declaration using struct literal T{}

        ft2.Printf("dilbert2   %T  %[1]#v\n\n",dilbert2)
        ft2.Printf("dilbert2   %T  %[1]v\n\n",dilbert2)

	dilbert8 := new(adhoc)
	*dilbert8= dilbert2						// shallow copy of dilbert2 - copies nested field to
        ft2.Printf("dilbert8   %T  %[1]#v\n\n",dilbert8)
        ft2.Printf("dilbert8   %T  %[1]v\n\n",dilbert8)

        return

	ft2.Printf(" 0x1000 %b %[1]d \n", 0x1000)
        ft2.Printf(" 0x7FFFFFFF %b %[1]d \n", 0x7FFFFFFF)
        ft2.Printf(" (1 << 31) %b %[1]d \n", (1 << 31) )
        var xx uint64 = 0xFFFFFFFFFFFFFFFF
        ft2.Printf(" 0xFFFFFFFFFFFFFFFF uint64 %b %[1]d \n", xx)
        //var xx2 = 0xFFFFFFFFFFFFFFFF
        //ft2.Printf(" 0xFFFFFFFFFFFFFFFF int64 %b %[1]d \n", xx2)
        ft2.Println(1235%2)
	//



	cat := adhoc{}						// short var declaration using struct literal T{} - allocates memory for cat,fields are initialised wiht zero values
	ft2.Println("cat.c -> ",cat.c)                          // prove cat is initialised to zero type values
	cat = zoo 						// var=var ie. struct var = struct var - shallow  copy taen
	cat22:=zoo						// short var declaration using ordinary assignment. Create new memory struct and allocates from values from zoo.

	ft2.Println("zoo -                        ",zoo)
	ft2.Println("*cat - should be copy of zoo  ",cat )
	ft2.Println("*cat22 - should be copy of zoo  ",cat22 )

	cat22.c=true					 	// to prove cat22 is separate allocation - change a field value and compare to struct that initialised it.
	ft2.Println("\n\nzoo -                        ",zoo)
	ft2.Println("*cat22 - should have different field c value ",cat22 )
        ft2.Printf("\n******************************\n")




	//
	ft2.Println(" copy using pointers ");
        zoo2 := &zoo						// short var declaration. create pointer to zoo 
        //cat3 := &adhoc{}
	cat3 := new(adhoc)					// short var declaration. Allocate struct and assign zero-type-values to felds. &adhoc{} & new(adhoc) are the same
	
        ft2.Printf("\n cat3 field values should all be zero values: %v\n",cat3)   // NB go dereferenes cat3 automatically ie takes *cat3
        *cat3=*zoo2						//
	ft2.Printf("\n*cat3- should be copy of zoo2,zoo \n%v\n",*cat3)
        cat2 := &adhoc{}
        *cat2 = cat 			// *struct var = struct var
        cat2.c=1==1
	ft2.Println("*cat2: ",*cat2)
	//
        ft2.Printf("\n******************************\n")
	XX := new(nested)
	//XX := nested{}
	(*XX).A=make(map[string][]string,2)
        (*XX).A["cat"]=[]string{"cute","little"}	
        (*XX).A["dog"]=[]string{"large","load"}	
	(*XX).B= [5]int{4,6,23,5,3}
	//
	zoo.X=*XX				// shallow copy
	ft2.Printf("\nzoo: %v",zoo)
	newZoo := new(adhoc)
	*newZoo = zoo				// copies adhoc & nested structs so its deep copy
	ft2.Printf("\nnewZoo: %v",newZoo)
}
