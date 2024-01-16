package main

/*for to execute this example make in terminal:
go run .\functions\1 - functions.go
*/

import (
	"fmt"
)

// a function can take zero ou many arguments

// THis is a strcture of function func namefunction() type_return_function { }
func hello() string {
	return "hello"
}

// Function that receives two params and returns a sum
func sum(n1 int, n2 int) int {

	return n1 + n2

}

// All go programs must have  a main function
func main() {

	//Here we declared a variable
	var msg string
	//here we will call the function hello
	msg = hello()

	//Prin the msg variable that received the hello funtion return
	fmt.Println(msg)

	//call the function sum and print result
	fmt.Println(sum(3, 3))

}
