package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 1. Create Server listener, using TCPv4 on localhost, port 5555
	listener, err := net.Listen("tcp4", "localhost:5555")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Listening for connections on port 5555...")
	// 2. Accept incoming client connection request
	connection, err := listener.Accept()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Client connected.")
	// Define buffered reader to read byte slice containing message from connection
	bufferedReader := bufio.NewReader(connection)
	// Infinite loop, read up to and including newline, then write response
	for {
		messageBytes, err := bufferedReader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Message received: %s", messageBytes)
		connection.Write(messageBytes)
		if string(messageBytes) == "quit\n" {
			connection.Close()
			listener.Close()
			os.Exit(0)
		}
	}
}
