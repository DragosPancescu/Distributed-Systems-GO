package main

import (
	"fmt"
	"strconv"
)

func main() {

	cfg := read_config("Data\\config.json")

	fmt.Println("Server Started on port:" + strconv.Itoa(cfg.Port))

}
