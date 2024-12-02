package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
)

const maxDiff = 3

func main() {
	file := util.FileFromArgs()
	defer file.Close()

	scan := bufio.NewScanner(file)
	var safeCount int
	var secondSafeCount int
	var errorIndex int

	for scan.Scan() {
		row := util.StringToInts(scan.Text(), " ")

		if errorIndex = findErrorIndex(row, -1); errorIndex == -1 {
			safeCount++
			secondSafeCount++
		} else {
			// Try removing any of the values arround the error index
			for i := range 3 {
				if findErrorIndex(row, errorIndex+1-i) == -1 {
					secondSafeCount++
					break
				}
			}
		}
	}

	fmt.Println(safeCount)
	fmt.Println(secondSafeCount)
}

func findErrorIndex(row []int, skip int) int {
	var diff int
	var sign int8

	for i := 0; i < len(row)-1; i++ {
		if i == skip {
			continue
		} else if i+1 == skip {
			if i+2 == len(row) {
				continue
			}

			diff = row[i] - row[i+2]
		} else {
			diff = row[i] - row[i+1]
		}

		// Diff bigger than max, unsafe!				  // direction changed, unsafe!
		if diff == 0 || diff > maxDiff || diff < -maxDiff || (diff < 0 && sign == 1) || (diff > 0 && sign == -1) {
			return i
		}

		// Track the direction, for future usage
		if sign == 0 {
			if diff < 0 {
				sign = -1
			} else {
				sign = 1
			}
		}
	}

	return -1
}
