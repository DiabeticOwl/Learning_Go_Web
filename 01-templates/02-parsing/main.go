// Program that creates an html file and "parse" it through a template.
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// A template is a container with all the parsed templates.
	tpl, err := template.ParseFiles("file.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// "ExecuteTemplate" is used when the template container has more than one.
	// This is in order to specify the specific template to execute.
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
