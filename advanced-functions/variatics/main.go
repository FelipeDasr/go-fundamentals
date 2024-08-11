package main

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	println(sum(1, 2, 3, 4, 5))
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}