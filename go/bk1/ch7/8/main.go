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
	i, ok := i.(Person)
	if !ok {
		fmt.Printf("Invalid type specified for type with value %T \n", i)
	}

	j, ok := i.(Employee)

	if !ok {
		fmt.Printf("Invalid type specified for type with value %v \n", j)
	}
}