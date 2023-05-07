package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input-day2.txt
var inputDay2 string

//go:embed input-day3.txt
var inputDay3 string

type Hand int

func (hand Hand) val() int {
	return int(hand)
}

type Strategy int

func (strat Strategy) val() int {
	return int(strat)
}

const (
	Rock Hand = iota + 1
	Paper
	Scissors

	Lose Strategy = 0
	Draw Strategy = 3
	Win  Strategy = 6
)

type HandPairs struct {
	Opponent Hand
	Mine     Strategy
}

func getStrategy(strategy string) Strategy {
	switch strategy {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		panic("Wrong strategy[" + strategy + "]")
	}
}

func getHand(hand string) Hand {
	switch hand {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("Wrong hand[" + hand + "]")
	}
}

func calculateHands(opponent Hand, mine Strategy) int {
	if mine == Draw {
		return Draw.val() + opponent.val()
	}
	if mine == Win {
		switch opponent {
		case Rock:
			return Win.val() + Paper.val()
		case Paper:
			return Win.val() + Scissors.val()
		case Scissors:
			return Win.val() + Rock.val()
		}
	}
	if mine == Lose {
		switch opponent {
		case Rock:
			return Lose.val() + Scissors.val()
		case Paper:
			return Lose.val() + Rock.val()
		case Scissors:
			return Lose.val() + Paper.val()
		}
	}
	return 0
}

func parseInput(input string) []HandPairs {
	handPairs := []HandPairs{}
	hands := strings.Split(input, "\n")
	for _, v := range hands {
		if len(v) == 0 {
			break
		}
		battle := strings.Split(v, " ")
		newHandPairs := HandPairs{getHand(battle[0]), getStrategy(battle[1])}
		handPairs = append(handPairs, newHandPairs)
	}
	return handPairs
}

func day2() {
	handPairs := parseInput(inputDay2)
	totalPoints := 0
	for _, v := range handPairs {
		totalPoints += calculateHands(v.Opponent, v.Mine)
	}
	fmt.Println(totalPoints)
}

type Rucksack struct {
	compartmentOne string
	compartmentTwo string
}

func (r Rucksack) getItems() string {
	return r.compartmentOne + r.compartmentTwo
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

func day3() {
	rucksacks := parseRucksack(inputDay3)
	commonItemPrioritySum := 0
	for _, v := range rucksacks {
		// fmt.Printf("%c %v\n", findCommonItem(v), getPriority(findCommonItem(v)))
		commonItemPrioritySum += getPriority(findCommonItem(v))
	}
	fmt.Printf("commonItemPrioritySum : %v\n", commonItemPrioritySum)

	teamBadgePriority := 0
	for i := 0; i < len(rucksacks); i += 3 {
		fmt.Printf("team %v : %v, %v, %v\n", i, rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		teamBadgePriority += getPriority(findTeamBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2]))
	}
	fmt.Printf("teamBadgePriority : %v\n", teamBadgePriority)

}

func main() {
	day3()
}
