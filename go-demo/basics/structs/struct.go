package employees

import "fmt"

type Employee struct { // to export variables they need to be capital letter
	Name     string
	Age      int
	IsRemote bool
}

func main() {

	employee1 := Employee{
		Name:     "Erik",
		Age:      30,
		IsRemote: true,
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
