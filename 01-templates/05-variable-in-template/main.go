// Program that will print a html template with data passed.
// Here a variable will be instantiated in the template file.
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Johann Cruz")
	if err != nil {
		log.Fatalln(err)
	}
}
