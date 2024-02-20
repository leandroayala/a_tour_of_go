/*
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/
package main

import "fmt"

func main() {

	//no executed immediatly, only next function
	defer fmt.Println("world")

	//Executed immediately and next execute "world"
	fmt.Println("hello")
}
