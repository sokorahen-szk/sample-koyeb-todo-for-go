package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: nil,
	}
	http.HandleFunc("/", handler)
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb")
}
