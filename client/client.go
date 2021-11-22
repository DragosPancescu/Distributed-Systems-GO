package client

// This package contains the client related information

import (
	"net"
	"strings"
)

type Client struct {
	Name       string
	Connection net.Conn
	Color      string
}

func Create_client(name string, conn net.Conn, color string) Client {
	// Remove trailing \n
	name = strings.TrimRight(name, "\n")

	// Create new client
	client := Client{
		Name:       name,
		Connection: conn,
		Color:      color,
	}

	return client
}
