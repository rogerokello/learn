package main

import "fmt"


type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{5}
	fmt.Println(myAdder.AddTo(10))

	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	tenAdder := Adder{10}

	f2 := Adder.AddTo
	fmt.Println(f2(tenAdder, 10))
}