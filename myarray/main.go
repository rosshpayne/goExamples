package main

import (
        "unsafe"
	"fmt"
	)
var ar [10]int = [...]int{1,2,3,4,5,6,7,8,9,10}
var arp *[4]int							// arp is a pointer to array of int
var pa uintptr
type xx struct {
       n  int
       pa uintptr             // pointer to ar
              }
type yy struct {
       pa uintptr
              }

func main () {
         fmt.Println("ar:  ",ar)
         arp := &ar
         fmt.Printf("ar        %T %[1]v\n",ar)
         fmt.Println()

         fmt.Printf("arp       %T %[1]v\n", arp)		// *[10]int &[1 2 3 4 5 6 7 8 9 10]     ptr to array of int
         fmt.Printf("arp[2]    %T %[1]v\n", arp[2]) 		// <pointer to array>[index]   go compier auto deref pointer
         fmt.Printf("(*arp)[2] %T %[1]d\n",(*arp)[2])		// <deref pointer to array>[index]  manually deref pointer arp so it becomes array ref.
         fmt.Println()

         fmt.Printf("ar[2]     %d\n",ar[2])			// <arrayTypeName>[index]
         arpslice:=arp[1:3]
         fmt.Printf("arsplice %v\n",arpslice)
         fmt.Printf("ar[5:8] %T  %[1]v\n",ar[5:8] )
         //
         // There two examples are he same as the last two - using unsafe for no real reason as arp it pointer to [4]int anyway.
         //
         pad := (* [4]int)(unsafe.Pointer(arp))[2]	  // use pointer to array to access index entry: *[4]int -> unsafe.Pointer -> *[4]int
         fmt.Printf("1 pad %d\n",pad)
         pad  = (*(* [4]int)(unsafe.Pointer(arp)))[2]     // or use deref pointer to array to access index entry. 
							  // Much the same as pointer to struct or struct to access struct field.
         fmt.Printf("2 pad %d\n",pad)
         //
         // here we save pointer to array to unintptr to be used later
         //
         pa   = (uintptr)(unsafe.Pointer(&ar)) 		 // *[4]int -> usafe.Pointer -> uintptr
         //:1
         s1:=xx{pa: (uintptr)(unsafe.Pointer(&ar)),n: 22,}
         s2:=yy{pa: (uintptr)(unsafe.Pointer(&ar))}
         //s2:=yy{pa: pa}
         fmt.Printf("1  %d\n",((* [4]int)(unsafe.Pointer(&ar)))[2])   // here were saying to interpret pa as pointer to [4]int
									 // need to use unsafe as we cannot deref uintptr by itself
									 // all pointer manipulation needs to go via unsafe - a compiler directive really.
         //
         fmt.Printf("1  %T\n",(* [4]int)(unsafe.Pointer(pa)))   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("2 %T\n",((*[4]int)(unsafe.Pointer(pa))))   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("3 %T\n",((* [4]int)(unsafe.Pointer(&s2.pa))))   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("4 %T\n",((* [4]int)(unsafe.Pointer(&s1.pa))))   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("5 %T\n",(* [4]int)(unsafe.Pointer(s1.pa)))   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("6 %T\n",(* [4]int)(unsafe.Pointer(s2.pa)))   // here were saying to interpret pa as pointer to [4]int
         //
         fmt.Printf("1  %d\n",(* [4]int)(unsafe.Pointer(pa))[2])   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("2 %d\n",((*[4]int)(unsafe.Pointer(pa)))[2])   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("3 %d\n",((* [4]int)(unsafe.Pointer(&s2.pa)))[2])   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("4 %d\n",((* [4]int)(unsafe.Pointer(&s1.pa)))[2])   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("5 %d\n",(* [4]int)(unsafe.Pointer(s1.pa))[2])   // here were saying to interpret pa as pointer to [4]int
         fmt.Printf("6 %d\n",(* [4]int)(unsafe.Pointer(s2.pa))[2])   // here were saying to interpret pa as pointer to [4]int
         //
         fmt.Printf("7 %d\n",(* [4]int)(unsafe.Pointer(&s2.pa))[2:3])   // here were saying to interpret pa as pointer to [4]int
         
}
