package main

import (
	"net"
)

type client struct {
	Name       string
	Connection net.Conn
	Commands   chan<- CMD_ID
}
