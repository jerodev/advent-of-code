package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
)

const symbol = '+'

func main() {
	file := util.FileFromArgs()

	// Map the grid
	b := make([]byte, 1)
	grid := [][]byte{}
	line := []byte{}
	for {
		_, err := file.Read(b)

		if err == io.EOF || b[0] == '\n' {
			grid = append(grid, line)
			line = []byte{}

			if err == io.EOF {
				break
			}

			continue
		}

		// Convert all symbols to pluses
		if b[0] != '.' && (b[0] < '0' || b[0] > '9') {
			b[0] = symbol
		}

		line = append(line, b[0])
	}

	// Loop the grid and find numbers
	sum := 0
	for row := 0; row < len(grid); row++ {
		number := 0

		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '.' || grid[row][col] == symbol {
				if number > 0 {
					numberLength := int(math.Floor(math.Log10(float64(number))) + 1)

					if findAdjacentSymbol(row, col-numberLength, col-1, grid) {
						sum += number
					}

					number = 0
				}

				continue
			}

			number = number*10 + int(grid[row][col]) - 0x30
		}

		// End of the line, check if we still have a number
		if number > 0 {
			numberLength := int(math.Floor(math.Log10(float64(number))) + 1)

			if findAdjacentSymbol(row, len(grid[row])-1-numberLength, len(grid[row])-2, grid) {
				sum += number
			}

			number = 0
		}
	}

	fmt.Println(sum)
}

func findAdjacentSymbol(row, colStart, colEnd int, grid [][]byte) bool {
	start := colStart - 1
	if start < 0 {
		start = 0
	}

	end := colEnd + 1
	if end >= len(grid[row]) {
		end--
	}

	// Check one row above and below
	for col := start; col <= end; col++ {
		if row > 0 && grid[row-1][col] != '.' {
			return true
		}
		if row+1 < len(grid) && grid[row+1][col] != '.' {
			return true
		}
	}

	// Check left and right
	return grid[row][start] == symbol || grid[row][end] == symbol
}
