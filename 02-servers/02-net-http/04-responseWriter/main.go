// Program that sets up a TCP server using the "net/http" package.
package main

import (
	"fmt"
	"net/http"
)

type person struct {
	Name    string
	Surname string
	Age     int
}

// ServeHTTP makes type person an http.Handler type.
func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// By using the method "Header" we can assign new header attributes or
	// modify existing ones.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("My-Header", "This is my header :D.")

	fmt.Fprintf(w, "<h1>I'm a human that writes. My name is %v. :D</h1>", p.Name)
}

func main() {
	p := person{
		Name:    "Johann",
		Surname: "Cruz",
		Age:     80,
	}

	// Enables a TCP server in the assigned port.
	http.ListenAndServe(":8080", p)
}
