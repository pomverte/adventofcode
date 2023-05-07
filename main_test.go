package main

import (
	"fmt"
	"testing"
)

func TestDay3ParseRucksack(t *testing.T) {
	result := parseRucksack(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`)
	expected := []Rucksack{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
	}
	if len(result) != len(expected) {
		t.Errorf("parseRucksack(...) has %v items, expected %v", len(result), len(expected))
	}
	for k, v := range result {
		if v != expected[k] {
			t.Errorf("parseRucksack[%v] has %v, expected %v", k, v, expected[k])
		}
	}
}

func TestDay3FindCommonItem(t *testing.T) {
	result := findCommonItem(Rucksack{"vJrwpWtwJgWr", "hcsFMMfFFhFp"})
	expected := 'p'
	if result != expected {
		t.Errorf("findCommonItem(vJrwpWtwJgWr, hcsFMMfFFhFp) returned %c, expected %c", result, expected)
	}
}

func TestDay3GetPriority(t *testing.T) {
	result := getPriority('p')
	expected := 16
	if result != expected {
		t.Errorf("getPriority('A') returned %v, expected %v", result, expected)
	}

	result = getPriority('L')
	expected = 38
	if result != expected {
		t.Errorf("getPriority('L') returned %v, expected %v", result, expected)
	}

	result = getPriority('P')
	expected = 42
	if result != expected {
		t.Errorf("getPriority('P') returned %v, expected %v", result, expected)
	}

	result = getPriority('v')
	expected = 22
	if result != expected {
		t.Errorf("getPriority('v') returned %v, expected %v", result, expected)
	}

	result = getPriority('t')
	expected = 20
	if result != expected {
		t.Errorf("getPriority('t') returned %v, expected %v", result, expected)
	}

	result = getPriority('s')
	expected = 19
	if result != expected {
		t.Errorf("getPriority('t') returned %v, expected %v", result, expected)
	}
}

func TestRunesEquals(t *testing.T) {
	r1 := []rune{'F', 'M', 'f', 'r', 's'}
	r2 := []rune{'r', 's', 'F', 'M', 'f'}
	if !runesEquals(r1, r2) {
		t.Errorf("runesEquals(%c, %c) should be equals", r1, r2)
	}
}

func TestRunesEqualsFail1(t *testing.T) {
	r1 := []rune{'F', 'M', 'f', 'r', 's'}
	r2 := []rune{'r', 's', 'F', 'M', 'f', 't'}
	if runesEquals(r1, r2) {
		t.Errorf("runesEquals(%c, %c) should NOT be equals", r1, r2)
	}
}

func TestRunesEqualsFail2(t *testing.T) {
	r1 := []rune{'F', 'M', 'f', 'r', 'a'}
	r2 := []rune{'r', 's', 'F', 'M', 'f'}
	if runesEquals(r1, r2) {
		t.Errorf("runesEquals(%c, %c) should NOT be equals", r1, r2)
	}
}
func TestDay3FindCommonItemBetweenRucksack(t *testing.T) {
	result := findCommonItemBetweenRucksack(
		Rucksack{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		Rucksack{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"})
	expected := []rune{'r', 's', 'F', 'M', 'f'}
	for _, v := range expected {
		fmt.Println(v)
	}
	if !runesEquals(result, expected) {
		t.Errorf("findCommonItemBetweenRucksack(r1, r2) returned %c, expected %c", result, expected)
	}
}

func TestDay3FindTeamBadge1(t *testing.T) {
	result := findTeamBadge(
		Rucksack{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		Rucksack{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		Rucksack{"PmmdzqPrV", "vPwwTWBwg"})
	var expected rune = 'r'
	if result != expected {
		t.Errorf("findTeamBadge(team) returned %c, expected %c", result, expected)
	}
}

func TestDay3FindTeamBadge2(t *testing.T) {
	result := findTeamBadge(
		Rucksack{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
		Rucksack{"ttgJtRGJ", "QctTZtZT"},
		Rucksack{"CrZsJsPPZsGz", "wwsLwLmpwMDw"})
	var expected rune = 'Z'
	if result != expected {
		t.Errorf("findTeamBadge(team) returned %c, expected %c", result, expected)
	}
}
