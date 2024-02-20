package main

import "fmt"

func main() {

	arr := [4]int{0, 1, 2, 3}

	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	//lets Take a sliceof two last elements
	var newArray []int = arr[2:4] //this select  arra[low : high]

	fmt.Println(newArray)

}
