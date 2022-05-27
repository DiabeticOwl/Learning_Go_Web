// Program that will render a template with a struct value.
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name    string
	Surname string
	Age     int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	aPerson := person{
		Name:    "Johann",
		Surname: "Cruz",
		Age:     165,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", aPerson)
	if err != nil {
		log.Fatalln(err)
	}
}
