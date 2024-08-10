package main

import "fmt"

type Person struct {
	name     string
	age      uint
	document string
}

type Student struct {
	Person     // Embedding Person struct (extends)
	course     string
	university string
}

func main() {
	person := Person{
		"John Doe",
		19,
		"123456789",
	}

	student1 := Student{
		Person:     person,
		course:     "Computer Science",
		university: "MIT",
	}

	fmt.Println(student1)

	student2 := Student{
		Person {
			"Larry Page",
			21,
			"987654321",
		},
		"Business Administration",
		"Harvard",
	}

	fmt.Println(student2)
}