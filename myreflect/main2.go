package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Printf("value: %d\n", reflect.ValueOf(x))
	fmt.Println("value: ", reflect.ValueOf(x)) // reflect interigates value and does a itoa conversion.
	//
	fmt.Println("Kind(): ", reflect.ValueOf(x).Kind())   // reflect interigates value and does a itoa conversion.
	fmt.Println("Type(): ", reflect.ValueOf(x).Type())   // reflect interigates value and does a itoa conversion.
	fmt.Println("Float(): ", reflect.ValueOf(x).Float()) // reflect interigates value and does a itoa conversion.
	//
	c1 := make(chan int, 4)
	c1 <- 4
	c1 <- 5
	fmt.Println("len() : ", reflect.ValueOf(c1).Len()) // number of elements in channel buffer
	fmt.Println("ChanDir() : ", reflect.ValueOf(c1).Type().ChanDir())
	fmt.Printf("ChanDir() : %d\n", reflect.ValueOf(c1).Type().ChanDir())
	close(c1)
	//
	type MyInt int
	var z MyInt = 7
	v := reflect.ValueOf(z)
	fmt.Println("Myint Type(): ", v.Type())
	fmt.Println("Myint Kind(): ", v.Kind())
	//
	fmt.Println("print reflect.Value: ", v)
	fmt.Println("print interface: ", v.Interface())
	fmt.Printf("print interface: %d \n", v.Interface())
	//        k := reflect.ValueOf(z).Interface().(int)
	//        fmt.Println("k: ",k);
	//
	var y float64 = 3.4
	p := reflect.ValueOf(&y) // Note: take the address of y.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
        //
        v2 := p.Elem()
        fmt.Println("settability of v2:", v2.CanSet()) //  Now v is a settable reflection object, as the output demonstrates,
        v2.SetFloat(7.1)
        fmt.Println(v2.Interface())
        fmt.Println(y)                                 // changed 3.4 to 7.1
        //
        type T struct {
            A int
            B string
        }
        t := T{23, "skidoo"}
        //
        s := reflect.ValueOf(&t).Elem()      // reference to struct t
        typeOfT := s.Type()
        fmt.Println("s.Type() ",s.Type())
        fmt.Printf("s.Type()  %v\n",s.Type())
        for i := 0; i < s.NumField(); i++ {
            f := s.Field(i)
            fmt.Printf("%d: %s %s = %v\n", i,
                 typeOfT.Field(i).Name, f.Type(), f.Interface())
        }
        //
}
