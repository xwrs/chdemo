package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/api/down", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Shutting down!")
		go os.Exit(1)
	})
	fmt.Println("Service started")
	http.ListenAndServe(":80", nil)
}
