package main

import "fmt"

func main() {
	i := 10

	for i > 0 {
		fmt.Println(i)
		i--
	}

	fmt.Println()

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	names := [3]string{"John", "Doe", "Smith"}

	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}
}