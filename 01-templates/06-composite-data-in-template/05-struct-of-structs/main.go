// Program that will render an html page from a parsed template
// with a struct of structs.
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

type phone struct {
	Brand string
	Model string
}

type items struct {
	People     []person
	CellPhones []phone
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

	// ------------ People ------------
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
	// ------------ People ------------

	// ------------ Phones ------------
	phoneOne := phone{
		Brand: "Samsung",
		Model: "S22 Ultra 5G",
	}

	phoneTwo := phone{
		Brand: "Apple",
		Model: "iPhone 13 Pro Max",
	}
	// ------------ Phones ------------

	people := []person{personOne, personTwo, personThree, personFour}
	cellPhones := []phone{phoneOne, phoneTwo}

	data := items{
		People:     people,
		CellPhones: cellPhones,
	}

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
