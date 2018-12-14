package main

import (
	"fmt"
	"time"
	"reflect"
	 "unsafe"
	 _ "os"
)
type emb struct{
	a int
	b float64
}
func (v *emb) sum() int { return v.a+int(v.b) }

type Animal struct {
	a     []string			`ab"'c`
	st    [3]float32
	cnt   int			`protobuf:"1"`
        err   error
        fld   []int
//	emb
}

var sendValue int = 5

func main() {
	var a1 Animal
	var refi interface{}

 	a1=Animal{[]string{"as","asdf"},[3]float32{2.34,45.34,8886.433},1,nil,[]int{1,2,3,4}}  //,emb{a:2,b:3.44} }
        refi=a1

 	rv:=reflect.ValueOf(refi)
        fmt.Println("Kind():  ",rv.Type().Kind())
        fmt.Println(rv.Type().NumField())
        // fmt.Println("Elem()   ",rv.Elem())           // Panic's as rv is not an interface nor a pointer which is required for an Elem()
	fmt.Println("Field 3 value:   ",rv.Field(2) )
	fmt.Println("Field 1 value:   ",rv.Field(0) )

	for i := rv.Type().NumField(); i>0; i-- {
		fmt.Printf("\nField:  %#v",rv.Type().Field(i-1))
		fmt.Printf("\nField Tag:   %v",rv.Type().Field(i-1).Tag)
		//
		// make new objects from field
		//
		fmt.Printf("\nMake ArrayOf using field type:   %v",reflect.ArrayOf(5,rv.Type().Field(i-1).Type) )    // ArrayOf(..) returns Type
		fmt.Println()
	}

        //is:=[]int{2, 3}
	//fmt.Printf("\nValue  FieldByIndex:  ",rv.FieldByIndex(is) )
	//fmt.Printf("\nType  FieldByIndex:  ",rv.Type().FieldByIndex(is) )
	//
	// reflect.ValueOf(chan)
	//
        var chant chan int
        //refi=chant

	//rt:=reflect.ValueOf(chant).Type()				// pass chant to interface{} argument directly.

	ft:=reflect.TypeOf(chant)					// 
        fmt.Println(ft.Kind());
        fmt.Println(ft.String());

	//intt:=reflect.TypeOf(sendValue).Kind()
	intt:=reflect.TypeOf(sendValue)
	fmt.Printf("\n%T %[1]v",intt.String());					// string int

        //rt:=reflect.ChanOf(reflect.BothDir,intt )				// create channel using ChanOf of type int
	rt:=reflect.ChanOf(reflect.BothDir,reflect.TypeOf(sendValue))
        fmt.Printf("\nChanOf: %T  %[1]v ",rt.Kind())
        fmt.Printf("\nChanOf: %T  %[1]v ",rt.String())
        fmt.Printf("\nString(): %T  %[1]v  ",rt.String())
        fmt.Printf("\nKind(): %T  %[1]v  ",rt.Kind())
	fmt.Printf("\n\n")
	//
	// make channell with reflect.MakeChan
	//
	fmt.Println("Now Make a channel using reflect.MakeChan with buffer 1")
	rv =reflect.MakeChan(rt,1)
	fmt.Println("string    ",rv.String())
        fmt.Println("Kind():   ",rv.Kind())
	fmt.Println("Send value on chan using rv.Send(reflect.ValueOf(SendValue)")
        rv.Send(reflect.ValueOf(sendValue))
	fmt.Println(" Now close chain by calling vclose()");
	go func(rv reflect.Value) {
		x,ok:=rv.Recv()
		if ok {
			fmt.Println(" ******** Value received on channel:  ",x)
	        } else {
			fmt.Println(" ******* Value received on channel errored...")
		}
	}(rv)
	time.Sleep(time.Second*2)
	rv.Close()
	//
	// ptr to chan
	//
	rv2:=reflect.New(rt)			// New(type) returns a Value representing a pointer to a new zero value for the specified type.
        fmt.Println("String of New(rt) where rt is a chan:  ",rv2.String())
        fmt.Println("Kind of New(rt):  ",rv2.Kind())
	//
	// Address of rv - chan type buffered 5
	//
//	raddr:=rva.Addr()
	//fmt.Println(raddr.String())
	//
	//  ValueOf(interface{}) returns Value
	//
	ic:=5
        rv=reflect.ValueOf(&ic)					// rv:  ptr to int 5
	//xx:=*(*int)(unsafe.Pointer(&ic))
	fmt.Printf("\n********++++ rv.Pointer(): %d",*(* int)(unsafe.Pointer(rv.Pointer()) ))  // use unsafe.Pointer to access rv.Pointer()

	fmt.Println("\n\nreflect.ValueOf(&ic) where ic=5\n");
        fmt.Println("Value rv.String(): ",rv.String())
        fmt.Println("Value rv.Kind(): ",rv.Kind())
        fmt.Println("rv.Elem()         ",rv.Elem())           // Elem() returns the value that the interface v contains or that the pointer v points to.
        fmt.Println("rv.Elem().Int()   ",rv.Elem().Int()) 
	fmt.Println("\n Now take .Type()..")

	rt=rv.Type()						// rt: rv.Type()
        fmt.Println("Type rt.String: ",rt.String())
        fmt.Println("Type rt.Kind(): ",rt.Kind())
        fmt.Println("CanAddr? : ",rv.Elem().CanAddr())
	//
	// make ptr to int, using New - which is equivalent to a MakePtr like the other MakeChan, MakeFunc, MakeMap, MakeSlice..
	//
	fmt.Println("\n\n New(rv.Elem().Type()) where rv is above..\n")   // rv is a ptr to int derived from ic=5

	rtpv:=reflect.New(rv.Elem().Type())			// MakePtr  equivalent: make a pointer to rv, which is an int of value 5.

        fmt.Println("Value String()  : ",rtpv.String())
        fmt.Println("Value Kind()  : ",rtpv.Kind())
	//
	// Value.Elem() returns the value that the interface v contains or that the pointer v points to. 
	//              It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil
	//
	fmt.Println("rtpv.Elem(): ", rtpv.Elem().String())		// <int Value>   Elem() doesn't print anytng useful, need to use String() to get at it.
        fmt.Println("CanAddr? : ",rtpv.Elem().CanAddr())		// here value is an ptr to int 5
        fmt.Println("CanSet()? : ",rtpv.Elem().CanSet())		
	//
	//  set value and show it
	//
	if rtpv.Elem().CanSet() {		// Elem returns the value that the interface v contains or that the pointer v points to.
		fmt.Println("\n\n Having established we can CanSet a value set it to 99\n")
        	rtpv.Elem().SetInt(99)
	} else {
		fmt.Println("CanSet is false..")
	}
        fmt.Println("rvtp int ",rtpv.Elem())
        fmt.Println("rvtp int ",rtpv.Elem().Int())
}
