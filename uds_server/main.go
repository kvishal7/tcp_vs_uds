package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const msg = "hello from uds server"

var s = "/Users/test/testsockets/test.sock"

func main() {

	if _, err := os.Stat(s); err == nil {
		os.Remove(s)
	}
	// Create a Unix domain socket and listen for incoming connections.
	socket, err := net.Listen("unix", s)
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup the sockfile.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(s)
		fmt.Printf("\nExiting...\n")
		os.Exit(1)
	}()

	for {
		// Accept an incoming connection.
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a separate goroutine.
		// Create a buffer for incoming data.
		buf := make([]byte, 4096)

		// Read data from the connection.
		_, err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(string(buf))

		// Echo the data back to the connection.
		// _, err = conn.Write([]byte(msg))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		conn.Close()
	}
}
