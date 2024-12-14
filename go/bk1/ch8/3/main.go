package main

import "fmt"

func Filter[T any](s []T, f func(T)bool)[]T {
	r := make([]T,0,len(s))

	for _,v := range(s) {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

type Student struct {
	name string
	mark int
}

func main(){
	students := []Student{
		{"Roger", 90},
		{"Jerry", 51},
		{"Pat", 40},
		{"Joe", 45},
	}
	fltd := Filter(students, func(s Student) bool {
		return s.mark > 50
	})
	fmt.Println(fltd)
}
