package main

import "fmt"

/*
In Go we has one looping cosntruct, the for loop

*/
func main() {

	sum := 0

	//the for has the strcture:
	//Initial statement ( i:= 0); the condition expression (i <10); post statement (i++)
	//he init statement: executed before the first iteration
	//the post statement: executed at the end of every iteration
	for i := 0; i < 10; i++ {
		sum += 1
	}

	//The lopp will finish when the conditional evaluates to false

	fmt.Println(sum)

	//the initial and post statements are optinals
	cont := 1
	for cont < 1000 {
		cont += cont
	}
}
