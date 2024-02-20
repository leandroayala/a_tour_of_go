package main

/*
to execute this example:  go run .\structs\struct.go
*/

import "fmt"

/*
A struct is a collection of fields
*/
type Vertex struct {
	X int
	Y int
}

func main() {

	//Using the struct
	vertex := Vertex{1, 2}
	fmt.Println(vertex)

	//to Access fields of structs to use a dot
	fmt.Println(vertex.X)

}
