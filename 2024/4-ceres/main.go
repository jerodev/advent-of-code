package main

import (
	"advent-of-code/util"
	"fmt"
	"strings"
)

const xmas = "XMAS"

var grid []string

func main() {
	grid = strings.Split(util.ReadFileFromArgs(), "\n")

	var xmasCount, xCount int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[x][y] == 'X' {
				xmasCount += findXmas(x, y)
			} else if grid[x][y] == 'A' {
				xCount += findXmasX(x, y)
			}
		}
	}

	fmt.Println(xmasCount)
	fmt.Println(xCount)
}

// findXmas returns the amount of xmas strings originating from this point
func findXmas(x, y int) int {
	var xmasCount, ddx, ddy int

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			for i := 1; i < len(xmas); i++ {
				ddx = x + (dx * i)
				if ddx < 0 || ddx >= len(grid[0]) {
					break
				}

				ddy = y + (dy * i)
				if ddy < 0 || ddy >= len(grid) {
					break
				}

				if grid[ddx][ddy] != xmas[i] {
					break
				}

				if i == len(xmas)-1 {
					xmasCount++
					break
				}
			}
		}
	}

	return xmasCount
}

func findXmasX(x, y int) int {
	if x == 0 || y == 0 || x == len(grid[0])-1 || y == len(grid)-1 {
		return 0
	}

	if !((grid[x-1][y-1] == 'M' && grid[x+1][y+1] == 'S') || (grid[x-1][y-1] == 'S' && grid[x+1][y+1] == 'M')) {
		return 0
	}

	if !((grid[x-1][y+1] == 'S' && grid[x+1][y-1] == 'M') || (grid[x-1][y+1] == 'M' && grid[x+1][y-1] == 'S')) {
		return 0
	}

	return 1
}
