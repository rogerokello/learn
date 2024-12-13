package main

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) push(item T) {
	s.vals = append(s.vals, item)
}

func (s *Stack[T]) pop()(bool, T) {
	if len(s.vals) == 0 {
		var zero T
		return false, zero
	}

	item := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return true, item
}

func main() {
	b := &Stack[int]{}
	b.push(1)
	b.push(2)
	b.push(3)
	fmt.Println(b.pop())
	fmt.Println(b.pop())
	fmt.Println(b.pop())
}