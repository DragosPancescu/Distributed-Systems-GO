package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/DragosPancescu/SD-Tema1/client"
	"github.com/DragosPancescu/SD-Tema1/common"
)

func handle_connection(conn net.Conn) {
	conn.Write([]byte("Welcome to the server!\n"))
	conn.Write([]byte("Please choose a name: "))

	name, _ := bufio.NewReader(conn).ReadString('\n')
	new_client := client.Create_client(name, conn)

	fmt.Println(new_client.Name + " s-a alaturat serverului!")

	for {
		message, err := bufio.NewReader(new_client.Connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			new_client.Connection.Close()
			return
		}

		// Sanitize the string
		message = strings.TrimRight(message, "\n")

		// Handle help message
		if message == "/help" {
			fmt.Println(new_client.Name + " sent a help request")
			help_string := common.Command_help()
			fmt.Println("Server is sending a response to " + new_client.Name)
			new_client.Connection.Write([]byte(help_string))
		}
	}
}

func main() {

	cfg := common.Read_config("Data\\config.json")

	server, err := net.Listen(cfg.Conn_type, cfg.Host+": "+strconv.Itoa(cfg.Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	fmt.Println("Server Started on port:" + strconv.Itoa(cfg.Port))
	defer server.Close()

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handle_connection(connection)
	}
}
