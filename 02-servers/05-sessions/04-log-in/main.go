// Program that allows the user to sign up, log in and log out by saving their
// info and assign it an UUID.
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

	sb, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	dbUsers["mail@mail.com"] = user{
		UserName:  "mail@mail.com",
		Password:  sb,
		FirstName: "Fulano",
		LastName:  "Fulene",
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup/", signUp)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/logout/", logout)
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

// login will execute the given template and receive the information passed
// from the form displayed and check if it is in our database of users.
// If it is the function will create a session and assign it to that user.
// If it is not then the user will be redirected to "/".
func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		ui := r.FormValue("username")
		pw := r.FormValue("password")

		u, ok := dbUsers[ui]
		if !ok {
			http.Error(w,
				"Incorrect Username or Password.",
				http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pw))
		if err != nil {
			http.Error(w,
				"Incorrect Username or Password.",
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

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

// logout will check that the user is logged in and remove their session and
// record from "dbSessions". If the user isn't logged in or the cookie isn't
// found it will redirect to "/".
func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// err is thrown because it is already checked in "alreadyLoggedIn".
	c, _ := r.Cookie("session")

	c.MaxAge = -1
	c.Path = "/"
	http.SetCookie(w, c)

	delete(dbSessions, c.Value)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
