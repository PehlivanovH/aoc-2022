package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hasOverlap func(startFirst int, endFirst int, startSecond int, endSecond int) bool

func completeOverlap(startFirst int, endFirst int, startSecond int, endSecond int) bool {
	return startFirst <= startSecond && endFirst >= endSecond
}

func anyOverlap(startFirst int, endFirst int, startSecond int, endSecond int) bool {
	return (startSecond >= startFirst && startSecond <= endFirst) || (endSecond >= startFirst && endSecond <= endFirst)
}

func main() {
	input, _ := os.ReadFile("input")
	strInput := string(input)

	fmt.Println(sunOverlap(strInput, completeOverlap))
	fmt.Println(sunOverlap(strInput, anyOverlap))
}

func sunOverlap(assignmentPairs string, hasOverlap hasOverlap) int {

	var sum int
	for _, line := range strings.Split(assignmentPairs, "\n") {
		pairs := strings.Split(line, ",")
		if overlap(pairs[0], pairs[1], hasOverlap) {
			sum++
		}
	}

	return sum
}

func overlap(first string, second string, hasOverlap hasOverlap) bool {
	firstSections := strings.Split(first, "-")
	secondSections := strings.Split(second, "-")

	startFirst, _ := strconv.Atoi(firstSections[0])
	endFirst, _ := strconv.Atoi(firstSections[1])
	startSecond, _ := strconv.Atoi(secondSections[0])
	endSecond, _ := strconv.Atoi(secondSections[1])

	return hasOverlap(startFirst, endFirst, startSecond, endSecond) || hasOverlap(startSecond, endSecond, startFirst, endFirst)

}
