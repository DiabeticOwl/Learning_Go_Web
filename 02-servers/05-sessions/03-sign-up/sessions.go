package main

import (
	"net/http"
)

// getUser returns the user found in the session.
// Returns a zero value of type user if none found.
func getUser(r *http.Request) user {
	var u user

	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	if ui, ok := dbSessions[c.Value]; ok {
		u = dbUsers[ui]
	}

	return u
}

// alreadyLoggedIn returns a bool indicating whether the user is already
// in session or not.
func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	_, ok := dbUsers[dbSessions[c.Value]]
	return ok
}
