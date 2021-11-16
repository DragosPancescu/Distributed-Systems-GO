package client

import (
	"net"
	"strings"
)

type Client struct {
	Name       string
	Connection net.Conn
}

func Create_client(name string, conn net.Conn) Client {
	// Remove trailing \n
	name = strings.TrimRight(name, "\n")

	// Create new client
	client := Client{
		Name:       name,
		Connection: conn,
	}

	return client
}
