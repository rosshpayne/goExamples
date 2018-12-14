// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"
	"time"
	_ "fmt"
	"log"
	"context"

	memo "gopl.io/ch9/memo-9.3"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody
var wait int = 1900

func TestSeq(t *testing.T) {
   	defer func() {						// recover can only run in defer. Here is an anonymous func.
       	if err := recover(); err != nil {
	     t.Error((err.(memo.Error)).Error())
       	 }
    	}()
	log.Printf("\nTimeout: %d\n",wait)
	//
	//  using context from client side applies timeout to all Get-http requests, not inidividual Http requests as previous iterations have.
        //   Alternatively we could apply context to memotest HTTP function
	//
	ctx,cancelFunc:=context.WithTimeout(context.Background(),time.Duration(int64(wait))*time.Millisecond)
	defer cancelFunc()

	m := memo.New(ctx, httpGetBody )
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestChanWrongType(t *testing.T) {
   	defer func() {						// recover can only run in defer. Here is an anonymous func.
       	 if err := recover(); err != nil {
	     e:=err.(memo.Error)	
	     t.Error(e.Error()) 	
       	 }
    	}()
	t.Skip()
	var wait float64 = 64
	log.Printf("Timeout: %d\n",wait)

	m := memo.New(nil, httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestTooManyArgs(t *testing.T) {
   	defer func() {						// recover can only run in defer. Here is an anonymous func.
       	 if err := recover(); err != nil {
	     e:=err.(memo.Error)	
	     t.Error(e.Error()) 	
       	 }
    	}()
	t.Skip()
	log.Printf("Timeout: %d\n",wait)
	m := memo.New(nil, httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestNilContext(t *testing.T) {
	m := memo.New(nil, httpGetBody )
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestConcurrent(t *testing.T) {
	wait:=500

	ctx,cancelFunc:=context.WithTimeout(context.Background(),time.Duration(int64(wait))*time.Millisecond)
	defer cancelFunc()

	m := memo.New(ctx, httpGetBody )
	defer m.Close()
	memotest.Concurrent(t, m)
}
