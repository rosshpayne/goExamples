package main

import (
       "fmt"
       )

var s1=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
var s0=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}     //s0 & s1 must have different underlying arrays hence allocate one to s0
                                                                     // if s0 & s1 share underlying array i.e. s0:=s1 any changes made by s1 slice to underlying array
 								     // will be reflected in s0 as its the same array.

func main () {
       fmt.Println("s1[:3]  ",s1[:3])      //1,2,3
       // cut  element values 5,6 ie. id 6,7
       s1:=append(s1[:6],s1[8:]...)
       fmt.Println(" len, cap s1 = ",len(s1), cap(s1),s1)
       //
       // cut 2 elements starting at n
       //
       s1=initslice(s1,s0)
       n:=12      // remove 12,13
       s1=append(s1[:n],s1[n+2:]...)
       fmt.Println("at 12:  len, cap s1 = ",len(s1), cap(s1),s1)
       //
       // cut 5 elements starting at n
       //
       s1=initslice(s1,s0)
       n=15      // remove 12,13
       s1=append(s1[:n],s1[n+5:]...)
       fmt.Println("at 15:  len, cap s1 = ",len(s1), cap(s1),s1)
       //
       // delete single element
       //
       s1=initslice(s1,s0)
       n=15                  // delete index 15
       s1=append(s1[:n],s1[n+1:]...)
       fmt.Println("delete index 15:  len, cap s1 = ",len(s1), cap(s1),s1)
       //
       // delete single element using copy
       //
       s1=initslice(s1,s0)
       n=15                  // delete index 15
       s1=s1[:n+copy(s1[n:],s1[n+1:]) ]
       //s1=s1[:len(s1)-1]
       fmt.Println("delete index 15:  len, cap s1 = ",len(s1), cap(s1),s1)
}

func initslice(a, b []int) []int {
      // a & b slices initialised to calling arguments by value as all go arguments are.
      // however, slice is a reference type meaning a component of it contains a pointer in this a pointer to the underlying array
      // any changes to the underlying array by a & b slices will be be reflected in the calling program variables however the len & cap values
      // will not as they are not by reference only by value. Hence to get len & cap passed back we return the slice and the len&cap values will be copied
      // into the calling variable.
      // copy b into a
      //a=a[0:0]                // truncate a
      a=append(a[:0],b[:]...)   // truncate a and append b
      fmt.Println("*len, cap a = ",len(a), cap(a),a)
      fmt.Println("*len, cap b = ",len(b), cap(b),b)
      return a
} 
