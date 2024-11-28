package main

import "testing"

func TestEmployee(t *testing.T) {
	firstName := "Roger"
	lastName := "Okello"
	id := 1

	testCases := []struct {
		name     string
		firstName, lastName string
		id int
		option int
		expected Employee
	}{
		{
			"Option 1", firstName, lastName, id, 1,
			Employee{firstName, lastName, 1},
		},
		{
			"Option 2", firstName, lastName, id, 2,
			Employee{firstName, lastName, 1},
		},
		{
			"Option 3", firstName, lastName, id, 3,
			Employee{firstName, lastName, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
				result := NewEmployee(tc.firstName,tc.lastName,tc.id,tc.option)
				if result != tc.expected {
						t.Errorf("Expected %+v, got %+v", tc.expected, result)
				}
		})
	}

}