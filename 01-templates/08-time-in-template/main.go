// Program that will render an html page from a parsed template.
// This template will receive the current date of execution as a
// value to display.
package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template
var fm = template.FuncMap{
	"tFormat": europeanTimeFormat,
}

// europeanTimeFormat receives a time type value and returns it with the
// date and time notation in Europe.
func europeanTimeFormat(t time.Time) string {
	// Based on the specific time stamp in Go (01/02 03:04:05PM '06 -0700)
	// 01 Meaning month
	// 02 Meaning day
	// 03 Meaning hours
	// 04 Meaning minutes
	// 05 Meaning seconds
	// '06 Meaning year
	// 0700 Standing for GMT-XXXX
	return t.Format("02/01/2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	// File in which the rendered template will be.
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
