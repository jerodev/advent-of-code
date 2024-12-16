package main

import (
	"advent-of-code/util"
	"bufio"
	"bytes"
	"fmt"
)

const (
	boundry = '#'
	box     = 'O'
	empty   = '.'
	up      = '^'
	right   = '>'
	down    = 'v'
	left    = '<'
)

var moves = map[byte]position{
	up:    {0, -1},
	right: {1, 0},
	down:  {0, 1},
	left:  {-1, 0},
}

var grid [][]byte
var robot position

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	// Read grid
	var x int
	for scan.Scan() {
		row := scan.Bytes()
		if len(row) == 0 {
			break
		}

		x = bytes.IndexRune(row, '@')
		if x > -1 {
			robot.X = x
			robot.Y = len(grid)
		}

		// Prevent scanner memory override
		gridRow := make([]byte, len(row))
		copy(gridRow, row)
		grid = append(grid, gridRow)
	}

	// Move stuf arround
	for scan.Scan() {
		for _, b := range scan.Bytes() {
			moveObject(robot.X, robot.Y, b, true)
		}
	}

	// Find boxes in our grid
	var result int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == box {
				result += 100*y + x
			}
		}
	}

	fmt.Println(result)
}

type position struct {
	X, Y int
}

func moveObject(x, y int, direction byte, isRobot bool) bool {
	dx := x + moves[direction].X
	dy := y + moves[direction].Y

	if grid[dy][dx] == boundry {
		return false
	}

	if grid[dy][dx] == box && !moveObject(dx, dy, direction, false) {
		return false
	}

	// All is well, move the object
	grid[dy][dx] = grid[y][x]
	grid[y][x] = empty

	if isRobot {
		robot.X = dx
		robot.Y = dy
	}

	return true
}
