package main

import "fmt"

func mathCalc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := mathCalc(5, 3)
	fmt.Println(sum, sub)

	fmt.Println(mathCalc(5, 8))
}