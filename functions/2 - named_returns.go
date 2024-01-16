package main

import "fmt"

/*
in thus example we used named values in return
note that the return is empty
A return statement without arguments returns the named return values. This is known as a "naked" return.

*/
//Here we have a named arguments x,y also called naked
func split(sum int) (x, y int) {

	x = sum * 4 / 9
	y = sum - x

	return //nake return -> will return x,y values
}

func main() {

	fmt.Println(split(17)) //print the result of split func
}
