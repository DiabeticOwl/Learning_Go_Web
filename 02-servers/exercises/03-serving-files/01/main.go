// # ListenAndServe on port 8080 of localhost

// ## For the default route "/"

// Have a func called "foo" which writes to the response "foo ran".

// ## For the route "/dog/"

// Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response `<h1>This is from dog</h1>` and also shows a picture of a dog when the template is executed.

// Use "http.ServeFile" to serve the file "dog.jpeg"
package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	// This create the address used in the template.
	http.HandleFunc("/dog.jpeg", dogW)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "foo ran")
}

func dogW(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpeg")
}

func dog(w http.ResponseWriter, r *http.Request) {
	// If tpl is used for parsing the "*.gohtml" file/s reloading the page
	// will panic the application as the template was already parsed once.
	tpl = template.Must(template.ParseGlob("*.gohtml"))
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}
