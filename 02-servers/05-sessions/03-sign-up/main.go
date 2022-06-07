// Program that allows the user to sign up by saving their info and assign it
// an UUID.
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	Password  []byte
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
	http.HandleFunc("/signup/", signUp)
	http.HandleFunc("/library/", library)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

// index will execute the given template with the user gotten with "getUser".
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", getUser(r))
}

// library will execute the given template with the user gotten with "getUser"
// if the result of "alreadyLoggedIn" is true.
func library(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "library.gohtml", getUser(r))
}

// signUp will execute the given template and save the data that the user
// posted through the form displayed, encrypting some of it and associating
// it with an UUID. signUp will redirect to "/" if "alreadyLoggedIn"
// returns true.
func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		ui := r.FormValue("username")
		pw := r.FormValue("password")
		fn := r.FormValue("firstname")
		ln := r.FormValue("lastname")

		if _, ok := dbUsers[ui]; ok {
			http.Error(w,
				"The submitted username is already in use.",
				http.StatusForbidden)

			return
		}

		sID := uuid.New()
		c := &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			Path:     "/",
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = ui

		// Encrypting password with bcrypt.
		sb, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w,
				"Internal Server Error",
				http.StatusInternalServerError)

			return
		}

		u := user{
			UserName:  ui,
			Password:  sb,
			FirstName: fn,
			LastName:  ln,
		}
		dbUsers[ui] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
