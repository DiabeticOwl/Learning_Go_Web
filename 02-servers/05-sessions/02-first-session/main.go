// Program that saves the session of an user by recording he/she's details
// and associating them with an UUID.
package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var tpl *template.Template

// User ID, user
var dbUsers = make(map[string]user)

// Session ID, User ID
var dbSessions = make(map[string]string)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/library/", library)

	http.ListenAndServe(":8080", nil)
}

// index will check if there is a cookie named "session" and record it
// alongside with the user passed through the template's form. If in
// method GET there is a cookie and a user already in memory it will
// only show the user's details.
func index(w http.ResponseWriter, r *http.Request) {
	// Searching for cookie.
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.New()

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			Path:  "/",
		}

		http.SetCookie(w, c)
	}

	// Instantiating and searching for user.
	var u user
	if ui, ok := dbSessions[c.Value]; ok {
		u = dbUsers[ui]
	}

	// If user's form was submitted.
	if r.Method == http.MethodPost {
		ui := r.FormValue("username")
		fn := r.FormValue("firstname")
		ln := r.FormValue("lastname")

		u = user{
			UserName:  ui,
			FirstName: fn,
			LastName:  ln,
		}

		dbSessions[c.Value] = ui
		dbUsers[ui] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// library will check if there is a cookie named "session" and if there is an
// user associated with it. If there are will execute a given template showing
// the user's details.
func library(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ui, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "library.gohtml", dbUsers[ui])
}
