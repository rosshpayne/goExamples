package main_test

import (
	"log"
	"net/http"
	"io/ioutil"

	"testing"
	)

func TestShow (t *testing.T) {
	resp, err := http.Get("http://localhost:9080/show")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(body))
}	

