package main

import "testing"

func TestGreet(t *testing.T) {
	expected := "Yo Tsap"

	result := Greet()

	if expected != result {
		t.Errorf("Greet() = %q; want %q", result, expected)
	}
}