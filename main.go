package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	port := 8000
	http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb")
}
