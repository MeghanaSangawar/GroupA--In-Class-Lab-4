package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the handler function for the home endpoint ("/") page
	// and return a message to the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the API Server!")
	})

	// Define the handler function for the about endpoint ("/about")

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is an About Page \n Group-A In Class Lab 4 \n")
	})

	// Start the server on port 8080
	// Print a message to the console

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
