package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan string)

	go write("Hello77", channel)
	
	for msg := range channel {
		fmt.Println(msg)
	}
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		(channel) <- text // Send text to channel
		time.Sleep(time.Second)
	}

	close(channel)
}
