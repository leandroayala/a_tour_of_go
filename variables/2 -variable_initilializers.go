package main

/*for to execute this example make in terminal:
go run .\variables\2 - variables_initializers.go
*/

import "fmt"

var i, j int = 1, 2 // Declared variables with initializers

func main() {
	var c, python, java = true, false, "no!" //If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	fmt.Println(c, python, java)
}
