package main



import (
       "unsafe"
 	"reflect"
       "fmt"
       )
var a = [3][4]int{  
   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
   {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
   {8, 9, 10, 11},   /*  initializers for row indexed by 2 */
}

var xx uint32 = 0x7FFFFFF
var yy 

var f float64 = 0.0 //34341e56 //678.42
var fp uintptr
var fpp *float64

var  n  int = 9
var  ns []int = []int{5,6,34,234,525,323,987}

func main () {
     fmt.Println(xx/1024/1024)
     //
     // note: unsafe.Pointer is a type not a function.  It permits you to convert any a pointer to any variable so it can be converted to a pointer to another type.
     //
     // A pointer value of any type can be converted to a unsafe.Pointer.
     // A unsafe.Pointer can be converted to a pointer value of any type. (*int) (*flaot64) ..
     // A uintptr can be converted to a unsafe.Pointer.
     // A unsafe.Pointer can be converted to a uintptr.
     //
     // uintptr types can be used in pointer arithmetic
     //
     // to convert to unintptr must use unsafe.Pointer as it is the only way to override the compiler 
     // to force interpretation of pointer as something other than original type.
     //
     // in the following we do:  &f -> uintptr -> (*uint64)      Note: spitting conversion across two statements is an dangerous operation. (see Go guide P357)
     // 							 Should use single as in p function value below
     //
     //fp = &f 			                 // Error: cannot use &f (type *float64) as type uintptr in assignmen - must convert to unsafe.Pointer, as in next line..
     fp = (uintptr)(unsafe.Pointer(&f) )         // convert *float64 to uintptr using unsafe.Pointer   (*float64)-> uintptr
						 // converting to uinptr enables us to do arithmetic on the pointer, as uintptr holds the pointer numeric value.
     fpp = &f


     fmt.Printf("\b%T %[1]b %[1]g",f)					   // cannot %b on a float64 - gets rubbish
     //
     //  these are the same..   they show that we can inspect the float64 as  bit stream, which we cannot ordinarily do. We must interrupt the float as int then we can.
     //				we cannot assing int=float64 but we can do it via a pointer to a float and use unsafe.Pointer conversion to permit conversion to another pointer type
     //                         the name of the package, unsafe, reminds us that we can do things outside of the normal Go compiler, hence potentially unsafe.
     //
     fmt.Printf("\n%T %064[1]b",*(* uint64)(unsafe.Pointer(&f)) )  // convert (*float64) to (*uint64) and dereference to interpret float64 as binary.
     fmt.Printf("\n%T %064[1]b",*(* uint64)(unsafe.Pointer(fp)) )  // convert uintptr to (* uint64) and dereference to interpret float64 as binary.
     //
     fmt.Printf("\n%T %064[1]x",*(* uint64)(unsafe.Pointer(fp)) )  // convert uintptr to (* uint64) and dereference to interpret float64 as binary.
     fmt.Printf("\n%T %[1]g",*(* uint64)(unsafe.Pointer(fp)) )  // convert uintptr to (* uint64) and dereference to interpret float64 as binary.
     //
     pi := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer())) // new(int) returns pointer to an int. Pointer() returns uintptr.
     fmt.Printf("\npi %d\n",*pi)
     //
     fmt.Printf("\n\n\n")
     //
     b:=a[:2:2]
     fmt.Println("a[:2:2]   ", b)
     c:=a[1:2]
     fmt.Println("len(b) ",len(b))
     fmt.Println("len(a) ",len(a[1][:2]))
     fmt.Println("len(a) ",len(a[1][:2]))
     fmt.Println("len(a[:2:3]) ",len(a[:2:3]))
     fmt.Println("len(a[:2:5]) ",len(a[:2:2]))
     fmt.Println(c)
     // 
     fmt.Printf("\n\n\n")
     //
     //func p (f float64) uint64 {                     // doesn't work
     //
     p:=func (f float64) uint64 {                     //  function literal - assigned to variable. p is a functiono	 value.
           return *(* uint64)(unsafe.Pointer(&f))     //   *T(<oldtype>) - dereference of a conversion to T of pointer to float64. Brackets required 
    						      //                   treat &f as pointer to unint64 and dereference 
          }
     //
     // more examples of showing the bit pattern of a float64, using unsafe.Pointer to reinterrupt as an int value
     //
     var f float64 = 45e-2
     fmt.Printf("%T\n",p)
     fmt.Printf("1 %T %064[1]b %[1]d \n",p(34341e56));    // # adds leading zeros - interprets float bit pattern as uint64 via unsafe.Pointer
     fmt.Printf("2 %T %064[1]b %[1]d \n",p(34341e5));    // # adds leading zeros
     fmt.Printf("3 %T %064[1]b %[1]d \n",p(1000e2));    // # adds leading zeros
     fmt.Printf("4 %T %064[1]b %[1]d \n",p(f));    // # adds leading zeros
     fmt.Printf("5 %T %064[1]b %[1]d \n",p(1.0));    // # adds leading zeros
     fmt.Printf("2 %16x \n",p(1.0));                   // hex with size 16
     fmt.Printf("3 %x \n",p(1.0));     		      // hex with default size (16)
     fmt.Printf("4 %#x \n",p(1.0));                    // hex with leading 0x
     fmt.Printf("5 %#x \n",p(1<<32-1));                // hex with leading 0x
     fmt.Printf("6 %#o \n",p(1<<64-1));                // octal with leading 0
     fmt.Printf("7 %d %#[1]x %#64[1]b\n",uint64(1<<64-1),p((1<<64)-1))
     fmt.Printf("8 %g %#64[1]b %#64[1]b\n",float64(1<<64-2),p((1<<64)-2))	// NB cannot show bit pattern of float64 must use unsafe.Pointer ... as above
     //
     //
     //
     fmt.Printf("\n\n\n")
     //
     v:=reflect.ValueOf(&n).Pointer()				// take *n -> reflect.Value -> .Pointer() -> uintptr
     fmt.Printf("%d", *(*int)(unsafe.Pointer(v)) )
     //
     v=reflect.ValueOf(ns).Pointer()				// uintptr - pointer to slice.  Only reference types allowed. slice,map,func,pointer,chan
     fmt.Printf("\n%d",   *(*int)(unsafe.Pointer(v+(unsafe.Sizeof(ns[0])*5))) )    // value of 5th index into slice.
}

