package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/DragosPancescu/SD-Tema1/client"
	"github.com/DragosPancescu/SD-Tema1/common"
)

func handle_connection(conn net.Conn, cfg common.Config) {
	// Welcome messages to the client, prompting for a name
	conn.Write([]byte("Welcome to the server!\n"))
	conn.Write([]byte("Please choose a name: "))

	// Handle error
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err.Error())
		conn.Close()
		return
	}

	// Creates a client object
	new_client := client.Create_client(name, conn, common.Get_random_color())

	fmt.Println(common.Color_string((new_client.Name + " connected to the server!"), new_client.Color))
	new_client.Connection.Write([]byte(common.Color_string(("Hi " + new_client.Name + ", please use command /help to see the server capabilitites.\n"), new_client.Color)))

	for {
		// Reads a message sent by a client
		message, err := bufio.NewReader(new_client.Connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			new_client.Connection.Close()
			return
		}

		// Parse the message
		command, args := common.Parse_message(message)

		// Check for max args length
		if len(args) > cfg.Max_args_size {
			fmt.Println(common.Color_string(("Server is sending a response to " + new_client.Name + "."), new_client.Color))
			new_client.Connection.Write([]byte(common.Color_string("Too many arguments for you command (max arguments number is "+strconv.Itoa(cfg.Max_args_size)+").\n", new_client.Color)))
		} else {

			switch command {
			// Handle help command.
			case "/help":
				{
					common.Handle_help_command(new_client, args)
				}
			// Handle exit command.
			case "/exit":
				{
					common.Handle_exit_command(new_client, args)
					return
				}
			// Handle command1 command.
			case "/command1":
				{
					common.Handle_command1(new_client, args)
				}
			// Handle command2 command.
			case "/command2":
				{
					common.Handle_command2(new_client, args)
				}
			// Handle command3 command.
			case "/command3":
				{
					common.Handle_command3(new_client, args)
				}
			// Handle command4 command.
			case "/command4":
				{
					common.Handle_command4(new_client, args)
				}
			// Unknown command, alert the user and prompt for another command.
			default:
				{
					common.Handle_unknown_command(new_client, args)
				}
			}
		}
	}
}

func main() {

	cfg := common.Read_config("Data\\config.json")

	// Starts listening to the designed host:port
	server, err := net.Listen(cfg.Conn_type, cfg.Host+": "+strconv.Itoa(cfg.Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	fmt.Println("Server started on port:" + strconv.Itoa(cfg.Port))
	defer server.Close()

	// Loop for accepting connections from users.
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handle_connection(connection, cfg)
	}
}
