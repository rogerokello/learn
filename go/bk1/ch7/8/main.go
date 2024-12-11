package main

import "fmt"

type Person struct{
	FirstName string
	LastName  string
	Age       int
}

type Employee struct {
	Person
}

func main() {
	var i any = Person{
		"Roger",
		"Okello",
		20,
	}
	j, ok := i.(Person)
	if !ok {
		fmt.Printf("Invalid type specified for type with value %T \n", j)
	}

	k, ok := i.(Employee)

	if !ok {
		fmt.Printf("Invalid type specified for type with value %v \n", k)
	}
}