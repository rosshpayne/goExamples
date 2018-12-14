package main

import (
	"fmt"
)

type height struct {
	cm int
}
var mymap map[string]height = map[string]height {
				"abc": height{1},
				"def": height{2},
				"hij": height{3},
				"klm": height{4},
			}

type LogLevelType uint

const (
        // LogOff states that no logging should be performed by the SDK. This is the
        // default state of the SDK, and should be use to disable all logging.
        LogOff LogLevelType = iota * 0x1000

        // LogDebug state that debug output should be logged by the SDK. This should
        // be used to inspect request made and responses received.
        LogDebug
        LogDebug1
        LogDebug2
        LogDebug3
)


func main () {
	fmt.Printf("\n Logoff = %012b %[1]d",LogOff)
	fmt.Printf("\n LogDebug  = %b %[1]d",LogDebug)
	fmt.Printf("\n LogDebug1 = %b %[1]d",LogDebug1)
	fmt.Printf("\n LogDebug2 = %b %[1]d",LogDebug2)
	fmt.Printf("\n LogDebug3 = %b %[1]d",LogDebug3)
	return
	var k int = 99
        //var v string = "local glasses"

        fmt.Printf("1<<8 %b\n",1<<8)
	fmt.Println(mymap)
        
        if _,ok  := mymap["abc"] ; !ok {
			fmt.Println("no value for abc")
	}
        fmt.Println("abc: ",mymap["abc"])
        if v,ok  := mymap["ser"] ; !ok {
			fmt.Println("no key for ser")
			fmt.Printf(" however map returns a value (zero value of type): %#v\n", v)
	}
        mymap["ser"] =height{22}
        fmt.Println("ser: ",mymap["ser"])
        fmt.Println("ser2: ",mymap["ser2"])				// ser2 does not exist but map will output zero value, hence use v,ok := map["ser2"]
        fmt.Println(k)
	for k,v := range mymap {					// k,v are local to for even if k exists outside for, k in for will override k in outer scope
		fmt.Printf("%s %d\n",k,v)
 	}
        fmt.Println(k)
}

