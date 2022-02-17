package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting gohello...")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v", r)
		fmt.Fprintf(w, "OK\n")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v", r)
		fmt.Fprintf(w, "Hello, Golang\n")
	})

	panic(http.ListenAndServe(":8080", nil))
}
