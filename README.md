# Welcome to Client-Server Application for the Distributed Systems Laboratory

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)

<img src="https://iconape.com/wp-content/png_logo_vector/go-gopher-favicon.png" width="200" height="200" />

## To start the server run the following command:

- `go run main/main.go`

The server starts on port 8080 by default, but it can be configured from [Data\config.json](https://github.com/DragosPancescu/SD-Tema1/blob/main/Data/config.json)  

## To connect to the server you can either use an utility program like NetCat or the built-in module *connection*:

The built-in way is done by running the following command if the configured host is 8080.
- `go run Connection/main/main.go localhost 8080`

Or if you have changed the port number:

- `go run Connection/main/main.go localhost [port_number]`

## Commands

1. `/help` - shows a list of commands the user can issue to the server.
2. `/command1 [args]` - returns a new list comprised of elements that on position **i** (0 <= i <= N) have the characters coresponding to the characters on position **i** from every starting element (args should be a list of strings). 
3. `/command2 [args]` - returns the number of integers that are perfect squares, found in the arguments list (arg element should look like *ab1cd5i*)
4. `/command3 [args]` - returns the sum of the arguments list after their digits are reversed (args should be a list of integers).
5. `/command4 [args]` - returns the arithmetic mean of the elements that have the digit sum between the first 2 arguments (args should be a list of integers)
6. `/command5 [args]` - returns a list of converted to base 10 arg elements that are in base 2(args should be a list of strings and base2 integers)
7. `/exit` - exits the server
