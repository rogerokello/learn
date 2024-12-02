package main

import "fmt"

type person struct {
	FirstName string
	LastName *string
}

func main() {
	a := person{
		FirstName: "Roger",
	}
	b := "Okello"
	var c *string = &b
	a.LastName = c

	fmt.Println(*a.LastName)

	d := "otim"
	a.LastName = makePtr(d)
	fmt.Println(a)
}

func makePtr[T any](t T) (*T) {
	return &t
}