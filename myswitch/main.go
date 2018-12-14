package main

import (
	"fmt"
)

func main() {
	op := 4
	switch {                          // no operand specified. Like "switch true { "
					  //  Switch checks for first case that is true (from the top) runs commands and automatically breaks
 					  // to force evaluation of other cases specifiy fallthrough.
        case op<4 : 
                fmt.Println("4")
//                fallthrough		 // first case to return true .  break out of switch
        case op==2 : 
                fmt.Println("2")
                fallthrough		// must keep specifying fallthrough to keep falling through
        default:
                fmt.Println("Default")  // will execute default but only on fallthrough if another case is matched.
        case op<3 : 
                fmt.Println("3")
		break;			// this is default behaviour
	}
        // switch op {                      // operand type determines type used in case if op is int then case must specify int.
	fmt.Println("========================")
	switch op {
	 case 4:
         case 3: fmt.Println(3)	
 	 case 2: fmt.Println(2)
         default: 
	}
}
