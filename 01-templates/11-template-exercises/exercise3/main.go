// Program that will render a template from a CSV file.
package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

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

	// Opening CSV file.
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Reading CSV file. The returned value should be a slice of
	// slice of string type.
	csvR, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", csvR)
	if err != nil {
		log.Fatalln(err)
	}
}
