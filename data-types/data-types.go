package main

import (
	"errors"
	"fmt"
)

func main() {
	// Integer data types
	var int64Variable int64 = 10000000000000
	inferredIntVariable := 20000000000000

	fmt.Println(int64Variable, inferredIntVariable)

	// Floating-point data types
	var float64Variable float64 = 1.23456789
	inferredFloatVariable := 9.87654321

	fmt.Println(float64Variable, inferredFloatVariable)

	// Boolean data types
	var boolVariable bool = true
	inferredBoolVariable := false

	fmt.Println(boolVariable, inferredBoolVariable)

	// String data types
	var stringVariable string = "String variable"
	inferredStringVariable := "Inferred string variable"

	fmt.Println(stringVariable, inferredStringVariable)

	char := 'a'
	fmt.Sprintln("Char ASCII value: ", char)

	// Errors
	var nilVariable error
	fmt.Println("null error variable ", nilVariable)

	customError:= errors.New("This is a custom error")
	fmt.Println(customError)
}