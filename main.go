package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	var caloriesForElves []int
	inputs := strings.Split(input, "\n\n")

	for _, v := range inputs {
		caloriesPerElve := 0
		for _, v := range strings.Split(v, "\n") {
			cal, _ := strconv.Atoi(v)
			caloriesPerElve += cal
		}
		caloriesForElves = append(caloriesForElves, caloriesPerElve)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesForElves)))

	top3CaloriesPerElveTotal := 0
	for i := 0; i < 3; i++ {
		top3CaloriesPerElveTotal += caloriesForElves[i]
	}
	fmt.Println(top3CaloriesPerElveTotal)
}
