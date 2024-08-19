package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			channel1 <- "Every second - Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Every two seconds - Channel 2"
		}
	}()

	for {
		select {
		case channel1Msg := <-channel1:
			fmt.Println(channel1Msg)
		case channel2Msg := <-channel2:
			fmt.Println(channel2Msg)
		}
	}
}