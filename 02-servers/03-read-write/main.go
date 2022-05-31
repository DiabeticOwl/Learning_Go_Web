// Program that enables a "tcp" server in the port "8080".
// When the server receives a connection it prints some messages.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

		fmt.Fprintln(conn, "\n***** You have 10 seconds to write",
			"something before you are disconnected. ****")

		// Launches handle concurrently for each connection.
		go handle(conn)
	}
}

// handle will take a connection and scan the value received.
// It also sets a dead line for the connection, meaning that
// the connection will be timed out at provided time.
func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(conn)

	// Scans each line of what the connection passes/writes.
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		// Prints into the connection a formatted string with the value
		// written.
		fmt.Fprintf(conn, "Did you just said: '%v'?\n", ln)
	}
	defer conn.Close()

	// Now this line will run after the connection is broken,
	// either by the user or by the deadline.
	fmt.Println("I will run! :)")
}
