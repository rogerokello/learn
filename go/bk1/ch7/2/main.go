package main

import (
	"math/rand"
)

type Tree struct {
	val int
	left, right *Tree
}

func (t *Tree) Insert(val int) *Tree{
	if t == nil {
		return &Tree{val:val}
	}
	if val < t.val {
		t.left = t.left.Insert(val)
	}
	if val > t.val {
		t.right = t.right.Insert(val)
	}
	return t
}

func main() {
	var tPtr *Tree
	for i:=0; i <10_000; i++ {
		a := rand.Intn(10_000)
		tPtr = tPtr.Insert(a)
	}
}
