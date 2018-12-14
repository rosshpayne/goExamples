package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	ss:=new(string)
        str:="ABC\u4e16defghijklmnopqrstuvwxyz"
	f, err := os.OpenFile("/tmp/notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f,"writing to file %d", 42)
	f.Sync()
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	f, err = os.OpenFile("/tmp/notes.txt", os.O_RDONLY,0755)
	//var s []byte
	s:=make([]byte,20,20)
	n,err:=f.Read(s)
	fmt.Println("n ",n)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(s))
	f, err = os.OpenFile("/tmp/notes.txt", os.O_RDONLY,0755)
	//var str string
        fmt.Fscanf(f,"%s",s)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
 	fmt.Println("str: ",string(s)) 
	ss=&str
  	s1:=(*ss)[4:8]	
	fmt.Printf("%T   %[1]s   %d\n",s1,len(s1))
}
