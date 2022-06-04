// Take the previous program in the exercise #1 and change it so that:
// * a template is parsed and served
// * you pass data into the template
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
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dFunc)
	http.HandleFunc("/me/", meFunc)

	http.ListenAndServe(":8080", nil)
}
