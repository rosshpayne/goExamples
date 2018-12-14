package main

import (
      "fmt"
      "reflect"
      "unicode/utf8"
       )
var genrune []rune

//  string is an immutable sequence of bytes.
//  Text strings are interpreted as UTF8-encode sequence of Unicode case points - runes.
//  Strings are like arrays  - they are fixed in length (unless resassigned a new string which can be of different length - which is not like an array which cannot be reassigned))
//			     - however, unlike an array a string struct is like a slice (has a pointer to storage of string data)
//                           - but like an array a string is a value type (not reference like a slice)
//                           - can len(s) like array or slice
//                           - arrays have no cap() nor do strings
//  strings are like slice   - as they point to a storage array that can be shared with other strings.
//                           - because a string is immutable the string type is said to still be a value type because a function cannot change its contents.
//                           - so strings a very light when being passed around as its the pointer that gets passed around not the data.




func main () {

      fmt.Println(string(0x4eac))     // 0x is used to specify number as a bit/byte sequence. 
				      //  Number converted to string is interpreted as runes in printf
      fmt.Println("\u4eac")           // escaping using \u in a string says interpret following hex as a rune
      fmt.Println(string('\u4eac'))   // specifying rune literal uses same \u escaping. Must convert to string before printing.
      fmt.Printf("%s\n",string(0x41)) // 'A' - again number converted to string is interpreted as rune
      fmt.Println(string(0x4eac))     // specify rune as number in hex format (0x) convert to string and print as rune 
      fmt.Printf("%x\n",string(65)    // again number converted to string and print as hex
      fmt.Printf("%x\n",65)	      // 0x41
      ax:=0x41				// specify number using hex 0x instead of numbers
      ax2:=ax+1
      fmt.Println(string(ax2))		// string says interpret as ASCII (or more precisely a rune encoded in UTF8) 
      return
      var q string
      fmt.Printf("\n q = [%s]\n",q)
      s := string("the cat sat on the mat")
      t := s
      s = "..and now its different"			// reassign new string value to s - new storage allocated. old string is GC, but it won't be because t is pointing to it
      // s[5]='x' cannot assign to s[5]			//   but cannot change content of s
      fmt.Printf("t = %d   [%s]  \n",len(t),t)
      fmt.Printf("s = %d   [%s]  \n",len(s),s)
      fmt.Printf("s[4:8] %s",s[4:8])
      //
      s = "go/src/gopl.io" + t				// allocate new storage (string) to s ...old string GC
      fmt.Println("************************") 
      fmt.Println(" s=go/src/gopl.io + t")
      fmt.Println(s)
      fmt.Println(s[4:8])
      fmt.Println("len: ",len(s))
      //
      s2 := s[:5]                         // new string
      fmt.Println(s2)
      fmt.Println("len: ",len(s2))
      fmt.Printf("%T\n",s2)
      //
      s = "Left foot"
      t = s					// t shares underlying array with s. 
      s += ", right foot"			// reassign new storage to s
      fmt.Println("len s: ",len(s),s) 
      fmt.Println("len t: ",len(t),t)
      t += "..here"				// go will allocate new underlying array to t, as s already shares the array.
      fmt.Println("len s: ",len(s),s) 
      fmt.Println("len t: ",len(t),t)
      //
      ps , pt := &s, &t                    // pointers to string 
      fmt.Println("*ps : ",*ps)
      fmt.Println("*pt : ",*pt)
      fmt.Printf("ps: %p\n",ps);
      fmt.Printf("pt: %p\n",pt);
      if ps == pt {
         fmt.Println(" ps == pt")
      } else {
         fmt.Println(" ps != pt")        // ptrs are different - even though they have same underlying representation.
      }                                  // is the ptr pointing to the struct {data,len}??
      //
      t2 := t[5:8]
      pt2 := &t2
      fmt.Println(*pt2, *pt)
      fmt.Printf("pt2 type: %T\n",pt2);
      fmt.Printf("pt : %p\n",pt);
      fmt.Printf("pt2: %p\n",pt2);
      if pt == pt2 {
         fmt.Println(" pt == pt2")
      } else {
         fmt.Println(" pt != pt2")
      }
      //
      rt:=reflect.TypeOf(pt)
      rtelem := rt.Elem();
      fmt.Println("Kind():  ",rt.Kind())
      fmt.Println("TypeOf(pt).ELem(): ",rtelem.Kind())
      //
     const nihongo = "日本語"
     // strings are simply an ordered sequence of bytes []byte
     // only string literals & Go code are interpreted as UTF8
     for i:=0; i<len(nihongo); i++ {
           fmt.Printf("%d  %X %d %c\n",i,nihongo[i], nihongo[i],nihongo[i])
     } 
     // range operator interpretes string as UTF8
     for index, runeValue := range nihongo {
           fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
    // alternatively, use unicate/utf8 package to interpret string as UTF8
    fmt.Println(" alternatively, use unicate/utf8 package to interpret string as UTF8")
    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
	//
	s3:="Abc\r\n\tdef"
	fmt.Printf("%d %s\n",len(s3),s3)
	s3 =`Abc\r\n\tdef`
	fmt.Printf("%d %s\n",len(s3),s3)
}
