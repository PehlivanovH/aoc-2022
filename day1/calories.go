package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	strInput := string(input)

	fmt.Println(mostCalirues(strInput))
	fmt.Println(sumTop3Calories(string(input)))
}

func countCalories(elf string) int {
	items := strings.Split(elf, "\n")

	var sum int
	for _, item := range items {
		caloriesForItem, _ := strconv.Atoi(strings.TrimSpace(item))
		sum += caloriesForItem
	}

	return sum
}

func mostCalirues(inventories string) int {
	elves := strings.Split(inventories, "\n\n")

	caloriesForElf := make([]int, 0)
	for _, elf := range elves {
		caloriesForElf = append(caloriesForElf, countCalories(elf))
	}

	sort.Ints(caloriesForElf)

	return caloriesForElf[len(caloriesForElf)-1]
}

func sumTop3Calories(inventories string) int {
	elves := strings.Split(inventories, "\n\n")

	caloriesForElf := make([]int, 0)
	for _, elf := range elves {
		caloriesForElf = append(caloriesForElf, countCalories(elf))
	}

	sort.Ints(caloriesForElf)
	

	top3 := caloriesForElf[len(caloriesForElf)-3:]

	return top3[0] + top3[1] + top3[2]
}