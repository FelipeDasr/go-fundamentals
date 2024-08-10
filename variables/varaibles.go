package main

import "fmt"

func main() {
	var variable1 string = "Variable 1"
	fmt.Println(variable1)

	varaible2 := "variable 2"
	variable3 := "variable 3"

	fmt.Println(varaible2, variable3)

	// Invert the values of two variables
	varaible2, variable3 = variable3, varaible2

	fmt.Println(varaible2, variable3)

	var (
		variable4 string = "variable 4"
		variable5 string = "variable 5"
	)

	fmt.Println(variable4, variable5)

	const constant1 string = "constant 1"

	fmt.Println(constant1)
}