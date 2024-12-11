package main

import "fmt"

type IntCustom int
type FloatCustom float64
type Person struct {
	FirstName string
	LastName  string
	Age       int
}


func main() {
	var a int = 1
	var b IntCustom = 2
	var c FloatCustom = 2.1
	var d Person = Person{}
	var i any = a
	i = b
	i = c
	i = d

	switch i := i.(type) {
	case int:
		fmt.Printf("Looks like %v is an int \n", i)
	case IntCustom:
		fmt.Printf("Looks like %v is an IntCustom \n", i)
	case FloatCustom:
		fmt.Printf("Looks like %v is an FloatCustom \n", i)
	case Person:
		fmt.Printf("Looks like %v is an Person \n", i)
	}
}