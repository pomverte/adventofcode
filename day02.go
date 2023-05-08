package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode/utf8"
)

//go:embed day02.txt
var inputDay2 string

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

func parseInput(input string) []HandPairs {
	handPairs := []HandPairs{}
	hands := strings.Split(input, "\n")
	for _, v := range hands {
		if utf8.RuneCountInString(v) == 0 {
			break
		}
		battle := strings.Split(v, " ")
		newHandPairs := HandPairs{getHand(battle[0]), getStrategy(battle[1])}
		handPairs = append(handPairs, newHandPairs)
	}
	return handPairs
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

func day2() {
	handPairs := parseInput(inputDay2)
	totalPoints := 0
	for _, v := range handPairs {
		totalPoints += calculateHands(v.Opponent, v.Mine)
	}
	fmt.Println(totalPoints)
}
