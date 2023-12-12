package main

import (
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	file, _ := os.Open("example.txt")

	result := run(file, 1)
	if result != 21 {
		t.Errorf("Expected 21, got %d", result)
	}
}

func TestExamplePartTwo(t *testing.T) {
	file, _ := os.Open("example.txt")

	result := run(file, 5)
	if result != 525152 {
		t.Errorf("Expected 21, got %d", result)
	}
}
