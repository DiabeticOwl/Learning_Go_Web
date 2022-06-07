// Program that sets a cookie with an UUID.
package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.New()

		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			Path:  "/",

			// Attribute used for enable extra security in HTTPS servers.
			// Secure: true,

			// Prevents any javascript code to access this cookie.
			// Only HTTP can access it.
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}

	fmt.Println(c)
}
