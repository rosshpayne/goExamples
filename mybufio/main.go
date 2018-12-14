package main

import (
	"fmt"
	"os"
        "log"
        "bufio"
)

func main() {
        //
        // open file
        //
	file, err := os.Open("/Users/rosspayne/go/src/gopl.io/mylog/main.go")    // file is *os.File
        if err != nil {
                      log.Fatal(err);
                      }
        //  how to read a file, line by line, or word by word
        //  1. create *os.File object - pointer to os.File
        //  2. create scanner - using interface io.Reader to join scanner to os.File
        //  2a. optionally assign scan type - ia Split function
        //  3. Scan Reader to output lines
        //
        //scanbuf := make([]byte,32*1024);
   
        scanner := bufio.NewScanner(file)       // calls (*file).Read
                                                // default internal buffer size MaxScanTokenSize (64kB) 
        //scanner.Buffer(scanbuf,80)
        //fmt.Println(cap(scanbuf),len(scanbuf))

        //scanner.Split(bufio.ScanLines)
        scanner.Split(bufio.ScanWords)
        //
        // read from file
        //
	for {
		if scanner.Scan()  {
			fmt.Printf("%s\n",scanner.Text())
		} else {
                        if err:=scanner.Err(); err != nil {       // EOF will be nil as its not an error
                             log.Fatal(err)
                        } else {
                             break;  // EOF
                        }
                } 
	}
        //fmt.Println(cap(scanbuf),len(scanbuf))
}
