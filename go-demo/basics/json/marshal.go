package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	pers := Person{"Test", 40}

	data, _ := json.Marshal(pers)

	jsonStr := `{"name":"Alice","age":30}`

	var user Person
	err := json.Unmarshal([]byte(jsonStr), &user)

	if err != nil {
		fmt.Print("невалидный JSON:", err)
	}

	fmt.Print(user)

	fmt.Print(string(data))
	fmt.Print(pers)
}
