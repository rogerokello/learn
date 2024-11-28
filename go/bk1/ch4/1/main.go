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
		fmt.Println(v)
		if v % 2 == 0 && v % 3 == 0 {
			fmt.Println("Six")
			continue
		}

		if v % 2 == 0 {
			fmt.Println("Two")
			continue
		}

		if v % 3 == 0 {
			fmt.Println("Six")
			continue
		}

		fmt.Println("Never mind")

	}

}