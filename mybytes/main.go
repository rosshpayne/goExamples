package main

import (
	"fmt"
	"sort"
	"bytes"
)

func main() {
        var buf *bytes.Buffer;       // no allocation is performed - only states that buf must be a pointer to bytes.Buffer when allocated later on
	var buf2 bytes.Buffer;	     // allocation of Buffer struct is performed with no underlying (embedded in this case) []byte allocated.
				     //  embedded []byte will be allocated automatically when written to using Write(.)
				     //- however all  methods & functions assoicated with bytes.Buffer require or return *bytes.Buffer - so this 
				     //  declaration is not very useful unless we take & and assign it to another * var.:w

	if buf == nil {
        	fmt.Printf("\n buf Type  %T  %#v- is nill (no allocation)\n" ,buf,buf );
	} else {
        	fmt.Printf("\n buf Type  %T,   pointer  %p \n",buf, buf);
	}
        fmt.Printf("\n buf2 Type pointer %p   %T  %#v- is nill (no allocation)\n" ,&buf2,buf2, buf2 );

	//buf = new(bytes.Buffer)					// this will force the use of the bootstrap slice when a write() occurs
	//fmt.Printf("\nbuf = new(bytes.Buffer)\n")

        localbuf:=[]byte{'C','A','1'}
	buf = bytes.NewBuffer(localbuf)		// equiv to make(struct) in C
						//  this will allocate not use bootstrap slice.	
	//fmt.Printf("\nbuf = bytes.NewBuffer([]byte{'C','A','1'})\n")
	fmt.Printf("\nbuf = bytes.NewBuffer(localbuf)\n")

        localbuf[0]='X'				// shows that localbuf and embedded buf in Buffer are being shared...at them moment
						//  note should not access localbuf directly ones assigned to bytes.Buffer
	fmt.Printf("\n contains: [%c] \n ",buf.Bytes())

	if buf.Len() == 0  {
        	fmt.Printf("\n buf Type  %T  - no embedded []byte allocated though: %#v\n" ,buf,buf);
	} else {
        	fmt.Printf("\n buf Type  %T,   pointer  %p  len %d , cap  %d,  %#v\n",buf, buf, buf.Len() , buf.Cap(), buf);
	}
	//var w []byte = []byte{'A','B','C'}

	_,err:=buf.Write([]byte{'A','B','C'})
        if err != nil {
		panic(err)
	}
        localbuf[0]='Y'				// shows localbuf has been dettached from Buffer embedded []byte due to reallocation of above Write(..)
	fmt.Printf("\n written: %c \n ",buf.Bytes())
        fmt.Printf("\n buf : %#v\n" ,buf,buf);
	fmt.Printf("\n written: %c \n ",buf.Bytes())
	fmt.Printf("\n localbuf: %c \n ",localbuf)


        //fmt.Printf("\n Type  %T,   len  %d \n",buf2, &(buf2).Len());
	a := []string{"1", "2","3", "4", "6", "8", "9"}
	x := "5"
	var exact bool
//        f:=bytes.Compare([]byte("35"),[]byte("3"))
        //fmt.Println(f)
	i := sort.Search(len(a), func(i int) bool {
                ret := bytes.Compare([]byte(a[i]),[]byte(x))
                if ret == 0 {
                        exact = true
                }
              fmt.Println(a[i],x,ret,exact,(ret!=1))
                return ret != -1 
        })

        if !exact && i > 0 {
               fmt.Println("i--",i)
		i--
               fmt.Println("i--",i)
	}
        fmt.Printf("exact %v, i=%d,  x=%s  %v", exact,i,x,a)
}
