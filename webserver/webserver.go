package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe("localhost:9922", http.FileServer(http.Dir("/home/ubuntu/dev/src/greengridz/webserver/templates/"))))
}

