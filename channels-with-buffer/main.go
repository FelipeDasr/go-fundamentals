package main

import "fmt"

func main() {
	var channel = make(chan string, 2)

	channel <- "Hello"
	channel <- "World"

	fmt.Println(<-channel);
	fmt.Println(<-channel);
}