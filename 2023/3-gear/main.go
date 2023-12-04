package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
)

const (
	gearSymbol = '*'
	symbol     = '+'
)

type gear struct {
	row, col int
	numbers  []int
}

func newGear(row, col int) *gear {
	return &gear{row, col, []int{}}
}

func (c gear) value() int {
	if len(c.numbers) == 2 {
		return c.numbers[0] * c.numbers[1]
	}

	return 0
}

func main() {
	file := util.FileFromArgs()

	gears := []*gear{}

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
			if b[0] == gearSymbol {
				gears = append(gears, newGear(len(grid), len(line)))
			}

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
						findAdjacentGears(number, row, col-numberLength, col-1, grid, gears)
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

			if findAdjacentSymbol(row, len(grid[row])-numberLength, len(grid[row])-1, grid) {
				sum += number
				findAdjacentGears(number, row, len(grid[row])-numberLength, len(grid[row])-1, grid, gears)
			}
		}
	}

	fmt.Println(sum)

	// Calculate gears sum
	gearsSum := 0
	for _, g := range gears {
		gearsSum += g.value()
	}
	fmt.Println(gearsSum)
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

	// Check one row above and below for symbols
	for col := start; col <= end; col++ {
		if row > 0 && grid[row-1][col] == symbol {
			return true
		}
		if row+1 < len(grid) && grid[row+1][col] == symbol {
			return true
		}
	}

	// Check left and right for symbols
	return grid[row][start] == symbol || grid[row][end] == symbol
}

func findAdjacentGears(number, row, colStart, colEnd int, grid [][]byte, gears []*gear) {
	start := colStart - 1
	if start < 0 {
		start = 0
	}

	end := colEnd + 1
	if end >= len(grid[row]) {
		end--
	}

	// Check one row above and below for symbols
	for col := start; col <= end; col++ {
		if row > 0 && grid[row-1][col] == symbol {
			gear := findGearByCoordinates(gears, row-1, col)
			if gear != nil {
				gear.numbers = append(gear.numbers, number)
			}
		}
		if row+1 < len(grid) && grid[row+1][col] == symbol {
			gear := findGearByCoordinates(gears, row+1, col)
			if gear != nil {
				gear.numbers = append(gear.numbers, number)
			}
		}
	}

	// Check left and right for symbols
	if grid[row][start] == symbol {
		gear := findGearByCoordinates(gears, row, start)
		if gear != nil {
			gear.numbers = append(gear.numbers, number)
		}
	}
	if grid[row][end] == symbol {
		gear := findGearByCoordinates(gears, row, end)
		if gear != nil {
			gear.numbers = append(gear.numbers, number)
		}
	}
}

func findGearByCoordinates(gears []*gear, row, col int) *gear {
	for _, g := range gears {
		if g.row == row && g.col == col {
			return g
		}
	}

	return nil
}
