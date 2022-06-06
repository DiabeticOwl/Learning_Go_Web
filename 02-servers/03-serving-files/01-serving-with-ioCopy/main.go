// Program that sets up a DefaultServeMux at port 8080 and serves an
// image to it.
package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", catW)
	// The image that will be copied will be set in this address.
	http.HandleFunc("/cat.jpeg", catPic)

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

// catPic will copy the local image to the response.
func catPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("cat.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
