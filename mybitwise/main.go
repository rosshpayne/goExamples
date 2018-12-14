package main

import (
	"fmt"
)

func main() {
        var xx uint32
        xx=0x80 
        fmt.Printf("%b %d\n",xx,xx)
        fmt.Printf("%b %d\n",3<<8,3<<8)      // take 3 (bit pattern) and shift 8 bits to left
        yy:=xx>>2                            // take xx (bit pattern)  and shift 2 bits to right
        fmt.Printf("%b %d\n",yy,yy)
        xx='A' 
        fmt.Printf("%b %d\n",xx,xx)
        xx='z'
        fmt.Printf("z     %b %d\n",xx,xx)
        xx='z'>>1
        fmt.Printf("z>>1  %b %d\n",xx,xx)
        xx='z'>>2
        fmt.Printf("z>>2  %b %d\n",xx,xx)
        xx='z'>>3
        fmt.Printf("z>>3  %b %d\n",xx,xx)
        xx='z'>>4
        fmt.Printf("z>>4  %b %d\n",xx,xx)
        xx='z'>>5
        fmt.Printf("z>>5  %b %d\n",xx,xx)
        c:=31
        fmt.Printf("31 %b %d\n",c,c)
        xx='a'
        d:=uint(xx&31)                         // bitwise & 31 -> 11111
        fmt.Printf("uint(a&31) %b %d\n",d,d)
        c=31
        fmt.Printf("%b %d\n",c,c)
        xx='z'
        cc:=uint(xx&31)
        fmt.Printf("uint(z&31)    %b %d\n",cc,cc)
        c=1 << uint(xx&31)
        xx='A'
        fmt.Printf("A  %b %d\n",xx,xx)
        var c1 uint 
        c1=uint(xx&31)
        fmt.Printf("uint(xx&31)    %b %d\n",c1,c1)
        c1=1<<uint(xx&31)
        fmt.Printf("1 << uint(xx&31)    %b %d\n",c1,c1)

//        as[c>>5] |= 1 << uint(c&31)         // as[c>>5] = as[c>>5] | 1 << uint(c&31)
}
