package main

import "testing"

func TestIntListEquals(t *testing.T) {
	x := []int{2, 3, 4}
	y := []int{3, 4, 2}
	result := intListEquals(x, y)
	if !result {
		t.Errorf("intListEquals(%v, %v) returned %v", x, y, result)
	}
}
