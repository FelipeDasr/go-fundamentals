package main

import "fmt"

type User struct {
	name string
	age  int
}

func (user User) showInfo() {
	fmt.Println("Name:", user.name)
	fmt.Println("Age:", user.age)
}

func (user *User) changeName(name string) {
	user.name = name
}

func main() {
	user := User{name: "John", age: 30}
	user.showInfo()

	user.changeName("Alice")
	user.showInfo()
}