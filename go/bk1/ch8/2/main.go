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

func Reduce[T1, T2 any](s []T1, initializer T2,f func(T2,T1)T2 ) T2 {
	r := initializer

	for _,v := range s {
		r = f(r, v)
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
	items1 := Map(items, func(s string) string {
		return strconv.Itoa(len(s))
	})
	fmt.Println(items1)
	
	// We intend to get the length of all the items
	// so we first get all their length stored somewhere
	itemLengths := Map(items, func(s string) int {
		return len(s)
	})
	fmt.Println(itemLengths)

	sums := Reduce(itemLengths, 0, func(acc int, val int) int{
		return acc + val
	})
	fmt.Println(sums)
}