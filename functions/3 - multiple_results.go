package main

/*for to execute this example make in terminal:
go run .\functions\3 - multiple_results.go
*/

import "fmt"

//This function receives two arguments and returns two values
func concate(name1, name2 string) (string, string) {

	return name1, name2
}

func main() {

	//here we go use the function and his two arguments
	firstname, lastname := concate("Joe", "biden")

	fmt.Println(firstname + " " + lastname)

}
