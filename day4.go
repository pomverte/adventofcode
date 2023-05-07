package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day4.txt
var inputDay4 string

func sectionsFullyContained(pair string) bool {
	p := strings.Split(pair, ",")
	s1 := strings.Split(p[0], "-")
	s2 := strings.Split(p[1], "-")
	x, _ := strconv.Atoi(s1[0])
	y, _ := strconv.Atoi(s1[1])
	i, _ := strconv.Atoi(s2[0])
	j, _ := strconv.Atoi(s2[1])
	return (x <= i && y >= j) || (x >= i && y <= j)
}

func getNbSectionsFullyContained(inputDay4 string) int {
	result := 0
	for _, p := range strings.Split(inputDay4, "\n") {
		if len(p) == 0 {
			break
		}
		if sectionsFullyContained(p) {
			result++
		}
	}
	return result
}

func day4() {
	fmt.Printf("Number of assignment pairs where one range fully contain the other : %v\n", getNbSectionsFullyContained(inputDay4))
}
