// # Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc
// * **Constraint**: Do not change anything outside of func main
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You are in the home page.</h1>")
}

func dFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You are in the dog page.</h1>")
}

func meFunc(w http.ResponseWriter, r *http.Request) {
	name := "Johann Cruz"

	err := tpl.ExecuteTemplate(w, "index.gohtml", name)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	// Since type "HandlerFunc" has the method
	// "ServeHTTP(w ResponseWriter, r *Request)", this type is also a Handler.
	// Which is what "http.Handle" expects.
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dFunc))
	http.Handle("/me/", http.HandlerFunc(meFunc))

	http.ListenAndServe(":8080", nil)
}
