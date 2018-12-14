package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	h := json.RawMessage(`{"precomputed": true}`)    // conversion:  convert string to RawMessage.  Use RawMessage as h will be json & will not need to 
							//    have its value marshalled
	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))

	if mh,err:=h.MarshalJSON(); err == nil {        // h is a RawMessage, pass through unchanged please.
		os.Stdout.Write(mh)
       }
}
