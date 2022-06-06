// # Serve the files in the "starting-files" folder

// Use "http.FileServer"
package main

import (
	"log"
	"net/http"
)

func main() {
	// The initial dot in the directory's string describes that the files are
	// in the current directory, inside of "starting-files".
	// The usage of "/starting-files/" will result in the application rendering
	// the webpage in "/starting-files/" URI instead of home.
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./starting-files/"))))
}
