package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
)

const (
	boundry  = '#'
	box      = 'O'
	boxLeft  = '['
	boxRight = ']'
	empty    = '.'

	up    = '^'
	right = '>'
	down  = 'v'
	left  = '<'
)

var moves = map[byte]position{
	up:    {0, -1},
	right: {1, 0},
	down:  {0, 1},
	left:  {-1, 0},
}

var grid, grid2 [][]byte
var robot, robot2 position

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	// Read grid
	for scan.Scan() {
		row := scan.Bytes()
		if len(row) == 0 {
			break
		}

		// Part 1: Prevent scanner memory override
		gridRow := make([]byte, len(row))
		copy(gridRow, row)
		grid = append(grid, gridRow)

		// Part 2
		gridRow = make([]byte, 0, len(row)*2)
		for x, b := range row {
			switch b {
			case '@':
				robot.X = x
				robot.Y = len(grid) - 1
				robot2.X = 2 * x
				robot2.Y = len(grid2)
				gridRow = append(gridRow, '@', empty)
			case boundry:
				gridRow = append(gridRow, boundry, boundry)
			case box:
				gridRow = append(gridRow, boxLeft, boxRight)
			case empty:
				gridRow = append(gridRow, empty, empty)
			}
		}
		grid2 = append(grid2, gridRow)
	}

	util.PrintMatrix(grid2)

	// Move stuf arround
	for scan.Scan() {
		for _, b := range scan.Bytes() {
			moveObject(robot.X, robot.Y, b, true)
			moveRobot2(robot2.X, robot2.Y, b)

			// util.PrintMatrix(grid2)
		}
	}

	// Find boxes in our grid
	var result, result2 int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == box {
				result += 100*y + x
			}
		}
	}
	for y := range grid2 {
		for x := range grid2[y] {
			if grid2[y][x] == boxLeft {
				result2 += 100*y + x
			}
		}
	}

	fmt.Println(result)
	fmt.Println(result2)
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

func moveRobot2(x, y int, direction byte) {
	dx := x + moves[direction].X
	dy := y + moves[direction].Y

	if grid2[dy][dx] == boundry {
		return
	}

	// Always start from the left box part
	if grid2[dy][dx] == boxRight || grid2[dy][dx] == boxLeft {
		if !canBoxMove(dx, dy, direction) {
			return
		}

		moveBox(dx, dy, direction)
	}

	grid2[y][x] = empty
	robot2.X = dx
	robot2.Y = dy
	grid2[dy][dx] = '@'
}

func canBoxMove(x, y int, direction byte) bool {
	// Always asume the left of the box
	if grid2[y][x] == boxRight {
		x--
	}

	dx := x + moves[direction].X
	dy := y + moves[direction].Y

	if grid2[dy][dx] == boundry {
		return false
	}

	if direction == up || direction == down {
		if grid2[dy][dx+1] == boundry {
			return false
		}
		if grid2[dy][dx] == boxLeft || grid2[dy][dx] == boxRight {
			if !canBoxMove(dx, dy, direction) {
				return false
			}
		}
		if grid2[dy][dx+1] == boxLeft {
			if !canBoxMove(dx+1, dy, direction) {
				return false
			}
		}
	} else if direction == left {
		if grid2[dy][dx] == boxRight {
			return canBoxMove(dx-1, dy, direction)
		}
	} else {
		if grid2[dy][dx+1] == boundry {
			return false
		}
		if grid2[dy][dx+1] == boxLeft {
			return canBoxMove(dx+1, dy, direction)
		}
	}

	return true
}

func moveBox(x, y int, direction byte) {
	// Always asume the left of the box
	if grid2[y][x] == boxRight {
		x--
	}

	dx := x + moves[direction].X
	dy := y + moves[direction].Y

	if direction == right {
		if grid2[dy][dx+1] == boxLeft {
			moveBox(dx+1, dy, direction)
		}
	} else if grid2[dy][dx] == boxLeft || grid2[dy][dx] == boxRight {
		moveBox(dx, dy, direction)
	}

	if direction == up || direction == down {
		if grid2[dy][dx+1] == boxLeft || grid2[dy][dx+1] == boxRight {
			moveBox(dx+1, dy, direction)
		}
	}

	grid2[y][x] = empty
	grid2[y][x+1] = empty
	grid2[dy][dx] = boxLeft
	grid2[dy][dx+1] = boxRight
}
