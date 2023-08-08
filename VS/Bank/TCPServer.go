package main

import (
	"fmt"
	"log"
	"net"
)

func StartTCPServer() {
	// Define the HTML template
	listener, err := net.Listen("tcp", ":6543")
	if err != nil {
		log.Fatal("cannt connect to Port :6543 ", err)
	}

	defer CloseListener(listener) //am Ende schließe die Connection

	fmt.Println("Server started. Listening on port 6543...")

	for {
		// Verbindung von einem Client akzeptieren
		conn, err1 := listener.Accept()
		if err1 != nil {
			panic(err)
		}
		AnzahlTCPrequest++
		go handleRequest(conn)
	}
}
func CloseListener(listener net.Listener) {
	err := listener.Close()
	if err != nil {
		return
	}
}
