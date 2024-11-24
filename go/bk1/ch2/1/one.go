package main

import "fmt"

func main() {
	var i int = 20

	// Literals do not have types and that is why assignment of 
	// var f float64 = 20 would work even in we are initially
	// assigning an integer literal. In the above case we need to
	// do the casting because we have already associated the value
	// with a type.
	var f float64 = float64(i)

	fmt.Println("i is ", i, "and f is ", f)

}