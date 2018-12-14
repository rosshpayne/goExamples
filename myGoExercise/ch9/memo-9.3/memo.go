// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import  "errors"
import  "context"
import  "time"
//import  "fmt"
//!+Func

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type Error string						// create error Type which is unique to this package.  We can use type assertion to filter panics to this package.

func (e Error) Error() string {
	return string("** memo error: "+ e)
}


// A result is the result of calling a Func.
type result struct {
	value   interface{}
	dur  	time.Duration
	err     error
}

type entry struct {
	res   result
	ready chan struct{} // each result has a dedicated channel. Channel is closed for cached values however.
	//  I imagine a closed channel consumes next to zero resources as opposed to an open channel.
	//  So you can be generous when dealing with closed channels.
	//  Channel used to synchronise completion of http.Get call.
}


//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // each request to a memo cache has its own response channel
}

type Memo struct {
	ctx      context.Context
	requests chan request   // each function has a cache and its own request channel - which is closed by client.
}

// New returns a memoization of f.  					// Clients must subsequently call Close.    Close will close the request channel.
func New(ctx context.Context, f Func) *Memo { 
	if ctx == nil {
		ctx=context.Background()
	}
	memo := &Memo{ctx: ctx, requests: make(chan request)}         // construct struct and pass its pointer back.
	go memo.server(f) 						// start goroutine cache server.  Runs forever. One for each f function
	return memo
}

// Provide client
//  Package
//   Struct Memo which has method Get & New that are accessible
//   A Type interface M that specifies Get method

func (memo *Memo) Get(key string) (interface{}, error, time.Duration) {	// this is what the client accesses - the Get method.
	var res result
	response := make(chan result)					// create a channel to receive response back from server
	memo.requests <- request{key, response}				// send request to server via channel
	res = <-response 						//   wait for response from server
	return res.value, res.err, res.dur
}

func (memo *Memo) Close() { close(memo.requests) }


func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests { 				// Get requests from Request channel. Get Client to close this channel to shutdown server.
		e := cache[req.key]
		if e == nil {
			t0 := time.Now()
			e = &entry{ready: make(chan struct{}) } 		// make ready channel. One per respose.
			cache[req.key] = e
			go e.call(f, req.key) 				// call f(key)	 goroutine to populate cache.  Gets url data and dies.
			select {
			case <-memo.ctx.Done():
				delete(cache, req.key) 			// remove entry from cache
				t1 := time.Now()
			        res:=result{value: nil, err:  errors.New("Timeout: func call.."+req.key), dur: t1.Sub(t0) }
				req.response <- res
				continue				// jump to beginning of loop i.e. bypass e.deliver. e.ready will be closed in call goroutine, eventually.
			case <-e.ready:					// has ready closed?
			}
			t1 := time.Now()
			e.res.dur=t1.Sub(t0)
		}
		go e.deliver(req.response) 				// goroutine to send response - created and destroyed for each client-response.
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	t0 := time.Now()	
	e.res.value, e.res.err = f(key)
	t1 := time.Now()
	e.res.dur=t1.Sub(t0)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready 							//  each result is assoicated with a ready channel that is closed permanently after call.
	response <- e.res
}

//!-monitor
