package main

func cerinta1(input []string) []string {

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
