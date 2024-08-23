package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type JsonType struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonData := JsonType{"John", 30, "john@email.com"}
	valueJSON, err := json.Marshal(jsonData)

	if err != nil {
		log.Fatal("JSON marshaling failed: ", err)
	}

	fmt.Printf("\nMarshel: %s", bytes.NewBuffer(valueJSON))

	jsonInString := `{"name":"Mike","age":18,"email":"mike@email.com"}`

	var jsonToStruct JsonType

	if erro := json.Unmarshal([]byte(jsonInString), &jsonToStruct); err != nil {
		log.Fatal("JSON unmarshaling failed: ", erro)
	}

	fmt.Printf("\nUnmarshel: %v\n", jsonToStruct)
}