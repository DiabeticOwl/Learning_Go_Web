// Program that sets up a DefaultServeMux at port 8080 and serves an
// image to it using "FileServer".
package main

import (
	"io"
	"net/http"
)

func main() {
	// Using the Handler type that returns FileServer we can set an entire
	// directory to be tracked down by the application's server.
	http.Handle("/", http.FileServer(http.Dir(".")))
	// The image that will be copied will be set in this address.
	http.HandleFunc("/cat", catW)

	http.ListenAndServe(":8080", nil)
}

// catW makes makes the response able to render HTML tags and tells the
// page to look for the image served at a specific address.
func catW(w http.ResponseWriter, r *http.Request) {
	// Sets the "Content-Type" of the header of the response as follows
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<img src="/cat.jpeg" width="600" height="500">
	`)
}
