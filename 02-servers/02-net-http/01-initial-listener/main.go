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
	fmt.Fprintf(w, "I'm a human that writes. My name is %v. :D", p.Name)
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
