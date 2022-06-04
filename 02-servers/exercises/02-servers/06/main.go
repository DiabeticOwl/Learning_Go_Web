// # Building upon the code from the previous problem:
// * Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.
// * Add this data to your RESPONSE so that this data is displayed in the browser.
package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

var data struct {
	Method string
	URI    string
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// Read from connection.
	request(conn)

	// Write to connection.
	response(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	i := 0

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			data.Method = strings.Fields(ln)[0]
			data.URI = strings.Fields(ln)[1]
		}

		if ln == "" {
			break
		}

		i++
	}

	fmt.Println("Code got here.")
}

func response(conn net.Conn) {
	body := `
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Greetings</title>
		</head>
		<body>
			<h1>Greetings and Salutations! :D</h1>
			{{ if . }}
				<h1>Method: {{ .Method }}</h1>
				<h1>URI: "{{ .URI }}"</h1>
			{{ end }}
		</body>
	</html>
	`

	tpl, err := template.New("MyTemplate").Parse(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	// Header with the RC7230 HTTP protocol.

	// 	REQUEST LINE
	// GET / HTTP/1.1
	// method SP request-target SP HTTP-version CRLF
	// https://tools.ietf.org/html/rfc7230#section-3.1.1

	// RESPONSE (STATUS) LINE
	// HTTP/1.1 302 Found
	// HTTP-version SP status-code SP reason-phrase CRLF
	// https://tools.ietf.org/html/rfc7230#section-3.1.2

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.Execute(conn, data)
}
