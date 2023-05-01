package main

import (
	_ "embed"
	"fmt"
	"strings"
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

func day3() {
	rucksacks := parseRucksack(inputDay3)
	for _, v := range rucksacks {
		fmt.Printf("%c\n", findCommonItem(v))
	}
}

func main() {
	day3()
}
