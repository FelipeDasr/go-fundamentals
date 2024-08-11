package main

import "fmt"

func closure() func() string {
	text := "Hello, World!"

	return func() string {
		return text
	}
}

func main() {
	text := "Hello, World!"
	fmt.Println(text)

	fmt.Println(closure()())
}