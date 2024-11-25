package main

import "fmt"


func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("B value is ", b)
	fmt.Println("smallI value is ", smallI)
	fmt.Println("bigI value is ", bigI)
	
}