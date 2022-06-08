package main

import (
	"fmt"
	"net/http"
	"time"
)

// getUser returns the user found in the session.
// Returns a zero value of type user if none found.
func getUser(w http.ResponseWriter, r *http.Request) user {
	var u user

	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	if sess, ok := dbSessions[c.Value]; ok {
		u = dbUsers[sess.ui]

		restartLength(w, c)
	}

	return u
}

// alreadyLoggedIn returns a bool indicating whether the user is already
// in session or not.
func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	_, ok := dbUsers[dbSessions[c.Value].ui]

	if ok {
		restartLength(w, c)
	}

	return ok
}

func restartLength(w http.ResponseWriter, c *http.Cookie) {
	sess := dbSessions[c.Value]
	sess.lastActivity = time.Now()

	dbSessions[c.Value] = sess

	c.MaxAge = sessionLength
	c.Path = "/"

	http.SetCookie(w, c)
}

func cleanSessions() {
	i := 0

	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Duration(sessionLength) * time.Second) {
			fmt.Println(k, v.lastActivity)
			delete(dbSessions, k)

			i++
		}
	}

	if i > 0 {
		fmt.Printf("################ %v Sessions cleaned ################\n\n", i)
	}
}

// for demonstration purposes
func showSessions() {
	fmt.Println("**************** Currently alive sessions ****************")
	for k, v := range dbSessions {
		fmt.Printf("%v - %v- %v\n", k, v.lastActivity, v.ui)
	}
	fmt.Println()
}
