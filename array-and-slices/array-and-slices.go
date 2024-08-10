package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Array and Slices")

	var slice []string; // Dynamic size
	slice = append(slice, "Felipe Dos Anjos")
	slice = append(slice, "João Dos Anjos")

	fmt.Println("Dynamic size (slice): ", strings.Join(slice, ", "))

	var array0 = [...]int{1, 2, 3, 4, 5} // Fixed size
	fmt.Println("\nFixed size ([...]array): ", array0)

	var array1 [5]string; // Fixed size
	array1 = [5]string{"Felipe Dos Anjos", "Neusa Dos Anjos", "Tiana Dos Anjos", "Gaby Dos Anjos", "Davi Dos Anjos"}

	fmt.Println("\nFixed size (array): ", array1)
}