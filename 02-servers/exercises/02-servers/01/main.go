// Exercise #1
// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/"
// "/dog/"
// "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You are in the home page.</h1>")
}

func dFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You are in the dog page.</h1>")
}

func meFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Johann Cruz</h1>")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dFunc)
	http.HandleFunc("/me/", meFunc)

	http.ListenAndServe(":8080", nil)
}
