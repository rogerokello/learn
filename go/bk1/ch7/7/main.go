package main

import "fmt"

type MyInt int

func main() {
	var i any
	var j MyInt = 20
	i = j
	isMyInt := i.(MyInt)
	fmt.Println(isMyInt)

	// You cannot use type assertions even if two types share the
	// same underlying type. A type assertion must match the type
	// of the value stored in the interface.
	// The code below will panic if we do not add the ok, idiom
	// isMyInt = i.(int)
}
