package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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

func Command_help() string {

	fmt.Println("The server is processing the data...")
	// Opens the help file
	help_file, err := os.Open("Data\\help.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer help_file.Close()

	// Reads the file conntents
	file_contents, _ := ioutil.ReadAll(help_file)

	return string(file_contents)
}

func Parse_message(message string) (string, []string) {
	message = strings.TrimRight(message, "\n")

	message_list := strings.Split(message, " ")

	return message_list[0], message_list[1:]
}

func reverse_number(input string) int {
	// Reversed digits in a number
	reversed := 0
	factor := 1

	for i := 0; i < len(input); i++ {
		digit, _ := strconv.Atoi(string(input[i]))
		reversed += factor * digit
		factor *= 10
	}

	return reversed
}

/*
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
*/

func Reverse_sum(input []string) int {
	output := 0

	for i := 0; i < len(input); i++ {
		output += reverse_number(input[i])
	}

	return output
}
