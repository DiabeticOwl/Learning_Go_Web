// # Serve the files in the "starting-files" folder

// ## To get your images to serve, use only this

// ``` Go
//  fs := http.FileServer(http.Dir("public"))
// ```
package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	// This means that all files and directories in "public" will be tracked.
	fs := http.FileServer(http.Dir("public"))

	// The directory named "pics" in fs will be linked with the URI "/pics/".
	http.Handle("/pics/", fs)
	// The template will be executed and the files in the previously described
	// URI  will be found accordingly.
	http.HandleFunc("/", sTpl)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sTpl(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
