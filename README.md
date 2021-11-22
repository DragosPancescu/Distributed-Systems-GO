# Welcome to Client-Server Application for the Distributed Systems Laboratory

## I recommend using an utility program like NetCat for connecting to the server after running the command:
The server starts on port 8080 by default, but it can be configured from [Data\config.json](https://github.com/DragosPancescu/SD-Tema1/blob/main/Data/config.json)  
- `go run main/main.go`  

## Commands

- `/help` - shows a list of commands the user can issue to the server.
- `/command1 [args]` - returns a new list comprised of elements that on position **i** (0 <= i <= N) have the characters coresponding to the characters on position **i** from every starting element (args should be a list of strings). 
- `/command2 [args]` - returns the number of integers that are perfect squares, found in the arguments list (arg element should look like *ab1cd5i*)
- `/command3 [args]` - returns the sum of the arguments list after their digits are reversed (args should be a list of integers).
- `/command4 [args]` - returns the arithmetic mean of the elements that have the digit sum between the first 2 arguments (args should be a list of integers)
- `/command5 [args]` - TODO
- `/exit` - exits the server
