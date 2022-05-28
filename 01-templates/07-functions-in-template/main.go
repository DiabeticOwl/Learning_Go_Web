// Program that will render an html page from a parsed template.
// This template will receive two functions from a template.FuncMap.
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name    string
	Surname string
	Age     int
}

// fm contains two functions,
// "tu" that will receive a string and return itself in upper case and
// "ff", defined later.
var fm = template.FuncMap{
	"tu": strings.ToUpper,
	"ff": firstFour,
}

// firstFour will receive a string and return the first four characters of it.
func firstFour(s string) string {
	return strings.TrimSpace(s)[:4]
}

func init() {
	// The template is initialized first with the "New" method, then assigned
	// the desired function map and finally parsing the template file.
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
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
