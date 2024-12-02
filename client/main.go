package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 1. Attempt to dial and connect to server
	connection, err := net.Dial("tcp4", "127.0.0.1:5555")
	if err != nil {
		log.Fatal(err.Error())
	}
	// Define buffered readers to read user input and server response
	inputReader := bufio.NewReader(os.Stdin)
	responseReader := bufio.NewReader(connection)
	// Infinite loop, prompt user to write message then read and print response
	for {
		fmt.Print("Enter Message ('quit' to disconnect) > ")
		messageBytes, err := inputReader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		connection.Write(messageBytes)
		responseBytes, err := responseReader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		// Print response as string
		fmt.Print(string(responseBytes))
		if string(responseBytes) == "quit\n" {
			connection.Close()
			os.Exit(0)
		}
	}
}
