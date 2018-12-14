package main

import (
	"fmt"
)

//var aa [9]int = [...]int{11, 12, 13, 14, 15, 16, 18, 19, 20}    // array not slice
var aa []int = []int{11, 12, 13, 14, 15, 16, 18, 19, 20}    //  slice

type pgids []int

type xx struct { thd []int}

func main() {
        //
        //part:=aa[:len(aa)]						//part and aa share underlying array
        var vx xx
        vx.thd=make([]int,len(aa))
        copy(vx.thd,aa)
        fmt.Println(vx.thd)						// [11 12 13 14 15 16 18 19 20]
        vx.thd[0] =-1
        fmt.Println(vx.thd)
        fmt.Println(aa)	
	dst:=make(pgids,12);
        fmt.Printf("dst %d %d\n",len(dst),cap(dst))

        dst =aa[:]
        fmt.Printf("aa[:] %d %d %v\n",len(dst),cap(dst),dst)

        dst[4]=22
        fmt.Printf("dst[4] %d %d %v\n",len(dst),cap(dst),dst)

        dst2:=aa[3]							// take a single element of a slice
        fmt.Printf("aa[3] %d \n",dst2)

        merged:=dst[:0]							// merged & dst share underlying array
        fmt.Printf("dst[:0]  %d %d\n",len(merged),cap(merged))

        merged=append(merged,33);
        fmt.Printf("append 33 %d %d %v\n",len(merged),cap(merged),merged)
        fmt.Printf("dst %d %d %v\n",len(dst),cap(dst),dst)

        merged=append(merged,44);
        fmt.Printf("append 44  %d %d %v\n",len(merged),cap(merged),merged)
        fmt.Printf("dst %d %d %v\n",len(dst),cap(dst),dst)

        bb := make([]int, len(aa), len(aa) )
        fmt.Printf("bb %d %d %v\n",len(bb),cap(bb),bb)
        
        copy(bb,aa)       
        fmt.Printf("bb %d %d %v\n",len(bb),cap(bb),bb)
        for i, r := range aa {
		bb[i]=r
	}
        fmt.Println("bb   ",bb)
        
		         
}
