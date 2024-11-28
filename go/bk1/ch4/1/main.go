package main

import (
	"fmt"
	"math/rand"
)

func main() {
	iRand := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		n := rand.Intn(101)
		iRand = append(iRand, n)
	}

	fmt.Println(iRand)
}