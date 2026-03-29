package main

import "fmt"

type Employee struct {
	name     string
	age      int
	isRemote bool
}

func main() {

	employee1 := Employee{
		name:     "Erik",
		age:      30,
		isRemote: true,
	}

	fmt.Print(employee1)

	job := struct { // anonimous struct
		title  string
		salary int
	}{ // initialization
		title:  "Sofrware Engineer",
		salary: 100000,
	}

	fmt.Print(job.title, job.salary)

}
