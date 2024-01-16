package main

import (
	"fmt"
	"math/rand"
) /*for to execute this example make in terminal:
go run .\packages\1 - package.go
*/

/*This is a other way to import packages
in this example we works whith fmt package and math/randa other way is:
import "fmt"
import "math"
*/

//rand is a packcage by math package

func main() {
	//here we used the two packages
	fmt.Println("My favorite number is ", rand.Intn(10))
}
