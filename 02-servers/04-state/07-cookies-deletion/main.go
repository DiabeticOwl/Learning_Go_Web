// Program that enables the option of creating, reading and deleting a cookie.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set/", set)
	http.HandleFunc("/read/", read)
	http.HandleFunc("/expire/", expire)

	http.ListenAndServe(":8080", nil)
}

// "index" checks if there is a cookie named "session" and redirects to "/read"
// if there is. Otherwise, will print a text with a link to "/set".
func index(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		fmt.Fprintln(w, `
			<h1><a href="/set">Set a Cookie</a></h1>
		`)
		return
	}

	http.Redirect(w, r, "/read", http.StatusSeeOther)
}

// "set" checks if there is a cookie named "session" and redirects to "/read"
// if there is. Otherwise, will set the cookie with the path to "/" and
// print a text with a link to "/read".
func set(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/read", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
		// If Path is not set Go will assign the URI of the request to
		// the cookie each time it is assigned.
		Path: "/",
	})

	fmt.Fprintln(w, `
		<h1>Cookie Set! Read it <a href="/read">Here</a>!</h1>
	`)
}

// "read" will read the cookie named "session" and print it, alongside a link
// to "/expire". If there is no cookie found, will redirect to "/set".
func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `
		<h1>This is your cookie:</h1>
		<p>%v</p
		
		<br>

		<a href="/expire">Click here to delete it.</a>
	`, c)
}

// "expire" checks if there is a cookie named "session" and changes it's
// "MaxAge" and "Path" attributes. The user will then be redirected to "/".
// If there is not cookie found the user will be redirected to "/set".
func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1
	c.Path = "/"

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
