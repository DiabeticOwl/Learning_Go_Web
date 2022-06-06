// Program that will redirect the user using the HTTP Status Server Response
// 303 or "See Other".
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)

	http.ListenAndServe(":8080", nil)
}

// barred will print the received method and execute a template
// with a form. The form will send a POST response to "/bar".
func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred is:", r.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

// bar will print the received method and redirect the user to the
// given location and http.StatusSeeOther. "See Other" or 303 will
// send a GET response regardless of the method received.
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar is:", r.Method)

	// This action can be performed by using:
	// `http.Redirect(http.ResponseWriter, *http.Request, location. status-code)`
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

// foo will print the received method and execute a template with a form.
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo is:", r.Method)
}
