package main

import "fmt"

type person struct{
	FirstName string
	LastName string
	Age int
}

func main() {

	person1 := MakePerson("Roger", "Okello", 34)
	person2 := MakePersonPointer("Roger", "Okello", 34)

	fmt.Println(person1)
	fmt.Println(person2)

}

func MakePerson (FirstName string, LastName string, Age int) person {
	return person{
		FirstName: FirstName,
		LastName: LastName,
		Age: Age,
	}
}

func MakePersonPointer (FirstName string, LastName string, Age int) *person {
	return &person{
		FirstName: FirstName,
		LastName: LastName,
		Age: Age,
	}
}
