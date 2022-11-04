package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "api call for: %s\n", r.URL.Path)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	fmt.Printf("Server is started on port 80")
	http.ListenAndServe(":80", nil)
}
