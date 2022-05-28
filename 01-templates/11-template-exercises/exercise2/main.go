// Program that will render a template with a passed slice of structs.
// Named slice will be populated using the faker and math/rand packages.
package main

import (
	"log"
	"math/rand"
	"os"
	"text/template"

	"github.com/jaswdr/faker"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Region  string
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

	// Fixed seed initialized for consistency purpose in the data created.
	var seed int64 = 777

	// Fixed seed assigned to both the rand and faker packages.
	rand.Seed(seed)
	faker := faker.NewWithSeed(rand.NewSource(seed))

	// Fixed regions are initialized as stated in the README.md.
	regions := []string{"Southern", "Central", "Northern"}

	// hotels is an unlimited slice of hotel type.
	var hotels []hotel
	// Fixed number of hotels initialized.
	n := 50

	// Population of hotels.
	for i := 0; i < n; i++ {
		newHotel := hotel{
			Name:    faker.Company().Name(),
			Address: faker.Address().Address(),
			City:    faker.Address().City(),
			Region:  regions[rand.Intn(3)],
		}

		hotels = append(hotels, newHotel)
	}

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
