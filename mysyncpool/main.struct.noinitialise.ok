package main

import (
	"fmt"
        "sync"
	"time"
)
type element_ struct { 
              start   time.Time
              end     time.Time
              duration time.Duration
              location string
              orders    map[string][]string
                  }
  
// The Pool's New function should generally only return pointer
// types, since a pointer can be put into the return interface
// value without an allocation:

//var bufPool = sync.Pool{ New: func() interface{} { return new(element_) }, }   // struct literal using field names hence add , at end.
var bufPool sync.Pool

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(key, val string) {
	//b := bufPool.Get().(*element_)
	b := new(element_)
        b.start = time.Now()
        time.Sleep(2*time.Second)
        b.end   = time.Now()
        b.location = string("abcderf") 
        //b.orders = make(map[string][]string,10)
        //b.orders[key]=[]string{val,"abc"}
        b.orders=map[string][]string{key : []string{val,"abc"} }
	bufPool.Put(b)
	bufPool.Put(b)
}

func main() {
	Log( "path", "/search?q=flowers")
	buf:= bufPool.Get().(*element_)     // Get() returns interface{}, so must use Type Assertion to reveal value.
        fmt.Printf("\n%#v\n",buf)
	bufi := bufPool.Get()//.(*element_)     // Get() returns interface{}, so must use Type Assertion to reveal value. If pool is empty Get will run New and create element.
        if bufi  != nil {
             buf = bufi.(*element_)
             fmt.Printf("\n%#v\n",buf)
        } else {
             fmt.Println("bufi is nil")      // will never be nill as pool will generate new struct if empty on Get()
        }
	bufi  = bufPool.Get()//.(*element_)     // Get() returns interface{}, so must use Type Assertion to reveal value. If pool is empty Get will run New and create element.
        if bufi  != nil {
             buf = bufi.(*element_)
             fmt.Printf("\n%#v\n",buf)
        } else {
             fmt.Println("\nbufi is nil")      // will never be nill as pool will generate new struct if empty on Get()
        }
    
}
