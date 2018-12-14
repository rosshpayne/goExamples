package main

import (
	"fmt"
	"reflect"
	"io"
	"os"
	"bytes"
	"bufio"
)

func main() {

	//var bfc []byte = []byte{'A','B','C'}
	//bfc := []byte{'A','B','C','d'}                 // declare using either var or T{} literal ie. struct literal`
        bfc := `ABCd`                                    // declare as inferred type string and populate with raw string form. Must convert to []byte where necessary
        is  := []int{}

        fmt.Printf("is len,cap : %d, %d\n",len(is), cap(is))
        for i:=0; i< 10; i++ {
		is=append(is,1)
                is[i]=i*5
        }
        fmt.Printf("is len,cap : %d, %d\n",len(is), cap(is))
        fmt.Println(is)

	var r io.Reader
        var any interface{}

	bfc_out:=make([]byte,10)

//      It's important to be clear that whatever concrete value r may hold, r's type is always io.Reader: 
//      Go is statically typed and the static type of r is io.Reader.
//      However r's internal dynamic type maybe any type that supports the interface io.Reader.
//       so we have a static type and its dynamic type when referring to IVs

	ft := reflect.TypeOf(r)                               // ft is an interface ie. reflect.Type.  Note internally go applies conversion interface{}(r) to argument
        fmt.Println("Dynamic Type:  ",ft)		      // this will use ft's String() method ie. ft.String() will print its representation of type.

	r = os.Stdin

	ft = reflect.TypeOf(interface{}(r))
        fmt.Println("Dynamic Type:  ",ft)

	r = bufio.NewReader(r)

	ft = reflect.TypeOf(interface{}(r))
        fmt.Println("Dynamic Type:  ",ft)

        fmt.Println("************ NewBuffer([]byte(bfc)) **************")
//	r = new(bytes.Buffer)
	r = bytes.NewBuffer([]byte(bfc))

	ft = reflect.TypeOf(interface{}(r))
        fmt.Println("Dynamic Type:  ",ft)
	fv := reflect.ValueOf(r)
        fmt.Println("Value ",fv)

	if r,ok := r.(io.Writer); !ok {
	     fmt.Println(" R is not an io.Writer")
        } else {
	     ft = reflect.TypeOf(r)
             fmt.Println("Dynamic Type:  ",ft)
	     r.Write([]byte(bfc))                         // write to bytes.Buffer
//             now read from the buffer

             w := r.(io.Reader)                   // r's static type is io.Reader, but its dynamic type (i.e. concrete type) is anything that support io.Reader
	     ft = reflect.TypeOf(interface{}(w))
             fmt.Println("w's dynamic type:  ",ft)

	     fv := reflect.ValueOf(w)
             fmt.Println("Value ",fv)

             w.Read([]byte(bfc_out))
	     fmt.Println([]byte(bfc_out))

             any = w				// assign IV to ANY IV, this means it supports any methods, however lets restrict to those we know are in w's dynamic Type.
	     rw := any.(io.ReadWriter)
	     rw.Write([]byte(bfc))
	     rw.Write([]byte(bfc))
	     rw.Write([]byte(bfc))
             rw.Read(bfc_out)
	     fmt.Printf("%s\n",bfc_out)
	     
        }
}

