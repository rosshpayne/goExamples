package main

import (
	"bytes"
        "strconv"
	"fmt"
	"log"
)

func main() {
	var buf = new(bytes.Buffer) 				// declaration. Allocation zero initial size.   pointer to bytes.Buffer.  
	//var buf = bytes.NewBuffer([]byte)
        //var buf = bytes.NewBufferString(string)

	fmt.Printf("buf size:  %d %d", buf.Len(), buf.Cap())
	//
	// buf is a struct, to satisify interface io.Writer need pointer to struct hence &
	//
	//A Logger represents an active logging object that generates lines of output to an io.Writer. Each logging operation makes a single call to the Writer's Write method, *bytes.Buffer in ths case. A Logger can be used simultaneously from multiple goroutines; it guarantees to serialize access to the Writer.
        //
	logger := log.New(buf, "loggerX: ", log.Ltime)     // logger writes to io.Writer, not to stdout.
        //
	logger.Print("Hello, log file!") // callls byte.Write(..) to write to buffer buf. buffer expands to accommodate data
        for i:=0;i<500;i++ {
	   logger.Print("line...................................."+strconv.Itoa(i))  // each logger Print writes separate line
	   logger.Printf("buf size:  %d %d", buf.Len(), buf.Cap())
        }

	fmt.Print(buf) // print buffer containing logger entries to stdio
}



