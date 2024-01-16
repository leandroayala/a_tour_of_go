package main

import "fmt"

/*
When declaring a variable without specifying an explicit type
(either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
*/

func main() {

	v := 42 // if you changethis value for  42.5 or "my name" the type will be infered to variable
	fmt.Printf("v is of type %T\n", v)

}
