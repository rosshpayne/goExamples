package main


import (
	ft2 "fmt"
)
//===============================================================================
//  long
//      var v T = expression
//      var v T
//      var v = expression
//  short
//      v:=expression
//      v1,v2:=expression           - where expression is a func that returns two types.
//
//	where expression is a function that returns a type or a type literal
//================================================================================
//
// note first letter of identifier determines whether the identifer is visible outside of the package. This applies to all data types.
//
type adhoc struct {
	a int
	b uint16
	c bool
	X *nested
}

type nested struct {
	A  map[string][]string
	B  [5]int
	}

//  NOTE: IN GO ALL DECLARATION HAVE AN ASSIGNED (ZERO) VALUE, sometimes this maybe nill, other times memory is allocatino and initialised *****

var zoo adhoc = adhoc{a: 1, b: 2323,c: 1==3 }		// package variable - allocated from heap - avoid.  Must use long form ie. var. Short form not allowed here.
var nest adhoc = adhoc{a: 1, b: 232,c: 1==3 }

//  NB "make" only applies to []slice, map, chan 

func main() {

	//
	var func1 = func () adhoc { return adhoc{a: 66, c: true} }
	//
	//  demonstrating SIX TYPES OF VAR DECLARATION (5 static & 1 dynamic)
	//
	//  use var or type literal to declare a variable
        //      a variable is a piece of storage containing a value
        //
	//  in the case of struct and slice var will allocated both the struct (with each field populated with their zero value) and the metadata for a slice with no underlying array
	//
	var dilbert adhoc					// var v T       - "var" long form. Must be used in package declar option in func declars.

        dilbert3:=adhoc{}					// v:=expression - here expression is struct literal, T{}, specifying zero type values to be used.

	var dilbert5 adhoc = adhoc{a:1,b:3,c:true}              // var v T = expression  - in this example expression is struct literal but could be a func
								//                use long form if you want to override deferred type from expression and use expression merely 
								//                to intialise data.

	dilbert4:=adhoc{a:1,b:3,c:true}				// v:=expression - here expression is struct literal, T{},  specifying non-zero values
								//  note:  var x adhoc
								//         var x adhoc = adhoc{}
								//         x:=adhoc{}     	- are ALL equivalent, all allocating memory and populating with field zero values.

	dilbert6:=dilbert4					// v:=expression - inferred - short var declar using inferred type. Initial values from expression.

	dilbert7:=new(adhoc)                                    // new(T)        - dynamic

	dilbert8:=func1()					// v:=expression - where expression is a function returning a type

	var dilbert9 = func1()					// var v = expression - where expression is a func

	// here we haven;t shown compound declartion where an expression return multiple variables
	//  using multi return function must use short declar form.  
	//             v1,v2 := expression
	//  either v1 or v2 can be pre-declar provided it was in the local lexical scope. If not use are forced to use long declar form for the other variable
	//   and change the := to =

	ft2.Printf("\n\n ******* Six declartions of a struct  *********\n\n")
        ft2.Printf("\nType of dilbert:  %T  %[1]#v\n",dilbert)
        ft2.Printf("\nPointer used to access dilbert.. [%p]\n",&dilbert)
	ft2.Printf("\ndilbert.c = %v\n",dilbert.c)
	ft2.Printf("\ndilbert3.c = %v\n",dilbert3.c)
	ft2.Printf("\ndilert6  T:=T1   %#v ",dilbert6)
	ft2.Printf("\ndilert7  new(T)  %#v ",dilbert7)
	ft2.Printf("\ndilert8  new(T)  %#v ",dilbert8)
	ft2.Printf("\ndilert9  new(T)  %#v ",dilbert9)
	ft2.Printf("\n\n ** Assign 453 to field b \n")
        dilbert.b=453
        dilbert3.b=453
        dilbert4.b=453
        dilbert5.b=453
        dilbert6.b=453
        dilbert7.b=453
        ft2.Printf("\nType of dilbert:  %T  %[1]#v\n",dilbert)
        ft2.Printf("\nPointer used to access dilbert.. [%p]\n",&dilbert)
	ft2.Printf("\ndilbert.b = %d\n",dilbert.b)
	ft2.Printf("\ndilbert3.b = %d\n",dilbert3.b)
	ft2.Printf("\ndilert6  T:=T1   %#v ",dilbert6)
	ft2.Printf("\ndilert7  new(T)  %#v ",dilbert7)
        //
	//  slice declaration & initialisation at compile time (statically)
        //
	ft2.Printf("\n\n ******* Seven declartions of a slice  *********\n\n")
	var a1 []int				   // allocates memory for slice metadata [len int, cap int,  p *data]
	var a1a []int = []int{1,2,3,4}             // allocates memory for slice metadata and underlying data array and populates array with values.
        a2 := []int{}                              // len=cap=0    var a2 []int = []int{}, ie. underlying array is populated but has no values.
        a3 := []int{1: 55, 2: 655, 3: 435}
        a4 := make([]int,5,10)			   // slice contains int zero value 0, upto cap, but only len are accessible otherwise get out-of-bound erros.
        a5 := new([]int)
        a6 := a3[2:len(a3)]
        //
	if a1 == nil {
		ft2.Println("a1 is nil")
	}
	if a1a == nil {
		ft2.Println("a1a is nil")
	}
	if a2 == nil {
		ft2.Println("a2 is nil")
	}
        ft2.Printf("\nlen,cap for a1  []int               %#v  %d, %d  " ,a1, len(a1), cap(a1))
        ft2.Printf("\nlen,cap for a1a []int = []int{1,2,3,4} %#v , %d, %d " , a1a,len(a1a), cap(a1a))
        ft2.Printf("\nlen,cap for a2:=[]int{}             %#v,  %d, %d  " , a2,len(a2), cap(a2))
        ft2.Printf("\nlen,cap for a3 := []int{1: 55, 2: 655, 3: 435} %d, %d " , len(a3), cap(a3))
        ft2.Printf("\nlen,cap for a4 := make([]int,5,10)  %#v,  %d, %d  " , a4, len(a4), cap(a4))
        ft2.Printf("\nlen,cap for a5 := new([]int)        %#v,  %d, %d " , a5,len(*a5), cap(*a5))
        ft2.Printf("\nlen,cap for a6 := := a3[3:5])        %#v,  %d, %d " , a6, len(a6), cap(a6))
        ft2.Printf("\n")
        //
        //  append - slice allocate storage dynamically
	//           must use when static declaration has allocated zero space (see above) or when underlying is full & must be extended.
        //

        







	//candies := []string{"abc","def","klm"}                  // use slice literal to assign initial values.  NB always curlly brackets {} for literals
								//   alternative allocate indexes also   []string{2:"abc", 3:"def", 4:"klm"}
	candies := []string{2:"abc", 3:"def", 4:"klm"}		// short var declaration with struct literal assigning non-zero-type-values
	ft2.Printf("\n\n%s",candies[2])

	// note use of struct literals to populate fields in dilbert2  e.g.

				// :=  allocates memory for adhoc struct (and all its fields), then the struct literals populate the fields

				// map[string][]string{"ABC": []string{"CAT","DOG"}, }
				// [5]int{4,5,6,7,8}

        //dilbert2 := adhoc{1,2,true,&nested{A: make(map[string][]string), B: [5]int{4,5,6,7,8}}}

/**************************************************************************************************** */
        //var dilbert2  adhoc
        // dilbert2 = adhoc{1,2,true,&nested{A: map[string][]string{"ABC": []string{"CAT","DOG"}, } , B: [5]int{4,5,6,7,8}}}  // assignment via struct literal T{}
        // or
        dilbert2 := adhoc{1,2,true,&nested{A: map[string][]string{"ABC": []string{"CAT","DOG"}, } , B: [5]int{4,5,6,7,8}}}    // short var declaration using struct literal T{}
        dilbert2b := adhoc{1,2,true,&nested{A: map[string][]string{"ABC": []string{"CAT","DOG"}, } , B: [5]int{4,5,6,7,8}}}    // short var declaration using struct literal T{}

        ft2.Printf("dilbert2   %T  %[1]#v\n\n",dilbert2)
        ft2.Printf("dilbert2   %T  %[1]v\n\n",dilbert2)
        ft2.Printf("dilbert2b   %T  %[1]v\n\n",dilbert2b)

	if dilbert2 == dilbert2b {
		ft2.Println("******  diltert2 == diltbert2b     even if nested pointers are different but contain some results")
	} else {
		// this is the case here..
		// considered shallow comparison as contents of nested are the same but pointers to it are different.
		ft2.Println("******  diltert2 != diltbert2b     even if nested pointers are different but contain some results\n\n")   // this is the case here..
																   // consider shallow comparison 
	}

	if dilbert2.X == dilbert2b.X {
		ft2.Println("******  diltert2.X == diltbert2b.X")
	} else {
		ft2.Println("******  diltert2.X != diltbert2b.X \n\n")   // this is the case here..
	}
/*
	if *(dilbert2.X) == *(dilbert2b.X) {  // note: cannot compare structs as it contains a map and maps are not comparable
		ft2.Println("******  diltert2.X == diltbert2b.X")
	} else {
		ft2.Println("******  diltert2.X != diltbert2b.X \n\n")   // this is the case here..
	}
*/


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
	//(*XX).A=make(map[string][]string,2)			// same as XX.A=
	XX.A=make(map[string][]string,2)			// same as XX.A=
        (*XX).A["cat"]=[]string{"cute","little"}	        // 
        (*XX).A["dog"]=[]string{"large","load"}	
	(*XX).B= [5]int{4,6,23,5,3}

	ft2.Printf("\nXX = %#v\n",XX)
	//
	zoo.X=XX				// shallow copy   -  struct = struct
	ft2.Printf("\nzoo: %v",zoo)
	newZoo := new(adhoc)
	*newZoo = zoo				// copies adhoc & nested structs so its deep copy
	ft2.Printf("\nnewZoo: %v",newZoo)
        var def []int = []int{1,2,3,4,5,6,7,8,9}
        ft2.Printf("\n\ndef   %#v %d",def,len(def))
        ext(def)
        ft2.Printf("\ndef   %#v %d",def,len(def))
        extp(&def)
        ft2.Printf("\ndef   %#v %d",def,len(def))

	//
	//  byte.Buffer declarations
	//
	//var buf = new(bytes.Buffer) 				// declaration. Allocation zero initial size.   pointer to bytes.Buffer.  
	//var buf = bytes.NewBuffer([]byte)			// *Buffer
        //var buf = bytes.NewBufferString(string)		// *Buffer
	//var buf = bytes.Buffer				// Buffer. Go will apply & automatically so buf.Len() becomes &buf.Len()		
        
}


func ext (abc []int)  {
    abc = abc[2:]
    ft2.Printf("\nabc %#v, %d",abc,len(abc))
}
func extp (abc *[]int)  {
    (*abc) =( *abc)[2:]
    ft2.Printf("\nabc %#v, %d",*abc,len(*abc))
}
