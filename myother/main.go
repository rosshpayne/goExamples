package main


import (
	"net/http"
//	"strings"
        "fmt"
	"flag"
        "reflect"
)

const (
	msAsync      = 1 << iota // perform asynchronous writes
	msSync                   // perform synchronous writes
	msInvalidate             // invalidate cached data
)

var (
	commonLowerHeader = map[string]string{} // map literal - empty element. Simple way to declare and allocate a map. ALternate to using make()
)

type node struct {
         b  []byte
            }

func main () {
	fmt.Printf("%T\n",flag.CommandLine)
        fmt.Printf("%08b\n",msAsync)
        fmt.Printf("%08b\n",msSync)
        fmt.Printf("%08b\n",msInvalidate)
        z:=node{b: nil}
        fmt.Println(len(z.b))
        var myint = []int{1,4,5,6,7,888,6,5,4}
        fmt.Println(fmt.Sprint(myint))
	// commonCanonHeader :=   map[string]string{}                       // map literal, emtpy element - equivalent to the following
	 var commonCanonHeader map[string]string
         commonCanonHeader=make(map[string]string)
        fmt.Println("typeOf: ",reflect.TypeOf(commonCanonHeader).Kind())
        fmt.Println("typeOf: ",reflect.TypeOf(commonCanonHeader).Elem())
        //
        fmt.Printf("%08b %d\n",0x9,0x9)               // 0x  hexidecimal
        fmt.Printf("%08b %d\n",0x16,0x16)
        fmt.Printf("%08b %d\n",0x20,0x20)
        o:= 0666                                      // 0   octal
        fmt.Printf("%d %[1]o %#[1]o %08b\n",o,o)
        fmt.Printf("%032b\n", uint32((1<<16 - 1)<<16 ) )
        fmt.Printf("%032b %[1]d \n", uint32(1<<24 - 1))
        fmt.Printf("%08b %[1]d \n", byte(uint32(16 << 20)>>18) )
        // 
        //
	for _, v := range []string{
		"accept",
		"accept-charset",
		"accept-encoding",
		"accept-language",
		"accept-ranges",
		"age",
		"access-control-allow-origin",
		"allow",
		"authorization",
		"cache-control",
		"content-disposition",
		"content-encoding",
		"content-language",
		"content-length",
		"vary",
		"via",
		"www-authenticate",
	} {
		chk := http.CanonicalHeaderKey(v)
                fmt.Println(len(commonLowerHeader[chk]))
		commonLowerHeader[chk] = v
		commonCanonHeader[v] = chk
                fmt.Println(commonCanonHeader[v])
	}
}
