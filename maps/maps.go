package main

import "fmt"

func main() {

	user := map[string]string{
		"name": "John",
		"age":  "25",
	}

	fmt.Println("Name: ", user["name"])
	fmt.Println("Age: ", user["age"])

	user["document"] = "123456789"
	fmt.Println("Document: ", user["document"])
	
	delete(user, "document")
}