package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "OK - Healthy")
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
