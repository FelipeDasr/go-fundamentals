package main

import "fmt"

type User struct {
	name string
	age  uint
	address Address
}

type Address struct {
	city string
	street string
	number uint
	zipCode string
}

func main() {
	address := Address{"New York", "5th Avenue", 123, "10001"}

	var user1 User
	user1.name = "John"
	user1.age = 30
	user1.address = address

	fmt.Println(user1)

	user2 := User{"David", 13, address}
	fmt.Println(user2)

	uncompletedUser := User{name: "Alice", address: Address{"Los Angeles", "Sunset Boulevard", 456, "90001"}}
	fmt.Println(uncompletedUser)
}