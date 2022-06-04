// # Building upon the code from the previous problem:
// Add code to respond to the following METHODS & ROUTES:
// * GET /
// * GET /apply
// * POST /apply
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

			if data.Method == "GET" && data.URI == "/" {
				responseHome(conn)
			} else if data.Method == "GET" && data.URI == "/apply" {
				responseApply(conn)
			} else if data.Method == "POST" && data.URI == "/apply" {
				responseApplyPost(conn)
			}
		}

		if ln == "" {
			break
		}

		i++
	}

	fmt.Println("Code got here.")
}

func responseHome(conn net.Conn) {
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

			<ul>
				<li><a href="/apply">Apply Here</a></li>
			</ul>
		</body>
	</html>
	`

	tpl, err := template.New("MyTemplate").Parse(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.Execute(conn, data)
}

func responseApply(conn net.Conn) {
	body := `
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Apply Here</title>
		</head>
		<body>
			<h1>Apply here! :D</h1>
			{{ if . }}
				<h1>Method: {{ .Method }}</h1>
				<h1>URI: "{{ .URI }}"</h1>
			{{ end }}

			<ul>
				<li><a href="/">Home</a></li>
			</ul>

			<form action="/apply" method="post">
				<input type="submit" value="Press me to apply">
			</form>
		</body>
	</html>
	`

	tpl, err := template.New("MyTemplate").Parse(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.Execute(conn, data)
}

func responseApplyPost(conn net.Conn) {
	body := `
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Apply Here</title>
		</head>
		<body>
			<h1>You have applied! :D</h1>
			{{ if . }}
				<h1>Method: {{ .Method }}</h1>
				<h1>URI: "{{ .URI }}"</h1>
			{{ end }}

			<ul>
				<li><a href="/">Home</a></li>
				<li><a href="/apply">Apply Again Here</a></li>
			</ul>
		</body>
	</html>
	`

	tpl, err := template.New("MyTemplate").Parse(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.Execute(conn, data)
}
