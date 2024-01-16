package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {

	/*
		The if statement is used to verify conditional
		and follow to into the conditional in case of true
	*/
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
