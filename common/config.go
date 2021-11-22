package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Struct to hold the config file data
type Config struct {
	Max_args_size int    `json:"max_args_size"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Conn_type     string `json:"conn_type"`
}

// Function that handles the reading of the config file
func Read_config(path string) Config {

	var json_output Config

	// Opens the JSON file
	json_file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	// Reads the file conntents
	file_contents, _ := ioutil.ReadAll(json_file)

	err = json.Unmarshal(file_contents, &json_output)
	if err != nil {
		fmt.Println(err)
	}

	return json_output
}
