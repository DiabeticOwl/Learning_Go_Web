// Program that prints a string with a value passed through the page's URL.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	// Value extracted from the URL: http;//localhost:8080?q=v
	v := r.FormValue("q")

	fmt.Fprintf(w, "My search was: %v", v)
}
