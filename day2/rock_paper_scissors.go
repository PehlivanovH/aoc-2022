package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var firstRound = map[string]map[string]int{
		"Y": { // paper +2
			"A": 8, // rock
			"B": 5, // paper
			"C": 2, // scissors
		},
		"X": { // rock +1
			"A": 4,
			"B": 1,
			"C": 7,
		},
		"Z": { // scissors +3
			"A": 3,
			"B": 9,
			"C": 6,
		},
	}

	var secondRound = map[string]map[string]int{
		"Y": { // draw +3
			"A": 4, // rock
			"B": 5, // paper
			"C": 6, // scissors
		},
		"X": { // lose 0
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Z": { // win +6
			"A": 8,
			"B": 9,
			"C": 7,
		},
	}

	fmt.Println(play(firstRound))
	fmt.Println(play(secondRound))
}

func play(rules map[string]map[string]int) int {
	input, _ := os.ReadFile("input")
	strInput := string(input)
	rounds := strings.Split(strInput, "\n")

	var sum int
	for _, round := range rounds {
		hands := strings.Split(round, " ")
		opponent := hands[0]
		me := hands[1]
		sum += rules[me][opponent]
	}

	return sum
}
