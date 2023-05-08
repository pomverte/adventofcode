package main

import (
	_ "embed"
)

func runesEquals(r1, r2 []rune) bool {
	if len(r1) != len(r2) {
		return false
	}
	for _, i := range r1 {
		found := false
		for _, j := range r2 {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func intListEquals(i1, i2 []int) bool {
	if len(i1) != len(i2) {
		return false
	}
	for _, i := range i1 {
		found := false
		for _, j := range i2 {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	day03()
	day04()
}
