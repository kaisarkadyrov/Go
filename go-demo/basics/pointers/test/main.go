package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	IsAvailable bool   `json:"isAvailable"`
}

func (p *Person) ChangeName(name string) string {
	p.Name = name
	return p.Name
}

func main() {

	p1 := Person{"Alex", 18, true}

	fmt.Println(p1.ChangeName("Steve"))

	data, _ := json.Marshal(p1)

	fmt.Println(string(data))
}
