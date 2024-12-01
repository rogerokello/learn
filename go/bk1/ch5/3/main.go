package main

import (
	"fmt"
)


func main() {
  baseBy10 := baseMultiplier(10)
	baseBy2 := baseMultiplier(2)
	baseBy10([]int{1,2,3,4})
	baseBy2([]int{1,2,3,4})
}

func baseMultiplier(base int) func([]int) {
	return func(items []int) {
		itemsMultiplied := []int{}
		for _, item := range items {
			itemsMultiplied = append(itemsMultiplied, base * item)
		}
		fmt.Println(itemsMultiplied)
	}
}