package main

import (
	"fmt"
)

type Queue[T comparable] struct {
	vals []T
}

func (q *Queue[T]) add(item T) {
	q.vals = append(q.vals, item)
}

func (q *Queue[T]) remove() (bool, T) {
	qL := len(q.vals)
	var zero T

	// if nothing is in the queue
	// return nothing
	if qL == 0 {
		return false, zero
	}

	var next T
	// If we only have one item in the queue
	// extract that item and generate a new empty queue
	if qL == 1 {
		next = q.vals[0]
		q.vals = []T{}
		return true, next
	}

	// Other wise extract the first item
	// and truncate the rest
	next = q.vals[0]
	q.vals = q.vals[1:]
	return true, next
}

func (qu *Queue[T]) Contains(item T) bool {
	for _, val := range qu.vals {
		if val == item {
			return true
		}
	}
	return false
}

func main() {
	qu := &Queue[int]{}
	qu.add(1)
	qu.add(2)
	qu.add(3)
	qu.add(4)
	fmt.Println("Does 4 exist in qu: ", qu.Contains(4))

	fmt.Println(qu.remove())
	fmt.Println(qu.remove())
	fmt.Println(qu.remove())
	fmt.Println(qu.remove())
	fmt.Println(qu.remove())
	fmt.Println(qu.remove())

	fmt.Println("Does 4 exist in qu: ", qu.Contains(4))
}