package main

import "fmt"

func toggleSign(x *int) {
	*x *= -1
}

func main() {
	number1 := 10

	toggleSign(&number1)
	fmt.Println(number1)

	toggleSign(&number1)
	fmt.Println(number1)
}