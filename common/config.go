package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Struct to map the config file to
type Config struct {
	Array_size   int    `json:"array_size"`
	Timeout_time int    `json:"timeout_time"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Conn_type    string `json:"conn_type"`
}

func Read_config(path string) Config {
	// Function that handles the reading of the config file
	// Input -> File path [string]
	// Output -> Config Object [config]

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
