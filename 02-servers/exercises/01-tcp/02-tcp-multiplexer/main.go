// Program that will work as a "multiplexer" that depending on the URI that
// the connection describes.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

// Available pages for the response function.
var pages map[string]string = map[string]string{
	"/":        "Home",
	"/home":    "Home",
	"/about":   "About Us",
	"/contact": "Contact Us",
}

func main() {
	// Uses net package to initiate the "tcp" server.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()

	// This loop will run while the server is running (in this
	// case until CTR+C is press on terminal).
	for {
		// The server will wait until a connection comes.
		// "Accept" method will return the connection.
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// Launches request concurrently for each connection.
		go request(conn)
	}
}

// request will take the connection passed and scan it's content, sending
// a response to it accordingly.
func request(conn net.Conn) {
	defer conn.Close()

	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			response(conn, ln)
		}

		// Loop is broken since the scanner is done.
		if ln == "" {
			break
		}

		i++
	}
}

// response will extract the method and URI from the header string passed
// and will write the correspondent template to the connection.
func response(conn net.Conn, header string) {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	m := strings.Fields(header)[0]
	u := strings.Fields(header)[1]

	fmt.Printf("The method is: %v\n", m)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")

	err := tpl.ExecuteTemplate(conn, "tpl.gohtml", pages[u])
	if err != nil {
		log.Fatalln(err)
	}
}
