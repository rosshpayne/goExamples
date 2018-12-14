package main

import (
	"fmt"
)

// this give an example of using a pointer to slice as an argument to a func.  Any change the func makes to the len of the slice will be 
// reflected in the caller.
// this proves that len is an attribute of the slice.  Append will append to the end of len of the slice.
//   extend a slice simply by adding increasing the b in slice[a:b]

var add []int = []int{90,91,92,93}
var add_array [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}

func main() {

	var mys []int					// does allocate memory to slice struct, mys!=nil, cap=len=0, however underlying array DOES NOT EXIST 
							// mys -> [ [ptr to underlying array], int (len), int (cap) ]  NO  [underlying array]

	var mys2_ [] int


	xs:=add_array[4:8]
	ys:=xs[:2]
        fmt.Printf("\n xs:  %T %v", xs,xs)	
        fmt.Printf("\n ys:  %T %v", ys,ys)	
        ys =append(ys[:1],555);
        fmt.Printf("\n ys:  %T %v", ys,ys)	
	fmt.Printf("\n %T, %[1]d",len(ys))

	fmt.Printf("\nvar mys []int:  %T %p %d %d ",mys,&mys,len(mys),cap(mys) )
	fmt.Printf("\nadd mys []int:  %T %p %d %d %v",add,&add,len(add),cap(add),add[:] )
	//fmt.Printf("\nvar mys []int:  %T %p %d %d %d",mys,&mys,len(mys),cap(mys), mys[0] )  // panics - cannot access underlying array element 0

        fmt.Printf("\nmys=append(mys,44)")		
        mys=append(mys,44);				//  appends element to position len 
        mys=append(mys,64);				//  appends element to position len 
        mys=append(mys,74);				//  appends element to position len 
        mys=append(mys,84);				//  appends element to position len 
        mys=append(mys,85);				//  appends element to position len 
        mys=append(mys,86);				//  appends element to position len 
        mys=append(mys,87);				//  appends element to position len 
        mys=append(mys,88);				//  appends element to position len 
        mys=append(mys,89);				//  appends element to position len 
	cat:=mys			// cat has len & cap of mys and same underly array of data
	cat2:=mys[4:7]
	fmt.Println("\ncat2:=mys[4:7]")
	fmt.Printf("\nvar mys []int:  %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])

        copy(mys[:2],add)								// copies appends add to beginning of destination slice mys, upto len(destination)
        fmt.Printf("\ncopy(mys[:2],add)")						// copies appends add to beginning of destination slice mys, upto len(destination)
	fmt.Printf("\n*  cat   %T %p %d %d %v",cat,&cat,len(cat),cap(cat),cat[:])
	fmt.Printf("\n   mys   %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])
	fmt.Printf("\n   cat2  %T %p %d %d %v",cat2,&cat2,len(cat2),cap(cat2),cat2[:])
        copy(mys[4:],add)			// copies add to mys[0]  
        fmt.Printf("\ncopy(mys[4:],add)")			// copies add to mys[0]  
	fmt.Printf("\n** mys   %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])

	cat=nil						// clears pointer to underly array.  Array candidate for GC, but still referenced by other vars.

	fmt.Printf("\n\n ***  cat=nil   %T %p %d %d %v",cat,&cat,len(cat),cap(cat),cat[:])
	fmt.Printf("\n     ** mys   %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])

	cat=append(cat,44)				// allocate new underlying array to cat (previously nil) 

	fmt.Printf("\n\n ***  cat=nil   %T %p %d %d %v",cat,&cat,len(cat),cap(cat),cat[:])
	fmt.Printf("\n     ** mys   %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])


	// can we manually extend a slice - NO must use append function.
        //mm:=mys[:2]     //  panic: runtime error: slice bounds out of range
        //mm[1]=-44

	mys_:=mys[:]					//  allocates memory to mys_ (the 3 attributes of slice) and points to same underlying array as mys
	fmt.Printf("\nnew slice assigned from existing one: mys_         %T %p %d %d %v",mys_,&mys_, len(mys_),cap(mys_),mys_[:])

	fmt.Printf("\n\nmys=make([]int,5,10) ");
	mys=make([]int,5,10)				// only allocates new underlying array - same mys struct . Old underlying array is GC.	
							//   to preserve old underlying array must copy it into new slice usig copy(mys,mys_)
	fmt.Printf("\n\nnew(mys)            %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])		// mys still points to same struct by has new underlying array
	copy(mys,mys_)					// copies underlying arrays only
	fmt.Printf("\nnew(mys) with copy   %T %p %d %d %v",mys,&mys,len(mys),cap(mys),mys[:])		// mys still points to same struct by has new underlying array
	fmt.Printf("\nold(mys_)         %T %p %d %d %v",mys_,&mys_, len(mys_),cap(mys_),mys_[:])

	mys2:=make([]int,5,10)				
        mys2[0]=11
        mys2[1]=21
        mys2[2]=31
        mys2[3]=41
        mys2[4]=51

	fmt.Printf("\n\nnmys2 %T %p %d %d",mys2,&mys2,len(mys2),cap(mys2))

	mys2_=mys2[:len(mys2)+2]             //**** can create new slice which is an extension to slice mys2 from len() 5 -> 8 provided its less than cap()
				   // mys2_ is a different slice struct but points to the same underlying array as mys2

	fmt.Printf("\n**mys2  %T %p %d %d %#v",mys2,&mys2,len(mys2),cap(mys2),mys2)
	fmt.Printf("\n**mys2_ %T %p %d %d %#v",mys2_,&mys2_,len(mys2_),cap(mys2_),mys2_)
        // mys2[6]=33               // but cannot manually extend actual mys2 by assigning value beyond len() // panic: runtime error: index out of range

        mys2=append(mys2,61)
	fmt.Printf("\nnmys2 %T %p %d %d %#v",mys2,&mys2,len(mys2),cap(mys2),mys2)
	fmt.Printf("\nmys2_ %T %p %d %d %#v",mys2_,&mys2_,len(mys2_),cap(mys2_),mys2_)


	mys3:=[]int{}					// declares & allocates slice struct   mys!=nil,  cap=len=0, underlying array DOES NOT EXIST
							// mys -> [ [ptr to underlying array], int (len), int (cap) ]  NO  [underlying array]

	//fmt.Printf("\n\nmys3  %T %d %d %d %d",mys3,&mys3,len(mys3),cap(mys3),mys3[0])   // cannot access underlying array element 0
	fmt.Printf("\nmys3  %T %d %d %d ",mys3,&mys3,len(mys3),cap(mys3))   

	mys4:=[]int{4,45,56,77,2,4,88,97}		// declares & allocates underlying array of 8  elements  mys!=nil,  cap=len=0, underlying array populated

	fmt.Printf("\n%T %d %d %d %d",mys4,&mys4,len(mys4),cap(mys4),mys4[0])

}
