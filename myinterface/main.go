package main

import (
       "io"
       "os"
       "fmt"
       "strings"
       )
type response struct {			// vars assigned this type are IV's.  They can contain ANY TYPE provide it has method Read(), Close() 
          io.ReadCloser			// interface :  concrete  type has only to satisfy this  interface.
}

// type assertion gives us a way of examining the concrete type (aka dynamic type) stored in an interface type (aka IV) 
//  remember an interface type (value) can store any type as the concrete type, provide that type satisfies the interface.
// There two semantics are available:
// type assertion 1:    x.(T) === iv.(T)     where x is an expression of an interface type and T is the asserted type. Checks x's concrete type matches T.
// type assertion 2:    x.(T) === iv.(T)     where x is an expression of an interface type and T is an asserted interface.  Check's whether x's concrete type satisfies T

func main () {
             //var r1 response
             //var r2 response
             var buf = make([]byte,40) 
             rd := strings.NewReader("CATINTHEHAT first reader ")
             fc,err := os.Create("/Users/rosspayne/go/src/gopl.io/myinterface/test.out")
             if err != nil { 
		panic(err) 
 	     }
	     //
	     //  test you can write to file.
	     // 
	     n,err:=fc.Write([]byte("Hello world.."))
	     if err != nil {
			panic(err)
	     } else {
		fmt.Printf("\nWritten bytes %d",n)
	     }

             //
             //  Body is an interface value.  In the response struct its type is io.ReadCloser which is an interface.
             //  This means that Body must support the following methods
             //
             //              Read(p [] byte) (n int, err error)
             //              Close() (err error) 
             //
             //  Using composition by struct embedding we can add support for io.ReadCloser.
             //  Here we define a struct with two anonymous interface fields of non-primitive types - interfaces.
             //  The struct has all the methods supported by the interfaces promoted from its field values.
             //  Once the struct is assigned values, rd,fc, the Struct assigned to Body has access
             //  to the Read & Close methods in rd, rc.
             //
	     r1:= struct {				// struct literal will allocate the struct. Has two anonymous fields comprising interface.
							//  assign Body a concrete type (of type struct) that satisfies the ReadCloser interface
  		io.Reader				//  these methods of the underlying concrete types, rd,fc will be promoted to r1.Body
                io.Closer				// transparent field name: Closer
  		}{
  		rd,              // only supports io.Readern & io.Writer
                fc,              // supports io.Reader & io.Closer. DOes not support io.Writer
  		}
             //
             n,err =r1.Read(buf)			//read from rd reader  - this method was promoted from anonymous field io.Reader (hiddent name: Reader)
						        //  no need to reference as r1.Reader.Read(buf)

             if err != nil { panic(err) }
             fmt.Printf("\nr1.Read(buf) got   %d  %s\n",n,string(buf))
	     //
	     //n,err=r1.Read(buf)				// read from os file, cannot be done becuase file is associated with io.Closer field.
	
             // fmt.Println("%d  %s",n,r1.Name())      // r1.Body.Name undefined (type io.ReadCloser has no field or method Name)
             //err = r1.Close() 			    // this will call the close() method of the fc instnace.
             //if err != nil { panic(err)  }
             //fmt.Println("File closed..") 
             //
             // type assertion	 -  does the r1.Closer concrete type  also satisfy the io.Writer interface, i.e. does one of its concrete type support this method.
	     //                     NB: we cannot write r1.(io.Writer) as the compiler will complain saying:
	     //                              **** invalid type assertion: r1.(io.Writer) (non-interface type struct { io.Reader; io.Closer } on left)
             //
             //  v,ok := r1.(io.Writer)		
	     fmt.Printf("\nr1.Closer is type:  %v",r1.Closer)
             v,ok := r1.Reader.(io.Writer)			// does r1.Closer concrete type  also satisfy the io.Writer interface.
	     if ok {
		fmt.Printf(" The type is %T\n",v)
                n,err=v.Write([]byte{'a','s','d'})
		if err == nil { 
			fmt.Printf("\n bytes written to fle: = %d", n)
		} else {
			panic(err)
		}
	     }
             v2,ok:=r1.Reader.(io.Reader)			// does r1.Closer concrete type  also satisfy the io.Reader interface.
             if ok {
		 n,err = r1.Read(buf)
		 fmt.Println(" use io.Reader ",string(buf))
	     }
             err = r1.Closer.Close() 			    // this will search for the Close() method in the struct r1 which is associated with the fc concrete type.
             if err != nil { 
		panic(err)  
	     }
             fmt.Printf("\nFile closed..") 

/*
             v = r1.Closer.(*os.File)				// is the concrete type an *os.File)
	     if ok {
                err:=v.Closer.Close()
		if err != nil {
			panic(err)
		}
             }
*/
/*					//  this will fail to compile...as 
	     r2 = struct {                         // the  x field does not promote its Read() method to r2.Body because its named (x)
                x io.Reader				//   the io.Closer is an anonymous embedded type so its concrete type will expose its Close() method.:w
                io.Closer
                }{
                rd,              // only supports io.Reader
                fc,              // supports io.Reader & io.Closer
                }
*/
}
