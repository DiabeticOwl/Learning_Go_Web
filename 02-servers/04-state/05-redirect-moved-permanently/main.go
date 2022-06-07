// Program that will redirect the user using the HTTP Status Server Response
// 301 or "Moved Permanently".
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)
}

// bar will print the received method and redirect the user to the
// given location and http.StatusMovedPermanently. "Moved Permanently" or 301 will
// save the method sended in the browse's cache so every time the user access
// this address it will be directly redirected to the location.
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar is:", r.Method)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// foo will print the received method and execute a template with a form.
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo is:", r.Method)
}
