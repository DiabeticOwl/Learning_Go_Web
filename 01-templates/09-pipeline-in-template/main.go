// Program that will render an html page from a parsed template.
// This template will receive multiple functions from a template.FuncMap.
// This template will use the pipeline character for consecutive execution
// of the mapped functions.
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
// "tu" that will receive a string and return itself in upper case.
// "ff", "sq", "dTwo" and "aFive" will be defined later.
var fm = template.FuncMap{
	"tu":    strings.ToUpper,
	"float": convFloat,
	"ff":    firstFour,
	"sq":    square,
	"dTwo":  divByTwo,
	"aFive": addFive,
}

// firstFour will receive a string and return the first four characters of it.
func firstFour(s string) string {
	return strings.TrimSpace(s)[:4]
}

// convFloat will receive an integer and return itself converted as a Float32
// type.
func convFloat(i int) float32 {
	return float32(i)
}

// square will receive a float32 type value and return its square value.
func square(f float32) float32 {
	return f * f
}

// divByTwo will receive a float32 type value and return itself divided by 2.
func divByTwo(f float32) float32 {
	return f / 2
}

// addFive will receive a float32 type value and return itself added by 5.
func addFive(f float32) float32 {
	return f + 5
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
