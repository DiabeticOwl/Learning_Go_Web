// Using cookies, track how many times a user has been to your website domain
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read/", read)

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("times-in-website")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "times-in-website",
			Value: "1",
		})
	} else {
		i, _ := strconv.Atoi(c.Value)
		c.Value = strconv.Itoa(i + 1)

		http.SetCookie(w, c)
	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER'S COOKIES.")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("times-in-website")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>This is the count of times you were in the home page: %v</h1>", c.Value)
}
