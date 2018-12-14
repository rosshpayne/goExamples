package main
/*
	tests if an undefined channel ie. nil value


	result:  nil channel blocks on sending or receiving.  Blocking is good as it gives the opportunity for active channels to work.
*/

import "fmt"
import "time"
import "os"

func main() {
	var ready chan int
	done := make(chan struct{})

	//fmt.Printf("%v", os.Args)
	if len(os.Args) > 1 {
		if os.Args[1] == "Y" {
			ready = make(chan int)
			go func () {
				ready<- 5
			}()
		}
	}
	fmt.Printf("\nready channel:  %#v\n",ready)
	go func() {
		time.Sleep(5000 * time.Millisecond)
		close(done)
	}()
        fmt.Println(" About to select..")
	select {
	case ready<-1:					// when ready is nil it blocks permanently. if not nil it will block waiting for send or simply not block.
			fmt.Println("case ready..")
	case <-done:
			fmt.Println("done terminated")
			return
		//default: 
		//	fmt.Println("case default ")
	}
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("End...")
}
