package utils

// This package contains the helper functions used by the commands issued to the server.

import (
	"math"
	"strconv"
	"unicode"
)

// Checks if string array has elements of the same length
func Check_elem_len(input []string) bool {
	length := len(input[0])

	for i := 1; i < len(input); i++ {
		if length != len(input[i]) {
			return false
		}
	}
	return true
}

// Reverses the digits in a number
// TODO: Refactor because it should check for less than 0 values in another function
func Reverse_number(input string) (int, bool) {

	// Check if number is negative
	is_negative := false
	if input[0] == '-' {
		is_negative = true
		input = input[1:]
	}

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
	}

	if is_negative {
		reversed *= -1
	}

	return reversed, true
}

// Extracts a number from a string (Ex: a1b4c2 -> 142)
func Extract_number(input string) (int, bool) {
	output := 0

	extracted_any := false
	for i := 0; i < len(input); i++ {
		if unicode.IsDigit(rune(input[i])) {
			digit, _ := strconv.Atoi(string(input[i]))

			// First digit should not be 0
			if output == 0 && digit == 0 {
				return -1, false
			}

			output = output*10 + digit
			extracted_any = true
		}
	}

	return output, extracted_any
}

// Checks if a numbers is a perfect square
func Check_perfect_square(input float64) bool {
	square_root := math.Sqrt(input)

	return square_root*square_root == input
}

// Calculates sum of the digits of a number
func Digits_sum(input int) int {
	input_string := strconv.Itoa(input)
	sum := 0

	for i := 0; i < len(input_string); i++ {
		digit, _ := strconv.Atoi(string(input_string[i]))
		sum += digit
	}

	return sum
}

// Check if string is base2 number
func Check_binary(input string) bool {

	for i := 0; i < len(input); i++ {
		if input[i] != '0' && input[i] != '1' {
			return false
		}
	}

	return true
}
