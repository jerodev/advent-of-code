package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strings"
)

const (
	TILE_EMPTY    = '.'
	TILE_SPLIT_H  = '-'
	TILE_SPLIT_V  = '|'
	TILE_MIRROR_L = '\\'
	TILE_MIRROR_R = '/'
)

type tile struct {
	Value        byte
	Energized    bool
	EntryHistory string
}

func main() {
	file := util.FileFromArgs()

	grid := [][]tile{
		{},
	}
	rowNumber := 0

	b := make([]byte, 1)

	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}

		if b[0] == '\n' {
			grid = append(grid, []tile{})
			rowNumber++
			continue
		}

		grid[rowNumber] = append(grid[rowNumber], tile{
			Value:        b[0],
			Energized:    false,
			EntryHistory: "",
		})
	}

	// Part 1
	fmt.Println(calculateEnergy(bounce(copyGrid(grid), 0, -1, 'R')))

	// Part 2
	maxEnergy := 0
	maxPos := []int{0, 0}
	e := 0
	for y := range grid {
		e = calculateEnergy(bounce(copyGrid(grid), y, -1, 'R'))
		if e > maxEnergy {
			maxPos = []int{y, -1}
			maxEnergy = e
		}

		e = calculateEnergy(bounce(copyGrid(grid), y, len(grid[y]), 'L'))
		if e > maxEnergy {
			maxPos = []int{y, len(grid[y])}
			maxEnergy = e
		}
	}
	for x := range grid[0] {
		e = calculateEnergy(bounce(copyGrid(grid), -1, x, 'D'))
		if e > maxEnergy {
			maxPos = []int{-1, x}
			maxEnergy = e
		}

		e = calculateEnergy(bounce(copyGrid(grid), len(grid), x, 'U'))
		if e > maxEnergy {
			maxPos = []int{len(grid), x}
			maxEnergy = e
		}
	}

	fmt.Println(maxPos)
	fmt.Println(maxEnergy)
}

func bounce(grid [][]tile, y, x int, direction byte) [][]tile {
	switch direction {
	case 'R':
		if x+1 < len(grid[y]) {
			x++
		} else {
			return grid
		}
	case 'L':
		if x > 0 {
			x--
		} else {
			return grid
		}
	case 'U':
		if y > 0 {
			y--
		} else {
			return grid
		}
	case 'D':
		if y+1 < len(grid) {
			y++
		} else {
			return grid
		}
	}

	grid[y][x].Energized = true

	// Continue forward if empty
	if grid[y][x].Value == TILE_EMPTY {
		return bounce(grid, y, x, direction)
	}

	// Change direction if mirror
	if grid[y][x].Value == TILE_MIRROR_L {
		tmp := string(direction)
		switch direction {
		case 'R':
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "R") { // Been there, done that
				return grid
			}
			direction = 'D'
		case 'L':
			if strings.Contains(grid[y][x].EntryHistory, "D") || strings.Contains(grid[y][x].EntryHistory, "L") {
				return grid
			}
			direction = 'U'
		case 'U':
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "R") {
				return grid
			}
			direction = 'L'
		case 'D':
			if strings.Contains(grid[y][x].EntryHistory, "D") || strings.Contains(grid[y][x].EntryHistory, "L") {
				return grid
			}
			direction = 'R'
		}

		grid[y][x].EntryHistory += tmp

		return bounce(grid, y, x, direction)
	} else if grid[y][x].Value == TILE_MIRROR_R {
		tmp := string(direction)
		switch direction {
		case 'R':
			if strings.Contains(grid[y][x].EntryHistory, "D") || strings.Contains(grid[y][x].EntryHistory, "R") { // Been there, done that
				return grid
			}
			direction = 'U'
		case 'L':
			if strings.Contains(grid[y][x].EntryHistory, "L") || strings.Contains(grid[y][x].EntryHistory, "U") {
				return grid
			}
			direction = 'D'
		case 'U':
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "L") {
				return grid
			}
			direction = 'R'
		case 'D':
			if strings.Contains(grid[y][x].EntryHistory, "R") || strings.Contains(grid[y][x].EntryHistory, "D") {
				return grid
			}
			direction = 'L'
		}

		grid[y][x].EntryHistory += tmp

		return bounce(grid, y, x, direction)
	}

	// Split if we hit a splitter in the middle
	if grid[y][x].Value == TILE_SPLIT_H {
		if direction == 'U' || direction == 'D' {
			if grid[y][x].EntryHistory != "" { // Passing once here always has the same result
				return grid
			}

			grid[y][x].EntryHistory += string(direction)

			grid = bounce(grid, y, x, 'L')
			grid = bounce(grid, y, x, 'R')
		} else {
			return bounce(grid, y, x, direction)
		}
	} else if grid[y][x].Value == TILE_SPLIT_V {
		if direction == 'L' || direction == 'R' {
			if grid[y][x].EntryHistory != "" { // Passing once here always has the same result
				return grid
			}

			grid[y][x].EntryHistory += string(direction)

			grid = bounce(grid, y, x, 'U')
			grid = bounce(grid, y, x, 'D')
		} else {
			return bounce(grid, y, x, direction)
		}
	}

	return grid
}

func calculateEnergy(grid [][]tile) int {
	energy := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x].Energized {
				energy++
			}
		}
	}

	return energy
}

func copyGrid(grid [][]tile) [][]tile {
	newGrid := make([][]tile, len(grid))
	for y := range grid {
		newGrid[y] = make([]tile, len(grid[y]))
		copy(newGrid[y], grid[y])
	}

	return newGrid
}
