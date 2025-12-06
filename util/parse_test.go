package util

import "testing"

func TestIntsToStrings(t *testing.T) {
	input := []int{2, 4, 6, 8, 10, 12}

	output := IntsToStrings(input)

	for i := range output {
		if (i+1)*2 != MustAtoi(output[i]) {
			t.Errorf("Expected value %v at index %v, but got %v", (i+1)*2, i, MustAtoi(output[i]))
		}
	}
}

func TestStringToInts(t *testing.T) {
	output := StringToInts("2, 4, 6, 8, 10, 12", ", ")
	for i := range output {
		if (i+1)*2 != output[i] {
			t.Errorf("Expected value %v at index %v, but got %v", (i+1)*2, i, output[i])
		}
	}

	output = StringToInts("2   4 6  8 10     12", " ")
	for i := range output {
		if (i+1)*2 != output[i] {
			t.Errorf("Expected value %v at index %v, but got %v", (i+1)*2, i, output[i])
		}
	}
}
