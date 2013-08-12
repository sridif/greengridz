package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe("demo.greengridz.com:9921", http.FileServer(http.Dir("/home/ec2-user/dev/src/greengridz/webserver/templates/"))))
}

