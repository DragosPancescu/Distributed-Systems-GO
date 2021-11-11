package main

type CMD_ID int

const (
	CMD_1 CMD_ID = iota
	CMD_2
	CMD_3
)

type command struct {
	Id     CMD_ID
	Client *client
	Args   []string
}

func command1(input []string) []string {

	var output []string

	for i := 1; i < len(input); i++ {
		var aux_string string
		for j := 1; j < len(input); j++ {
			aux_string += string(input[j][i])
		}
		output[i] = aux_string
	}

	return output
}
