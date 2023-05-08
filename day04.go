package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

//go:embed day04.txt
var inputDay04 string

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

func getNbSectionsFullyContained(inputDay04 string) int {
	result := 0
	for _, p := range strings.Split(inputDay04, "\n") {
		if utf8.RuneCountInString(p) == 0 {
			break
		}
		if sectionsFullyContained(p) {
			result++
		}
	}
	return result
}

func sectionsOverlapped(pair string) bool {
	p := strings.Split(pair, ",")
	s1 := strings.Split(p[0], "-")
	s2 := strings.Split(p[1], "-")
	// FIXME no error handling
	x, _ := strconv.Atoi(s1[0])
	y, _ := strconv.Atoi(s1[1])
	i, _ := strconv.Atoi(s2[0])
	j, _ := strconv.Atoi(s2[1])
	return (x <= i && y >= i) || (x >= i && x <= j)
}

func getNbSectionsOverlapped(inputDay04 string) int {
	result := 0
	for _, p := range strings.Split(inputDay04, "\n") {
		if utf8.RuneCountInString(p) == 0 {
			break
		}
		if sectionsOverlapped(p) {
			result++
		}
	}
	return result
}

func day04() {
	fmt.Println("=== DAY 4")
	fmt.Printf("Number of assignment pairs where one range fully contain the other : %v\n", getNbSectionsFullyContained(inputDay04))
	fmt.Printf("Number of assignment pairs with ranges overlaping : %v\n", getNbSectionsOverlapped(inputDay04))
}
