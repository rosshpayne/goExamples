package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func main() {

	// test your handler code

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!"+r.Host+" "+r.RemoteAddr+"</body></html>")
	}

	req := httptest.NewRequest("GET", "http://oracle.com/login", nil)	// simulates a http request to  your handler
	w := httptest.NewRecorder()
	handler(w, req)								// execute handler with the simulated request

	resp := w.Result()							// now examine the response fro handler...
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	h:=w.Header()
	for k,v:= range h {
		fmt.Println(k,v)
	}

}
