package main

import "fmt"

func sum(n1 int, n2 int) int8 {
	return int8(n1 + n2)
}

func mathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sumResult := sum(10, 20)
	fmt.Println(sumResult)

	functionConstant := func(){
		fmt.Println("This is a function inside a constant cariable")
	}

	sumResult1, subResult2 := mathCalcs(10, 20)
	fmt.Println(sumResult1, subResult2)

	usedValue, _:= mathCalcs(15, 20)
	fmt.Println(usedValue)

	functionConstant()
}