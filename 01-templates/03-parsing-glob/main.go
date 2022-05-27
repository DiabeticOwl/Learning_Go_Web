// Program that reads a set of gohtml files and parse them
// using ParseGlob and ExecuteTemplate.
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// The variable "tpl" is initialized as a pointer to a template.Template type.
var tpl *template.Template

// init function will run once before the function main.
// Used in this case to assign the templates we desire.
func init() {
	// template.Must is a method that checks if there is any error while
	// parsing the definition of the files passed.
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	// "ExecuteTemplate" is used when the template container has more than one.
	// This is in order to specify the specific template to execute.
	for i := 1; i <= 4; i++ {
		err := tpl.ExecuteTemplate(os.Stdout, fmt.Sprintf("file_%v.gohtml", i), nil)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
