package main

import (
	"context"
	"os/exec"
	"time"
	"log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4000*time.Millisecond)
	defer cancel()

	c:=exec.CommandContext(ctx, "ls","-l","/tmp")
	/*if err := c.Run(); err != nil {
		log.Panic(err)
	}*/
	if out_,err:=c.Output(); err != nil {
		log.Panic(err)
	} else {
		log.Println(string(out_))
 	}
}
