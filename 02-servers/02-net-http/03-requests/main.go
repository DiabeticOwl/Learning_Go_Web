// Program that sets up a TCP server using the "net/http" package.
// The template parsed in this server contains a form for the user
// to interact with.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
)

type person struct {
	Name    string
	Surname string
	Age     int
}

// ServeHTTP makes type person an http.Handler type.
// ServeHTTP will parse a template containing a form for the user and
// print the values submitted by him/she.
func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ParseForm needs to be run before the use of the property
	// of *http.Request called "Form".
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// Extra data from the variable "r" to be passed to the template.
	data := struct {
		Method        string
		URL           *url.URL
		FormData      map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}

	// The template will be initiated and executed with the ResponseWriter
	// and Request's Data.
	tpl := template.Must(template.ParseGlob("*.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", data)

	// If the form has values the object's attributes will be populated.
	if len(r.Form) > 0 {
		p.Name = r.FormValue("name")
		p.Surname = r.FormValue("surname")
		p.Age, _ = strconv.Atoi(r.FormValue("age"))
	}

	// Each time that the form is submitted,
	// the program will print the person's details.
	fmt.Println(p)
}

func main() {
	var p person

	// Enables a TCP server in the assigned port.
	http.ListenAndServe(":8080", p)
}
