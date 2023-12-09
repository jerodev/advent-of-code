package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"slices"
)

func main() {
	file := util.FileFromArgs()

	line := ""
	b := make([]byte, 1)

	nextValueSum := 0
	prevValueSum := 0

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			history := util.StringToInts(line, " ")
			nextValueSum += findNext(history)

			// Part 2: Calculate previous value
			slices.Reverse(history)
			prevValueSum += findNext(history)

			if err == io.EOF {
				break
			}

			line = ""
		}

		line += string(b[0])
	}

	fmt.Println(nextValueSum)
	fmt.Println(prevValueSum)
}

func findNext(history []int) int {
	sequences := [][]int{
		history,
	}

	lastDiff := history
	value := 0
	for {
		lastDiff = calculateDiffs(lastDiff)

		allSame := true
		value = lastDiff[0]
		for _, n := range lastDiff {
			if n != value {
				allSame = false
				break
			}
		}

		if allSame {
			break
		}

		sequences = append(sequences, lastDiff)
	}

	// Add value to the new value in the last sequence, spare some loops
	nextValue := sequences[len(sequences)-1][len(sequences[len(sequences)-1])-1] + value

	// Now calculate a new value for all other sequences
	for i := len(sequences) - 2; i >= 0; i-- {
		nextValue = sequences[i][len(sequences[i])-1] + nextValue
	}

	return nextValue
}

func calculateDiffs(history []int) []int {
	diffs := []int{}

	for i := range history {
		if i == 0 {
			continue
		}

		diffs = append(diffs, history[i]-history[i-1])
	}

	return diffs
}
