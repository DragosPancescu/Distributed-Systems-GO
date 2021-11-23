package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// Read command line arguments and handle errors
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Please provide host and port.")
		return
	} else if len(args) > 3 {
		fmt.Println("Too many arguments (please provide only host and port)")
		return
	}

	// Connect to the server
	conn, err := net.Dial("tcp", args[1]+":"+args[2])
	if err != nil {
		fmt.Println("Can't connect to the server: " + err.Error())
		return
	}

	// Get initial messages the server might send to the client
	// TODO: Make this work

	/*
		reader := bufio.NewReader(conn)
		for {
			message, _ := reader.ReadString('\n')
			fmt.Print(message)

			// Check if there is more data to read
			data, _ := reader.Peek(1)
			if data != nil {
				break
			}
		}*/

	reader := bufio.NewReader(conn)

	message, _ := reader.ReadString('\n')
	fmt.Print(message)

	message, _ = reader.ReadString('\n')
	fmt.Print(message)

	for {
		// Waiting for input
		fmt.Print(">> ")

		// Read message from the client
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			conn.Close()
			return
		}

		// Sanitize the message for server
		message = strings.TrimLeft(message, "> ")

		// Send it to the server
		writer := bufio.NewWriter(conn)
		writer.WriteString(message)
		writer.Flush()

		delim := '\n'
		if strings.Contains(message, "/help") {
			delim = '~'
		}

		// Get the response from the server
		response, err := bufio.NewReader(conn).ReadString(byte(delim))
		if err != nil {
			fmt.Println("Could not read response from server: " + err.Error())
			return
		}

		// Remove the carriage return
		response = strings.Trim(response, "\r")
		if delim == '~' {
			response = response[1:]
		}
		fmt.Print(response)

		if strings.Contains(response, "See you soon") {
			return
		}
	}
}
