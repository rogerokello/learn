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