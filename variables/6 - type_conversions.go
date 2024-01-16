package main

func main() {

	/*
		we can convert a variable type this way
		T(v) Type e variable
	*/
	var numA int = 42
	var numB float64 = float64(numA) // numB is a float64 of numA
	var numC uint = uint(numB)       // numC is auint

	// other way to declare a variable
	numA = 42
	numB = float64(numC)
	numC = uint(numB)

}
