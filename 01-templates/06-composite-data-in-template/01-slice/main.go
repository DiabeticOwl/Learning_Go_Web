// Program that will render a template with a passed slice of strings.
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
	sliceStr := []string{"Johann", "Cruz", "Hello", "World"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sliceStr)
	if err != nil {
		log.Fatalln(err)
	}
}
