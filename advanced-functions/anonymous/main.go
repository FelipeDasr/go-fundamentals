package main

func main() {
	(func() {
		println("Hello, world!")
	})()

	functionValue := func(a, b int) int {
		return a + b
	}(3, 7)

	println(functionValue)
}