package main

func main() {
	number := -5

	if number > 0 {
		println("Number is positive")
	} else {
		println("Number is negative")
	}

	// if init
	if otherNumber := 5; otherNumber > 0 {
		println("Other number is positive")
	} else {
		println("Other number is negative")
	}
}