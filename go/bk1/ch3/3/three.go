package main

import "fmt"

type Employee struct {
	firstName string
	lastName string
	id int
}

func main() {
	first := Employee {
		"Susan",
		"Aguti",
		1,
	}

	second := Employee {
		firstName: "Roger",
		lastName: "Okello",
		id: 1,
	}

	var third Employee
	third.firstName = "Billy"
	third.lastName = "Ayo"
	third.id = 2

	fmt.Println("first", first)
	fmt.Println("second", second)
	fmt.Println("third", third)

}


func NewEmployee(firstName string, lastName string, id int, option int) Employee {

	if option == 1 {
		return Employee {
			firstName,
			lastName,
			id,
		}
	}
	if option == 2 {
		return Employee {
			firstName: firstName,
			lastName: lastName,
			id: id,
		}
	}

	// Default
	var third Employee
	third.firstName = firstName
	third.lastName = lastName
	third.id = id

	return third

}
