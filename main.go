package main

import (
	"fmt"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server listening on port 80")
	http.ListenAndServe(":80", nil)
}
