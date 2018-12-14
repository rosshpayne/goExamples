package main

import (
       "fmt"
       )

/*  1.  Slice assignment s=s2[a:b] share the same underlying array. Slice appending will create new arrays if cap exceeded
    2.  A slice is bound by len not cap. So an Append to slice s will insert at s[:len+1] not to s[:cap+1] 
    3.  If len cannot be increased without exceeding cap then cap will be increased by allocating new array 
*/

func main () {
   
      var s1 = make([]int,10)              // len = cap = 10 & elements assigned int zero value 0.
      var s2 = []int{6,7,8,9,6,7,8}        // len = cap = 7
      var s3 = []int{}                     // len = cap = 0 s3=nil
      //
      fmt.Println("s1: ",len(s1),cap(s1), s1)
      fmt.Println("s2: ",len(s2),cap(s2), s2)
      fmt.Println("s3: ",len(s3),cap(s3), s3)
      //
      s4:= s2[:0]
      fmt.Printf("\ns2[:0]  %T %[1]v %v %v\n ",s4,len(s4),cap(s4))
      s5:= s2[3]
      fmt.Printf("s2[3]  %T %[1]v \n",s5)
      s4 = s2[3:4]
      fmt.Printf("s2[3:4]  %T %[1]v %v %v\n ",s4,len(s4),cap(s4))
      s4 = s2[3:5]
      fmt.Printf("s2[3:5]  %T %[1]v %v %v\n ",s4,len(s4),cap(s4))
      //
      s1 = append(s1,[]int{1,2,3,4,5}...);              // spread operator ... enumerate slice into its elements as args to function.
      fmt.Println()
      fmt.Println("s1 len cap = ",len(s1),cap(s1), s1);

      s3 = append(s3,1);              // spread operator ... enumerate slice into its elements as args to function.
      fmt.Println("s3 len cap = ",len(s3),cap(s3), s3);

      s3 = append(s3,2,3,4);         
      fmt.Println("s3 len cap = ",len(s3),cap(s3), s3);
      s3 = append(s3,s1[10:]...);    // spread operator ... enumerate slice into its elements as args to function.
      fmt.Println("s3 len cap = ",len(s3),cap(s3), s3);

      s3=s3[:len(s3)+1]                     // manually append zero type value to s3 - len is now equal to cap 
      fmt.Println("s3 len cap = ",len(s3),cap(s3), s3);

      //s3=s3[:len(s3)+1]                     // manually append zero type value beyond cap. Will panic with slice bound out of range.
      //fmt.Println("s3 len cap = ",len(s3),cap(s3), s3);
      //
      // pop off top
      //
      a:=[]int{1,2,3,4,5,6,7,8,9,10}
      l:=len(a)
      var x int
      for i:=1;i<=l;i++ {
            x, a = a[len(a)-1], a[:len(a)-1]       // don't use x,a := a[len(a)    it doesn't assign value to a and len(a) doesn't decrement each time.
            fmt.Println(x,len(a),cap(a),a)
      }
      //
      // pop off  back
      //
      a=append(a,1,2,3,4,5,6,7,8,9,10)
      fmt.Println(x,len(a),cap(a),a)
      for i:=1;i<l;i++ {
          x,a = a[1], a[1:]
          fmt.Println(x,len(a),cap(a),a)
      }
      //
      // push onto top
      //
      a=[]int{}
      for i:=1;i<=10;i++ {
          a=append(a,i)
          fmt.Println(len(a),cap(a),a)
      }
      //
      // push onto back
      //
      a=make([]int,1)
      for i:=1;i<=10;i++ {
          a = append([]int{i},a...)      // create a new slice with one entry  and append current one to it. Nice.
          fmt.Println(len(a),cap(a),a)
      }

}
