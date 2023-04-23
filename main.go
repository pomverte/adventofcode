package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	topCalories := 0
	inputs := strings.Split(input, "\n\n")
	for _, v := range inputs {
		caloriesPerElve := 0
		for _, v := range strings.Split(v, "\n") {
			cal, _ := strconv.Atoi(v)
			caloriesPerElve += cal
		}
		if topCalories < caloriesPerElve {
			topCalories = caloriesPerElve
		}
	}
	fmt.Printf("topCalories: %v\n", topCalories)
}
