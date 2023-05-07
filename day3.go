package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed day3.txt
var inputDay3 string

type Rucksack struct {
	compartmentOne string
	compartmentTwo string
}

func (r Rucksack) getItems() string {
	return r.compartmentOne + r.compartmentTwo
}

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

func parseRucksack(inputDay3 string) []Rucksack {
	rucksacks := []Rucksack{}
	for _, v := range strings.Split(inputDay3, "\n") {
		if len(v) == 0 {
			break
		}
		mid := len(v) / 2
		rucksacks = append(rucksacks, Rucksack{v[:mid], v[mid:]})
	}
	return rucksacks
}

func findCommonItem(rucksack Rucksack) rune {
	for _, one := range rucksack.compartmentOne {
		for _, two := range rucksack.compartmentTwo {
			if one == two {
				return one
			}
		}
	}
	panic("oups, there is no common item in " + rucksack.compartmentOne + " and " + rucksack.compartmentTwo)
}

func getPriority(letter rune) int {
	if unicode.IsLower(letter) {
		return int(unicode.ToUpper(letter) - 'A' + 1)
	}
	return int(unicode.ToUpper(letter) - 'A' + 1 + 26)
}

func findCommonItemBetweenRucksack(r1, r2 Rucksack) []rune {
	set := make(map[rune]bool)
	for _, a := range r1.getItems() {
		for _, b := range r2.getItems() {
			if a == b {
				set[a] = true
			}
		}
	}
	result := []rune{}
	for k := range set {
		result = append(result, k)
	}
	return result
}

func findTeamBadge(r1, r2, r3 Rucksack) rune {
	for _, v := range findCommonItemBetweenRucksack(r1, r2) {
		for _, i := range r3.getItems() {
			if v == i {
				return v
			}
		}
	}
	panic("No common item found")
}

func day3() {
	rucksacks := parseRucksack(inputDay3)
	commonItemPrioritySum := 0
	for _, v := range rucksacks {
		commonItemPrioritySum += getPriority(findCommonItem(v))
	}
	fmt.Printf("commonItemPrioritySum : %v\n", commonItemPrioritySum)

	teamBadgePriority := 0
	for i := 0; i < len(rucksacks); i += 3 {
		teamBadgePriority += getPriority(findTeamBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2]))
	}
	fmt.Printf("teamBadgePriority : %v\n", teamBadgePriority)

}
