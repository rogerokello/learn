package main

import "fmt"


func Double[T int | float32 | float64](a T) T{
	return 2*a
}

func main(){
	var a float32
	var b float64
	var c int

	a = 2.6
	b = 3.7
	c = 4

	fmt.Println("a", Double(a))
	fmt.Println("b", Double(b))
	fmt.Println("c", Double(c))
}