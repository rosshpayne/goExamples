package main

import (
	"os"
        "log"
        "bufio"
        "fmt"
       )

var newline byte = '\n' 

func main() {
        //
        // open file
        //
	file, err := os.Open("/Users/rosspayne/go/src/gopl.io/mylog/main.go")    // file is *os.File
        if err != nil {
                      log.Fatal(err);
                      }
        //
        //  create scanner - allocates internal buffer or you can create own buffer and assign.
        //
        var linebuf []byte
        reader := bufio.NewReader(file)       // ends up calling (*Reader).Read & (*os.File).Read
                                              // file has no buffer as the file represents the buffer
        //reader := bufio.NewReaderSize(file,2*1024)
        fmt.Printf("Reader's internal Buffer size: %d\n",reader.Buffered) 
        //
        // read from file
        //
	for {
		linebuf,err = reader.ReadSlice(newline)  // reader calls (*os.File).Read(Buffer) & returns Buffer
                
                fmt.Printf("len, cap : %d  %d ",len(linebuf),cap(linebuf))
                if err != nil {
                             log.Fatal(err)
                }
                fmt.Printf("%s",linebuf)
	}
}
