package main

import (
	"fmt"
	"time"
)

func main() {
	//     c := time.Tick(4 * time.Second)                   // returns channel & sends a "tick" to channel every duration of time
	t := time.NewTicker(5100*time.Millisecond)
	tl:= time.NewTicker(3*5500*time.Millisecond)
	//for now := t.C {                              // use range to receive from channel. Range will block on channel until sender sends.
	//       fmt.Printf("%v      %s      %s   %s\n", now, time.Now(), now.Format(time.Kitchen), now.Format(time.RFC1123Z))
	//}
loopExit:
	for {
		select {
		case t := <-t.C:
			fmt.Printf("ticker: %s\n", t.Format(time.RFC3339Nano))
		case <-tl.C : 
			fmt.Println("Break from  ticker..")
			break loopExit
		}
	}
	fmt.Println("Stopping ticker..")
	tl.Stop()
	t.Stop()
}
