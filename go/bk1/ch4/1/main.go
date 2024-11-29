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

	for _,v := range iRand {
		switch {
		case v % 2 == 0 && v % 3 == 0:
			fmt.Println(v)
			fmt.Println("Six")
		case v % 2 == 0:
			fmt.Println(v)
			fmt.Println("Two")
		case v % 3 == 0:
			fmt.Println(v)
			fmt.Println("Six")
		default:
			fmt.Println(v)
			fmt.Println("Never mind")
		}
	}

}