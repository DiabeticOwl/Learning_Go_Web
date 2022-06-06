// # Serve the files in the "starting-files" folder

// ## To get your images to serve, use

// ``` Go
//  func StripPrefix(prefix string, h Handler) Handler
//  func FileServer(root FileSystem) Handler
// ```

// Constraint: you are not allowed to change the route being used for images in the template file
package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	fs := http.FileServer(http.Dir("./starting-files/public"))
	fs = http.StripPrefix("/public", fs)

	http.Handle("/public/", fs)
	http.HandleFunc("/", sTpl)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sTpl(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("./starting-files/templates/*.gohtml"))

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
