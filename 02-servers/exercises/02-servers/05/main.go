// # Building upon the code from the previous exercise:
// ## In that previous exercise, we WROTE to the connection.
// * Now I want you to READ from the connection.
// * You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.
// * Use bufio.NewScanner() to read from the connection.
// * After all of the reading, include these lines of code:
// * `fmt.Println("Code got here.")`

// ## Launch your TCP server.
// * In your **web browser,** visit localhost:8080.
// * Now go back and look at your terminal.

// ## Add code to WRITE to the connection.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
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
		</body>
		</html>
	`
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
	fmt.Fprint(conn, body)
}
