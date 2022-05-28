package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

// This struct can be optimized by making the Fall, Spring and Summer
// attributes a slice of semester type.
type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

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

	years := []year{
		{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"CSCI-40", "Introduction to Programming in Go", "4"},
					{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"CSCI-50", "Advanced Go", "5"},
					{"CSCI-190", "Advanced Web Programming with Go", "5"},
					{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"CSCI-40", "Introduction to Programming in Go", "4"},
					{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"CSCI-50", "Advanced Go", "5"},
					{"CSCI-190", "Advanced Web Programming with Go", "5"},
					{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
			Summer: semester{
				Term: "Summer",
				Courses: []course{
					{"CSCI-300", "Summer Go Internship", "0"},
					{"CSCI-301", "Summer Go Web Internship", "0"},
				},
			},
		},
	}

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", years)
	if err != nil {
		log.Fatalln(err)
	}
}
