package main

import "fmt"

func main() {

	/*
		when a variable is declared,without value atribuition it assumes the "Zero" value
	*/

	//to string the zero value is empty string -> ""
	var zero string
	//to numeric types the zero value is 0
	var i int
	var x float64
	//to bool types the zero value is false
	var ok bool

	fmt.Printf("%q %v %v %v\n", zero, i, x, ok)

}
