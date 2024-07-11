package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the API Server!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is an About Page \n Group-A In Class Lab 4 \n")
	})
}
