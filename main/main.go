package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"github.com/DragosPancescu/SD-Tema1/common"
)

func main() {

	cfg := common.Read_config("Data\\config.json")

	server, err := net.Listen(cfg.Conn_type, cfg.Host+": "+strconv.Itoa(cfg.Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Server Started on port:" + strconv.Itoa(cfg.Port))
	defer server.Close()

	connection, _ := server.Accept()

	for {
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Println("Got a message -> " + message)
	}
}
