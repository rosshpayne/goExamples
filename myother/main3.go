package main

import (
    lib "mylib/math"
    "fmt"
)

type handleritf  interface {
          Xstring(arg1 int, arg2 string) string
                      }

type helperfunc struct {
                  f1 int
                  f2 int
                  f3 string
                   }


func (m helperfunc) maybe (a int, b string) string {
              return fmt.Sprintf("this is maybe:  %d , %s",a,b)
              }

type adapterfunc func(a int, b string) string

func (f adapterfunc) Xstring (a int, b string) string 

var (
    fruits  = []string{"apple", "banana", "cherry", "durian"}
    banned = "durian"
)

func main() {
    channel := lib.PermutateWithChannel(fruits)
    for myFruits := range channel {
        fmt.Println(myFruits)
        if myFruits[0] == banned {
            close(channel)
            //break
        }
    }
    x2:=0x7f

    fmt.Printf("%b %d",x2,x2)


    var yy  handleritf

    xx := helperfunc{1,2,"asdf"}
    yy  = adapterfunc(xx)

    xx.maybe(55,"Byebye")
}
