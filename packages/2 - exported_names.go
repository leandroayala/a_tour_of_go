package main

/*for to execute this example make in terminal:
go run .\packages\2 - exported_names.go
*/

import (
	"fmt"
	"math"
)

/*

Names are exported if it begins with capital letter
In this example, if you change math.pi, you get a error
because in math package Pi is exported not pi

*/

func main() {
	fmt.Println(math.Pi)
}
