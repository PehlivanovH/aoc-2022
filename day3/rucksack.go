package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	strInput := string(input)

	fmt.Println(sumForAllRucksacks(strInput))
	fmt.Println(sumForGroupItems(strInput))
}

func sumForGroupItems(input string) int {
	var sum int
	var counter int
	var group [3]string
	for _, rucksack := range strings.Split(input, "\n") {
		group[counter] = rucksack
		if (counter == 2) {
			sum += findBadge(group)
			counter = 0
			continue
		}
		counter++
	}

	return sum
}

func findBadge(group [3]string) int {
	badge := findCommonInGroup(group[0], group[1], group[2])
	return code(badge)
}

func findCommonInGroup(first string, second string, third string) rune {
	for _, item := range first {
		if strings.Contains(second, string(item)) && strings.Contains(third, string(item)) {
			return item
		}
	}
	return -1
}

func code(item rune) int {
	if item > 96 { // lower case
		return int((item) - 96)
	} else { // upper case
		return int((item - 38))
	}
}

func sumForAllRucksacks(input string) int {
	var sum int
	for _, rucksack := range strings.Split(input, "\n") {
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		commonItem := findCommon(compartment1, compartment2)
		sum += code(commonItem)
	}

	return sum
}

func findCommon(compartment1 string, compartment2 string) rune {
	for _, item := range compartment1 {
		if strings.Contains(compartment2, string(item)) {
			return item
		}
	}
	return -1
}
