package main

import (
	"fmt"
)
var decodeMap [256]byte
  
const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
const encodeHex = "0123456789ABCDEFGHIJKLMNOPQRSTUV"


func main() {
        var i int =  99 
        var u uint
       if i < 0 {
	       u = (^uint(i) << 1) | 1 // complement i, bit 0 is 1
       } else {
	       u = (uint(i) << 1) // do not complement i, bit 0 is 0
       }
       //iencodeUnsigned(u)
       fmt.Printf("^uint(i)   %b \n",^uint(i))
       fmt.Printf("^uint(i)   %o \n",^uint(i))
       fmt.Printf("uint(i)    %b \n",uint(i))
       fmt.Printf("uint(1)    %b \n",uint(1))
       fmt.Printf("^uint(1)   %b \n",^uint(1))
       fmt.Printf("(^uint(i)  %b \n",(^uint(i) << 1))
       fmt.Printf("%b \n",u)
 
       encoder := encodeStd

       	for i := 0; i < len(encoder); i++ {
                fmt.Printf("%c %d %b\n",encoder[i],encoder[i],byte(i)) 
  		decodeMap[encoder[i]] = byte(i)
       }

       for i,v:=range decodeMap {
          if v != 0 {
                fmt.Printf("%d [%d]\n",i,v)
          }
       }

       for i:=0;i<32;i++ {
          fmt.Printf("%d %b %x\n",i,i,i)
       }

}
