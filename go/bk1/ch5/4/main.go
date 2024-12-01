package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct {
		FirstName string
		LastName string
		Age int
	}

	people := []Person {
		{"Roger", "Okello", 30},
		{"Billy", "Ayo", 31},
		{"Linda", "Acen", 32},
	}

	sort.Slice(people, func(i,j int) bool {
		return people[i].Age < people[j].Age
	})

	targetAge := 32

	i := sort.Search(len(people), func(i int) bool {
		return people[i].Age >= targetAge
	})

	if i < len(people) && people[i].Age == targetAge {
		// x is present at data[i]
		fmt.Println(people[i].Age)
	} else {
		fmt.Println("Not found")
	}
}