package main

import (
	"testing"
)

func TestCaloriesCount(t *testing.T) {
	inventory := `1000
	2000
	3000`

	result := countCalories(inventory)

	if result != 6000 {
		t.Errorf("Expected 6000 but got %d", result)
	}
}

func TestCaloriesCount_WithMultipleInventories(t *testing.T) {
	inventories := "1000\n2051\n\n4000\n3000"

	result := mostCalirues(inventories)

	if result != 7000 {
		t.Errorf("Expected 7000 but got %d", result)
	}
}
