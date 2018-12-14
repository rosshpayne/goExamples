package main

import (
	"context"
	"fmt"
	"time"
)
type emptyStruct struct{a int}

func main() {
        var x1 *int = new(int)
        var x2 *int = x1
        *x1=22
        *x2=33
        fmt.Println(*x1,*x2)

	var y1 *emptyStruct = new(emptyStruct)
	var y2 *emptyStruct = y1
        fmt.Println(*y1,*y2)

	fmt.Printf("emptyContext type: %T\n",context.Background())

	//d := time.Now().Add(5 * time.Second)
	//	ctx, cancel := context.WithDeadline(context.Background(), d)

	ctx, cancel := context.WithCancel(context.Background())			// pass ctx to any goroutine you like and we can "cancel" very easily.

	fmt.Printf("%#v\n",ctx)

	//if cctx,ok := ctx.(context.cancelCtx); ok {				// cannot refer to unexported type context.cancelCtx
		//fmt.Println("ctx is a cancelCtx using type assertion")
	//}
	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	//time.Sleep(6*time.Second)
	time.Sleep(4*time.Second)
	cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("Done: ",ctx.Err())
	}

}
