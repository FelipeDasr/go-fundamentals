package main

import (
	"fmt"
)

func main() {
	fmt.Println("Internal arrays")

	var sliceFromMake = make([]int, 2, 2)

	fmt.Println("Length: ", len(sliceFromMake))
	fmt.Println("Capacity: ", cap(sliceFromMake))

	sliceFromMake = append(sliceFromMake, 1)

	fmt.Println("\n\nLength: ", len(sliceFromMake))
	fmt.Println("Capacity: ", cap(sliceFromMake))
}