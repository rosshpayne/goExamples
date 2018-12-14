package main

import (
       "io"
//       "os"
       "fmt"
       "strings"
       )
type response struct {
          Body  io.Reader
                }
func main () {
             var r1 response
             var buf = make([]byte,20) 
             rd := strings.NewReader("first reader ")
//             fc,err  := os.Open("/Users/rosspayne/go/src/gopl.io/http/main.go")
//             if err != nil { panic(err) }
	     r1.Body = struct {
  		io.Reader
  		}{
  		rd,
  		}
             n,err:=r1.Body.Read(buf)
             if err != nil { panic(err) }
             
   fmt.Println("%d  %s",n,string(buf))
}
