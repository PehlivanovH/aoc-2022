package main

import ("testing")

func TestCleaningCompleteOverlap(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	result := sunOverlap(input, completeOverlap);

	if result != 2 {
		t.Errorf("Expected 2 but got %d", result)
	}
}

func TestCleaningAnyOverlap(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	result := sunOverlap(input, anyOverlap);

	if result != 4 {
		t.Errorf("Expected 4 but got %d", result)
	}
}