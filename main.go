package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// handle POST localhost/*
	http.HandleFunc("POST localhost/{$}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("localhost %v %v", r.Method, r.URL.Path)
		fmt.Fprintf(w, "localhost %v %v\n", r.Method, r.URL.Path)
	})

	// handle GET /
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		fmt.Fprintf(w, "%v %v\n", r.Method, r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
