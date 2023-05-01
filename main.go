package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type Hand int

func (hand Hand) val() int {
	return int(hand)
}

const (
	Rock Hand = iota + 1
	Paper
	Scissors

	Win  = 6
	Draw = 3
)

type HandPairs struct {
	Opponent Hand
	Mine     Hand
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
		panic("Wrong input[" + hand + "]")
	}
}

func calculateHands(opponent Hand, mine Hand) int {
	if opponent == mine {
		// draw
		return opponent.val() + Draw
	}
	switch opponent {
	case Rock:
		if mine == Paper {
			// win Paper > Rock < Paper
			return Paper.val() + Win
		}
		// lose Rock > Scissors
		return Scissors.val()
	case Paper:
		if mine == Rock {
			return Rock.val()
		}
		return Scissors.val() + Win
	case Scissors:
		if mine == Rock {
			return Rock.val() + Win
		}
		return Paper.val()
	}
	panic("oups")
}

func parseInput(input string) []HandPairs {
	handPairs := []HandPairs{}
	hands := strings.Split(input, "\n")
	for _, v := range hands {
		if len(v) == 0 {
			break
		}
		battle := strings.Split(v, " ")
		newHandPairs := HandPairs{getHand(battle[0]), getHand(battle[1])}
		handPairs = append(handPairs, newHandPairs)
	}
	return handPairs
}

func main() {
	handPairs := parseInput(input)
	totalPoints := 0
	for _, v := range handPairs {
		totalPoints += calculateHands(v.Opponent, v.Mine)
	}
	fmt.Println(totalPoints)
}
