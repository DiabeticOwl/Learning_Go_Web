// Program that sets up a DefaultServeMux at port 8080 and serves an
// image to it using "FileServer" and "StripPrefix".
package main

import (
	"io"
	"net/http"
)

func main() {
	// Now the home page and every get method will be linked to catW function.
	http.HandleFunc("/", catW)

	// Using the Handler type that returns FileServer we can set an entire
	// directory to be tracked down by the application's server.
	// Adding the StripPrefix function the handler passed will link the
	// directory to a "Prefix". A prefix is a case-sensitive name that will
	// be used throughout the application as a nickname of the directory.
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./resources"))))

	http.ListenAndServe(":8080", nil)
}

// catW makes makes the response able to render HTML tags and tells the
// page to look for the image served at a specific address.
func catW(w http.ResponseWriter, r *http.Request) {
	// Sets the "Content-Type" of the header of the response as follows
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<img src="/assets/cat.jpeg" width="600" height="500">
	`)
}
