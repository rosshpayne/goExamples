package main

import (
	"fmt"
)


type inode struct {
	flags uint32
	pgid  uint
	key   []byte
	value []byte
}
type inodes []inode
func main() {
        inodes_:=make(inodes,0,5)				// make a slice
        fmt.Println(len(inodes_))
        inodes_ = append(inodes_, inode{})             // len increased by 1
        fmt.Println("inodes_ ",len(inodes_),cap(inodes_))
        //
        inodes_[0].flags=0x3
        inodes_[0].pgid=99
        fmt.Println(len(inodes_),inodes_)
        //inodes_ = append(inodes_, inodeo{})                       elements
        fmt.Println("src len : ",len(inodes_[0:]) )    // len = 1 0:len()-1 0:0 1     i=0,j=1 so if len is (j-1)-i + 1 = 1 where index[0] has len 1, index[2] has len 1
        fmt.Println("src len : ",len(inodes_[1:]) )    // len = 0 1:0           0     i=1,j=1 so len is     1-1-1+1 = 0
        fmt.Println("src len : ",len(inodes_[1:1]) )   // len = 0 1:0           0     i=1,j=1                         0
        fmt.Println("src len : ",len(inodes_[0:1]) )   // len = 1 0:0    i=0,j=1              (1-1)-0 + 1 =1
        fmt.Println("src len : ",len(inodes_[0:0]) )   // len = 0 0:-1    i=0,j=0              (-1)-0 + 1  =0
        //fmt.Println("src len : ",len(inodes_[2:]) )   //  slice bounds out of range
        //
	copy(inodes_[1:], inodes_[0 :])       // expand []inodes   note: dst has len 0 (see above), so no copy is done.
        fmt.Println(len(inodes_),inodes_[0])
}
