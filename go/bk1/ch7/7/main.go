package main

import "fmt"

type MyInt int

func main() {
	var i any
	var j MyInt = 20
	i = j
	isMyInt := i.(MyInt)
	fmt.Println(isMyInt)
}
