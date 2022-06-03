// Program that enables a "tcp" server in the port "8080".
// When the server receives a connection it prints some messages.
package main

import (
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

		// Other ways of printing in the connection are through the uses of
		// io.WriteString and fmt.Fprintf
		// The following messages will be written in the connection:
		fmt.Fprintln(conn, "\nHello from TCP server")
		fmt.Fprintln(conn, "How is your day?")

		conn.Close()
	}
}
