package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8000
	http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb")
}
