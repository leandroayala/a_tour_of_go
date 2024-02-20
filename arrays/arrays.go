package main

import "fmt"

/*

 */
func main() {

	//basic form declaration of array
	var a [10]int //array 10 position of integer (0 - 10)

	//put on array positions
	a[0] = 1
	a[1] = 2

	fmt.Println(a[0], a[1], a[9]) //a[9] value 0

}
