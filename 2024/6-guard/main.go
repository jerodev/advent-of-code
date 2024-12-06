package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"strings"
)

const (
	freeBlock     = '.'
	obstacleBlock = '#'
	visiteBlock   = 'X'
)

var direction = []int{-1, 0, 1, 0}
var grid [][]byte
var guardStart position

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	// Map the grid
	var row string
	var guardIndex int
	for scan.Scan() {
		row = scan.Text()

		guardIndex = strings.Index(row, "^")
		if guardIndex > -1 {
			guardStart = position{guardIndex, len(grid)}
			row = strings.Replace(row, "^", "X", 1)
		}

		grid = append(grid, []byte(row))
	}

	// Walk the grid
	fmt.Println(walkGrid(copyGrid(grid)))
}

type position struct {
	X, Y int
}

func walkGrid(grid [][]byte) int {
	guard := position{guardStart.X, guardStart.Y}

	stepCount := 1
	var nx, ny int
	dx, dy := 1, 0
	for {
		nx, ny = guard.X+direction[dx], guard.Y+direction[dy]

		// Did we leave the grid?
		if nx < 0 || ny < 0 || nx >= len(grid[0]) || ny >= len(grid) {
			return stepCount
		}

		if grid[ny][nx] == obstacleBlock {
			dx, dy = (dx+1)%4, (dy+1)%4
			continue
		} else if grid[ny][nx] == 0 {
			stepCount++
		} else if grid[ny][nx] > 4 {
			return -1 // Infinite loop detected!
		}

		grid[ny][nx]++
		guard.X, guard.Y = nx, ny
	}
}

func copyGrid(grid [][]byte) [][]byte {
	gridCopy := make([][]byte, len(grid))

	for i := range grid {
		barry := make([]byte, len(grid[i]))
		for j := range grid[i] {
			if grid[i][j] == freeBlock {
				barry[j] = 0
			} else {
				barry[j] = grid[i][j]
			}
		}
		gridCopy[i] = barry
	}

	return gridCopy
}
