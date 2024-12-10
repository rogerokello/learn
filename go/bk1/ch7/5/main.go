package main

import "fmt"

type Incrementer interface {
	Increment()
}

type Counter struct {
	val int
}

func (c *Counter) Increment() {
	c.val++
}

func main() {
	var c1 *Counter
	var c Counter
	var i1 Incrementer = c1

	fmt.Println(c, i1)

	/* 
		var c does not contain Increment in its method set because it is not a pointer.
		Only variables pointing to type counter will contain Increment in their method
		set because the Increment method has been defined with a pointer reciever.
		For this reason you cannot assign c to an interface defined with Increment as a
		method
	*/
	// i1 = c

}