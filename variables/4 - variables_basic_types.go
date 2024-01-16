package main

import "fmt"

/*
The basic types are bool, int string, byte
*/

func main() {

	/*  variables can be declared into blocls like import*/
	var (
		toBe   bool   = false
		toNum  int    = 0
		toName string = "name"
	)

	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", toNum, toNum)
	fmt.Printf("Type: %T Value: %v\n", toName, toName)
}
