package main

import "fmt"

/*
For is Go while
on Go not exists while statement, but the for works this way
*/

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
