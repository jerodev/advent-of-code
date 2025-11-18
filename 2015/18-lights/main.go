package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// Read input
	var grid, grid2 [][]bool
	var line string
	var x, y int

	for y, line = range strings.Split(input, "\n") {
		grid = append(grid, make([]bool, len(line)))
		grid2 = append(grid2, make([]bool, len(line)))

		for x = range line {
			if line[x] == '#' {
				grid[y][x] = true
				grid2[y][x] = true
			}
		}
	}

	for range 100 {
		grid = advance(grid, false)
		grid2 = advance(grid2, true)
	}

	// Force the corners before the new count
	grid2[0][0] = true
	grid2[0][len(grid[0])-1] = true
	grid2[len(grid)-1][0] = true
	grid2[len(grid)-1][len(grid[0])-1] = true

	util.PrintBoolMatrix(grid)

	// Count lights that are on
	var count, count2 int
	for y = range grid {
		for x = range grid[y] {
			if grid[y][x] {
				count++
			}
			if grid2[y][x] {
				count2++
			}
		}
	}
	fmt.Println(count)
	fmt.Println(count2)
}

func advance(grid [][]bool, forceCorners bool) [][]bool {
	newGrid := make([][]bool, len(grid))
	var n int

	if forceCorners {
		grid[0][0] = true
		grid[0][len(grid[0])-1] = true
		grid[len(grid)-1][0] = true
		grid[len(grid)-1][len(grid[0])-1] = true
	}

	for y := range grid {
		newGrid[y] = make([]bool, len(grid[y]))

		for x := range grid[y] {
			n = countNeighbours(grid, y, x)

			if grid[y][x] {
				if n != 2 && n != 3 {
					newGrid[y][x] = false
				} else {
					newGrid[y][x] = true
				}
			} else if !grid[y][x] && n == 3 {
				newGrid[y][x] = true
			}
		}
	}

	return newGrid
}

// countNeighbours counts how many neighbours to these coordinatest are on
// If 4 neighbours are counted we return earlier, more neighbours don't matter
func countNeighbours(grid [][]bool, y, x int) int {
	var count, i int

	// Row above
	if y > 0 {
		for i = -1; i < 2; i++ {
			if x+i < 0 || x+i >= len(grid) {
				continue
			}

			if grid[y-1][x+i] {
				count++
			}
		}
	}

	// Row bellow
	if y < len(grid)-1 {
		for i = -1; i < 2; i++ {
			if x+i < 0 || x+i >= len(grid) {
				continue
			}

			if grid[y+1][x+i] {
				count++
				if count > 3 {
					return count
				}
			}
		}
	}

	// Left
	if x > 0 && grid[y][x-1] {
		count++
		if count > 3 {
			return count
		}
	}

	// Right
	if x < len(grid)-1 && grid[y][x+1] {
		count++
	}

	return count
}
