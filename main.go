package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8000
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: nil,
	}
	http.HandleFunc("/", handler)
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb")
}
