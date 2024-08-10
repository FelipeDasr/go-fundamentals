package main

import (
	"fmt"
	"module/utils"
)

func main() {
	fmt.Println("Writing from the main package.")
	fmt.Println(utils.SumToNumbers(1, 2))
	fmt.Println(utils.SumFourNumbers(1, 2, 3, 4))
}