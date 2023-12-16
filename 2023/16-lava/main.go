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

	grid = bounce(grid, 0, -1, 'R')
	fmt.Println(calculateEnergy(grid))
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
			if strings.Contains(grid[y][x].EntryHistory, "D") || strings.Contains(grid[y][x].EntryHistory, "R") { // Been there, done that
				return grid
			}
			direction = 'D'
		case 'L':
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "L") {
				return grid
			}
			direction = 'U'
		case 'U':
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "L") {
				return grid
			}
			direction = 'L'
		case 'D':
			if strings.Contains(grid[y][x].EntryHistory, "D") || strings.Contains(grid[y][x].EntryHistory, "R") {
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
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "R") { // Been there, done that
				return grid
			}
			direction = 'U'
		case 'L':
			if strings.Contains(grid[y][x].EntryHistory, "L") || strings.Contains(grid[y][x].EntryHistory, "D") {
				return grid
			}
			direction = 'D'
		case 'U':
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "R") {
				return grid
			}
			direction = 'R'
		case 'D':
			if strings.Contains(grid[y][x].EntryHistory, "L") || strings.Contains(grid[y][x].EntryHistory, "D") {
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
			if strings.Contains(grid[y][x].EntryHistory, "U") || strings.Contains(grid[y][x].EntryHistory, "D") { // Been there, done that
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
			if strings.Contains(grid[y][x].EntryHistory, "L") || strings.Contains(grid[y][x].EntryHistory, "R") { // Been there, done that
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
	var resultGrid [][]byte

	for y := range grid {
		resultGrid = append(resultGrid, []byte{})
		for x := range grid[y] {
			resultGrid[y] = append(resultGrid[y], grid[y][x].Value)
			if grid[y][x].Energized {
				energy++
				resultGrid[y][x] = '#'
			}
		}
	}

	util.PrintMatrix(resultGrid)

	return energy
}
