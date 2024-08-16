package main

import "fmt"

func revertString(text *string) {
	var reverted = ""

	for i := len(*text) - 1; i >= 0; i-- {
		reverted += string((*text)[i])
	}

	*text = reverted
}

func main() {
	var text string = "Felipe dos anjos"
	revertString(&text)

	fmt.Println(text)
}