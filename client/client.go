package client

import (
	"net"
)

type Client struct {
	Name       string
	Connection net.Conn
	Commands   chan<- int
}

/*
	//Connect to server
	server_connection, _ := net.Dial("tcp", "127.0.0.1:8080")

	for {
		// Construim un mesaj (comanda)
		// -----

		// fmt.Printf(server_connection, mesaj + '\n')
	}
*/
