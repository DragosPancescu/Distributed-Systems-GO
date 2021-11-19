package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/DragosPancescu/SD-Tema1/client"
	"github.com/DragosPancescu/SD-Tema1/common"
)

func handle_connection(conn net.Conn) {
	conn.Write([]byte("Welcome to the server!\n"))
	conn.Write([]byte("Please choose a name: "))

	name, _ := bufio.NewReader(conn).ReadString('\n')
	new_client := client.Create_client(name, conn, common.Get_random_color())

	fmt.Println(common.Color_string((new_client.Name + " connected to the server!"), new_client.Color))
	new_client.Connection.Write([]byte(common.Color_string(("Hi " + new_client.Name + ", please use command /help to see the server capabilitites.\n"), new_client.Color)))

	for {
		message, err := bufio.NewReader(new_client.Connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			new_client.Connection.Close()
			return
		}

		// Parse the message
		command, args := common.Parse_message(message)

		switch command {
		// Handle help message
		case "/help":
			{
				fmt.Println(common.Color_string((new_client.Name + " sent a help request."), new_client.Color))
				help_string := common.Command_help()
				fmt.Println(common.Color_string(("Server is sending a response to " + new_client.Name + "."), new_client.Color))
				new_client.Connection.Write([]byte(common.Color_string(help_string, new_client.Color)))
			}
		case "/exit":
			{
				new_client.Connection.Write([]byte(common.Color_string(("See you soon " + new_client.Name + ".\n"), new_client.Color)))
				fmt.Println(common.Color_string((new_client.Name + " is leaving the server."), new_client.Color))
				new_client.Connection.Close()
			}
		case "/reverse_sum":
			{
				fmt.Println(common.Color_string((new_client.Name + " sent a reverse_sum request."), new_client.Color))
				reversed_sum := common.Reverse_sum(args)
				fmt.Println(common.Color_string(("Server is sending a response to " + new_client.Name + "."), new_client.Color))
				new_client.Connection.Write([]byte(common.Color_string("Reversed sum = "+strconv.Itoa(reversed_sum)+"\n", new_client.Color)))
			}
		default:
			{
				fmt.Println(common.Color_string((new_client.Name + " sent an unknown request"), new_client.Color))
				new_client.Connection.Write([]byte(common.Color_string("That request is not valid, please try again.\n", new_client.Color)))
			}
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

	fmt.Println("Server started on port:" + strconv.Itoa(cfg.Port))
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
