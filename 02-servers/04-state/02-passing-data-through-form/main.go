// Program that presents to the user a simple form. This form allows the user
// to upload a file, which will be opened, shown to the user and saved in the
// local server.
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	// Initializes the variable that will contain the file as a string.
	var s string
	fmt.Println(r.Method)

	if r.Method == http.MethodPost {
		// Extracts the file that is attached to the form in the request
		// with name "f".
		f, h, err := r.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Prints information about the file.
		fmt.Printf("\nfile: %v\nheader: %v\nerr: %v", f, h, err)

		// Reads the file's bytes.
		sb, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the file as a string.
		s = string(sb)

		// Instantiates a pointer to another file in which the already read
		// file will be written. This variable implements the writer interface.
		dst, err := os.Create(filepath.Join("./files/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Writes the file in the given path.
		_, err = dst.Write(sb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// There is no `return` as this is the end of the block regarding
			// the reading and writing of the file.
		}
	}

	// Executes the template with the value of s.
	// Either zero value or the string version of the file uploaded.
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}
