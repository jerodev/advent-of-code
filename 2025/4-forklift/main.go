package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var x, y int
	var l string

	lines := strings.Split(input, "\n")
	grid := make([][]bool, len(lines))
	for y, l = range lines {
		for x = range l {
			grid[y] = append(grid[y], l[x] == '@')
		}
	}

	// Part 1, remove rolls once
	var count int
	grid, count = removeRolls(grid)
	fmt.Println(count)

	// Part 2, remove rolls until no more rolls can be removed
	n := count
	for n > 0 {
		grid, n = removeRolls(grid)
		count += n
	}
	fmt.Println(count)
}

func removeRolls(grid [][]bool) ([][]bool, int) {
	var count, n, x, y int
	grid2 := make([][]bool, len(grid))

	for y = range grid {
		grid2[y] = append(grid2[y], grid[y]...)

		for x = range grid[y] {
			if !grid[y][x] {
				continue
			}

			n = 0

			// Row above & row below
			for i := max(0, x-1); i <= min(len(grid[y])-1, x+1); i++ {
				if y > 0 && grid[y-1][i] {
					n++
				}
				if y < len(grid)-1 && grid[y+1][i] {
					n++
				}
			}
			if x > 0 && grid[y][x-1] {
				n++
			}
			if x < len(grid[y])-1 && grid[y][x+1] {
				n++
			}

			if n < 4 {
				count++
				grid2[y][x] = false
			}
		}
	}

	return grid2, count
}
