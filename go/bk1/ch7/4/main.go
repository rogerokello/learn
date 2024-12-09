package main

import "fmt"


type Score int
type HighScore int
type VerHighScore HighScore

func (h HighScore) hasBonus() bool {
	return true
}

func main() {
	var i Score = 100
	var v VerHighScore = 10000
	fmt.Println("v", v)
	fmt.Println("i", i)

	var j HighScore = HighScore(i)
	fmt.Println("j", j)

	fmt.Println(j.hasBonus())
}