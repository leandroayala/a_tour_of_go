package main

import "fmt"

func main() {

	color := "red"

	/*
		the else statement are executed whand the ir condition is false
	*/
	if color == "green" {
		fmt.Println("color is green")
	} else {
		fmt.Println("color is " + color)
	}

}
