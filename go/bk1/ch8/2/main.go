package main

import (
	"fmt"
	"strconv"
)

func Map[T1, T2 any](s []T1, f func(T1)T2) []T2 {
	r := make([]T2, 0, len(s))

	for _,v := range s {
		r = append(r, f(v))
	}
	return r
}

func main(){
	items := []string {
		"Banana",
		"Orange",
		"Apple",
		"Pawpaw",
		"egg",
	}
	items = Map(items, func(s string) string {
		return strconv.Itoa(len(s))
	})
	fmt.Println(items)
}