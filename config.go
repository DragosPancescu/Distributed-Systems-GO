package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Struct to map the config file to
type config struct {
	Array_size              int    `json:"array_size"`
	Timeout_time            int    `json:"timeout_time"`
	Max_concurrent_routines int    `json:"max_concurrent_routines"`
	Host                    string `json:"host"`
	Port                    int    `json:"port"`
	Conn_type               string `json:"conn_type"`
}

func read_config(path string) config {
	// Function that handles the reading of the config file
	// Input -> File path [string]
	// Output -> Config Object [config]

	var jsonOutput config

	// Opens the JSON file
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Reads the file conntents
	fileContents, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(fileContents, &jsonOutput)
	if err != nil {
		fmt.Println(err)
	}

	return jsonOutput
}
