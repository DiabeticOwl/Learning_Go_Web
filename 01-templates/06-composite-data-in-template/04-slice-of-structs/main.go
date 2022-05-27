// Program that will render an html page from a parsed template.
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
	// File in which the rendered template will be.
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	personOne := person{
		Name:    "Johann",
		Surname: "Cruz",
		Age:     165,
	}

	personTwo := person{
		Name:    "Fulano",
		Surname: "DeTal",
		Age:     85,
	}

	personThree := person{
		Name:    "John",
		Surname: "Snow",
		Age:     30,
	}

	personFour := person{
		Name:    "David",
		Surname: "Tennant",
		Age:     51,
	}

	people := []person{personOne, personTwo, personThree, personFour}

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
}
