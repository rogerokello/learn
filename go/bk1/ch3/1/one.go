package main

import "fmt"


func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}

	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[3:]
	
	fmt.Println("First", first)
	fmt.Println("Second", second)
	fmt.Println("Third", third)
}