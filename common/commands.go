package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

// Command 1 from the homework
func Command1(input []string) []string {

	var output []string

	for i := 0; i < len(input[0]); i++ {
		var aux_string string
		for j := 0; j < len(input); j++ {
			aux_string += string(input[j][i])
		}
		output = append(output, aux_string)
	}

	return output
}

func reverse_number(input string) (int, bool) {

	// Checkf if number is negative
	is_negative := false
	if input[0] == '-' {
		is_negative = true
		input = input[1:]
	}

	// Reverse digits in a number
	input_int, err := strconv.Atoi(input)
	if err != nil {
		return -1, false
	}

	reversed := 0
	factor := 1

	for i := 0; i < len(input); i++ {
		digit := input_int % 10
		input_int /= 10

		reversed = (reversed * factor) + digit
		factor *= 10
		fmt.Println(strconv.Itoa(reversed))
	}

	if is_negative {
		reversed *= -1
	}

	return reversed, true
}

// Command 3 from the homework
func Reverse_sum(input []string) (int, bool) {
	output := 0

	for i := 0; i < len(input); i++ {
		reversed_number, good_format := reverse_number(input[i])

		if good_format {
			output += reversed_number
		} else {
			return -1, false
		}
	}

	return output, true
}
