// Program that enables a "tcp" server in the port "8080".
// When the server receives a connection it prints some messages.
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

		// Launches handle concurrently for each connection.
		go handle(conn)
	}
}

// handle will take a connection and scan the value received.
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	// Scans each line of what the connection passes/writes.
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// Currently the connection is still open so any code here won't run.
	fmt.Println("I won't run. :(")
}
