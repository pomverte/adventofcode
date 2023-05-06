package main

import (
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
	var expected int32 = 16
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
