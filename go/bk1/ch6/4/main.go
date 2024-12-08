package main

// import (
// 	"fmt"
// 	"time"
// )

type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	people := make([]Person, 0)
	// people := make([]Person, 10_000_000)
	// start := time.Now()

	for i:=0; i < 10_000_000 ; i++ {
		people = append(people,Person{"Okello","Roger",12})
	}

	// fmt.Println("Time for creation", time.Since(start))
}
