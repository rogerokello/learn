package main

import "testing"

func TestDouble(t *testing.T){
	intTests := []struct{
		param int
		result int
	}{
		{2,4},
		{3,6},
		{4,8},
	}

	for _,v := range intTests {
		result := Double(v.param)

		if result != v.result {
			t.Errorf("Double(%d) = %d; want %d", v.param, result, v.result)
		}
	}

	float32Tests := []struct{
		param float32
		result float32
	}{
		{1.1, 2.2},
		{2.2, 4.4},
		{3.3, 6.6},
	}

	for _,v := range float32Tests {
		result := Double(v.param)

		if result != v.result {
			t.Errorf("Double(%f) = %f; want %f", v.param, result, v.result)
		}
	}

	float64Tests := []struct{
		param float32
		result float32
	}{
		{1.1, 2.2},
		{2.2, 4.4},
		{3.3, 6.6},
	}

	for _,v := range float64Tests {
		result := Double(v.param)

		if result != v.result {
			t.Errorf("Double(%f) = %f; want %f", v.param, result, v.result)
		}
	}

}

