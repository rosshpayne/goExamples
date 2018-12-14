package main

import (
	"fmt"
)

var a1 [3]int = [...]int{1, 2, 3}
var b1 [9]int = [...]int{11, 12, 13, 14, 15, 16, 18, 19, 20}

type st struct {
	f1 int
	f2 *[3]int
	f3 *[9]int
}

var s1 []st = []st{st{f1: 1, f2: &a1, f3: &b1},      // struct literal declaration
	st{f1: 2, f3: &b1, f2: &a1},
	st{f1: 3, f3: &b1, f2: &a1},
	st{f1: 4, f3: &b1, f2: &a1},
	st{f1: 5, f2: &a1, f3: &b1},
	st{f1: 6, f2: &a1, f3: &b1},
	st{f1: 7, f2: &a1, f3: &b1},
	st{f1: 8, f3: &b1, f2: &a1},
	st{f1: 9, f2: &a1, f3: &b1},
	st{f1: 10, f3: &b1, f2: &a1},
	st{f1: 11, f2: &a1, f3: &b1},
}
var s3 []st = []st{st{f1: 1, f2: &a1, f3: &b1},
	st{f1: 2, f3: &b1, f2: &a1},
	st{f1: 3, f3: &b1, f2: &a1},
	st{f1: 4, f3: &b1, f2: &a1},
	st{f1: 5, f2: &a1, f3: &b1},
	st{f1: 6, f2: &a1, f3: &b1},
	st{f1: 7, f2: &a1, f3: &b1},
	st{f1: 8, f3: &b1, f2: &a1},
	st{f1: 9, f2: &a1, f3: &b1},
	st{f1: 10, f3: &b1, f2: &a1},
	st{f1: 11, f2: &a1, f3: &b1},
}


func main() {
        //
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i].f1, *(s1[i].f2), *(s1[i].f3))
	}
        s23:=s1[:0]      // len 0 cap >0
        fmt.Printf("s23: %d %d  %#v\n",len(s23),cap(s23),s23) 
        //
        // cut i & j  - this approach creates new underlying array and makes s1 avaialable for gc.
        //
        i,j:=4,6			// index start at 0 so 4 means index 5 when starting at 1 ie. 5,6
        s2:=make([]st,len(s1)-(j-i))        
        s21:=make([]st,len(s1)-(j-i))        
        s22:=make([]st,len(s1)-(j-i))        
        copy(s2[:i],s1[:i])			// s2[0:4],s1[:4] refers to elements 0,1,2,3 
        copy(s2[i:],s1[j:])			// s2[4:] ,s1[5:]
        fmt.Println("cut 5 & 6 ")
	for i := 0; i < len(s2); i++ {
		fmt.Println(i,s2[i].f1, *(s2[i].f2), *(s2[i].f3))
	}
        s21=s2[:]
        s22=s2
        fmt.Println(s21)
        fmt.Println(s22)
        fmt.Printf("s22 len, cap %d  %d\n",len(s22),cap(s22))
	for i := 0; i < len(s22); i++ {
		fmt.Println(i,s22[i].f1, *(s22[i].f2), *(s22[i].f3))
	}
        //
        // cut between i & j setting nil values
        //
        copy(s3[i:],s3[j:]);
        for i:=len(s3)-(j-i); i< len(s3); i++ {
               s3[i].f2, s3[i].f3  = nil, nil   // set pointer types to nil so referenced object can be gc
        } 
        s3=s3[:len(s3)-(j-i)]                   // shorten slice 
        //
        fmt.Println("s3 ",len(s3),cap(s3))
	for i := 0; i < len(s3) ; i++ {
		fmt.Println(s3[i].f1, *(s3[i].f2), *(s3[i].f3))
	}
        //
        // extend - add to tail
        //
        s3=append(s3,make([]st,5)...) 
        fmt.Println("expand s3 ",len(s3),cap(s3))
        //
        // expand - add to midpart of slice
        //
        i,x:=2,3
        //s4:=[]st{}      // len=0, s4!=nil
        //s4=append(append(append(s4[:],s3[:i]...),make([]st,x)...),s3[i:]...)
        fmt.Println("len(s3) ",len(s3))
        s3=append(s3[:i],append(make([]st,x),s3[i:]...)...)    // make slice will be gc.  No new underlying array. Expand insitu.
        fmt.Println("len(s3) ",len(s3))
	for i := 0; i < len(s3) ; i++ {
             if s3[i].f2!=nil {
		fmt.Println(s3[i].f1, *(s3[i].f2), *(s3[i].f3))
             } else {
		fmt.Println(s3[i].f1) 
             }
	}
        s4:=s3[5:8]      // copy cap over to s4, calc len =8-5=3, *data &s3[5] 
       fmt.Println("s4 len cap ",len(s4), cap(s4),s4[0]," ",s4[1]," ",s4[2])
        s4[0].f1=99
	for i := 0; i < len(s3) ; i++ {
             if s3[i].f2!=nil {
		fmt.Println(s3[i].f1, *(s3[i].f2), *(s3[i].f3))
             } else {
		fmt.Println(s3[i].f1) 
             }
	}
        s5:=s3[0:0] 
       fmt.Println("s5 len cap ",len(s5), cap(s5))

}
