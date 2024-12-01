package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	type Person struct {
		FirstName string
		LastName string
		Age      int
	}
	people := []Person {
		{"Roger", "Okello", 40},
		{"Billy", "Ayo", 45},
		{"Linda", "Acen", 38},
	}
	fmt.Println(people)
	sort.Slice(people, func(i,j int) bool {
		// return people[i].Age < people[j].Age
		return strings.Compare(people[i].FirstName, people[j].FirstName) == -1
	})
	fmt.Println(people)
}