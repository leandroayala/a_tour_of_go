package main

import "fmt"

/*for to execute this example make in terminal:
go run .\variables\1 - variables.go
*/

/*
Variables are declared with var statement
A var stattement can be at declared package level or function level
*/

var c, python, java bool //package level

func main() {

	var i int // function level

	fmt.Println(i, c, python, java)

}
