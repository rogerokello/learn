package main

import (
	"fmt"
	"strconv"
)

var opMap = map[string] func (int, int) int {}

func main() {

	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mult
	opMap["/"] = div

	items := [][]string{
		{"2", "+", "3"},
		{"3", "-", "2"},
		{"5", "*", "6"},
		{"6", "/", "2"},
	}

	for _, item := range items {
		oprnd1, err := strconv.Atoi(item[0])
		if err != nil {
			fmt.Println("Invalid option for the first operand, it must be a number")
			continue
		}
		operation, ok := opMap[item[1]]
		if !ok {
			fmt.Println("Invalid option for operator, it must be either +,-,* or / ")
			continue
		}
		oprnd2, err := strconv.Atoi(item[2])
		if err != nil {
			fmt.Println("Invalid option for the first operand, it must be a number")
			continue
		}
		result := operation(oprnd1, oprnd2)
		fmt.Println("Result is", result)
	}

}

func add (a int, b int) int {
	return a + b
}

func mult (a int, b int) int {
	return a * b
}

func sub (a int, b int) int {
	return a - b
}

func div (a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}