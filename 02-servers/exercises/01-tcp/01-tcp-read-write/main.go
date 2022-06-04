// Program that will render an html page from a parsed template and establish
// a TCP listener that will write the parsed template into new connections.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"text/template"
)

var tpl *template.Template

var filepath string = "index.html"

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

	// File in which the rendered template will be.
	nf, err := os.Create(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
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

		// Launches handle concurrently for each connection.
		go handle(conn)
	}
}

// handle takes in a connection and performs two operations to it before
// closing it.
func handle(conn net.Conn) {
	defer conn.Close()

	// Read from connection.
	request(conn)

	// Write to connection.
	response(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		// Loop is broken since the scanner is done.
		if ln == "" {
			break
		}
	}
}

func response(conn net.Conn) {
	body, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}

	// We need to adhere to the RC7230 HTTP protocol.

	// The first line is the status line.
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")

	fmt.Fprint(conn, string(body))
}
