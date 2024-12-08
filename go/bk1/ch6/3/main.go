package main

import "fmt"

func main() {

	c := []string{
		"last position",
	}

	fmt.Println("c", c, "len c", len(c), "Cap C", cap(c))
	UpdateSlice(c, "value1")
	fmt.Println("c", c, "len c", len(c), "Cap C", cap(c))
	GrowSlice(c, "value2")
	fmt.Println("c", c, "len c", len(c), "Cap C", cap(c))
}

func UpdateSlice(a []string, b string) {
	a[len(a)-1] = b
}

func GrowSlice(a []string, b string) {
	_ = append(a, b)
}