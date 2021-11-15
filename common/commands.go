package common

import (
	"github.com/DragosPancescu/SD-Tema1/client"
)

const (
	CMD_1 int = iota
	CMD_2
	CMD_3
)

type Command struct {
	Id     int
	Client client.Client
	Args   []string
}

func Command_handler(id string) {

}

func Command1(input []string) []string {

	var output []string

	for i := 0; i < len(input); i++ {
		var aux_string string
		for j := 0; j < len(input); j++ {
			aux_string += string(input[j][i])
		}
		output[i] = aux_string
	}

	return output
}
