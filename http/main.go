package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", HttpHealthCheck)
	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %s\n", err)
		os.Exit(1)
	}
}

func HttpHealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	fmt.Fprintln(w, "ok")
}

func health() error {
	return nil
}
