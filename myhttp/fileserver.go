package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":9003", http.FileServer(http.Dir("/usr/share/doc"))))
}
