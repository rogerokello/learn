package main

import "fmt"


func main() {
	deferExample()
}

func deferExample(){
	a := 10

	defer func(b int) {
		fmt.Println(" Defer 1 ", b)
	}(a)

	a = 30
	defer func(b int) {
		fmt.Println(" Defer 2 ",b)
	}(a)

	a = 40
	defer func(b int) {
		fmt.Println(" Defer 3 ",b)
	}(a)

}