package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
