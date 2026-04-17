package main

import "fmt"

type Stringer interface {
	String() string
}

type Dog struct {
	Name string
}

func (d Dog) String() string {
	return fmt.Sprintf("Hello, %s", d.Name)
}

func print(i any) {
	if v, ok := i.(Stringer); ok {
		fmt.Println(v.String())
	}
}	

func main() {
	var i any = "hello"

	fmt.Println(i.(string))
	print(Dog{"BOBIK"})
}
