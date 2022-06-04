// Program that sets up a TCP server using the "net/http" package
// and the DefaultServeMux.
package main

import (
	"fmt"
	"net/http"
)

func p(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a Person.")
}

func c(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a cat.")
}

func d(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a dog.")
}

func main() {
	// HandleFunc takes a pattern (used for assign the URL to which the
	// Handler will point) and a function that takes a ResponseWriter
	// and a Request as parameters. HandleFunc uses the function as Handler.
	http.HandleFunc("/person/", p)
	http.HandleFunc("/cat", c)
	http.HandleFunc("/dog", d)

	// Enables a TCP server in the assigned port.
	// Since the Handler is nil Go will use the DefaultServeMux.
	http.ListenAndServe(":8080", nil)
}
