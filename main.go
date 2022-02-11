package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World\n")
}

func main() {
	// "/" the root path
	http.HandleFunc("/", helloHandler)

	// nil is the default implementation of a HTTP multiplexer
	http.ListenAndServe(":8080", nil)
}
