// Program that will render a template with a passed map.
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
	mapStr := map[string]string{
		"Name":           "Johann",
		"Surname":        "Cruz",
		"A String":       "Hello",
		"Another String": "World",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", mapStr)
	if err != nil {
		log.Fatalln(err)
	}
}
