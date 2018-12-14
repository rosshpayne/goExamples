package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	data:=make([]byte,100)
	data=[]byte("2006-01-02T15:04:05Z07:00")
	fmt.Printf("\n%v\n",data)
	t1:=time.Now()
	err:=t1.UnmarshalText(data)
	if err!=nil {
	    log.Panic(err)
	}
	fmt.Printf("\n%v\n",t1)
}

