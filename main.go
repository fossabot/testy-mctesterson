package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm a teapot!")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
