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
			moveObject2(robot2.X, robot2.Y, b, true)

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

func moveObject2(x, y int, direction byte, isRobot bool) bool {
	dx := x + moves[direction].X
	dy := y + moves[direction].Y

	if grid2[dy][dx] == boundry {
		return false
	}

	// Always start from the left box part
	bx := dx
	if grid2[dy][dx] == boxRight {
		bx -= 1
	}

	if grid2[dy][bx] == boxLeft {
		switch direction {
		case up:
			if (grid2[dy-1][bx] == empty && grid2[dy-1][bx+1] == empty) || (grid2[dy-1][bx] == boxRight && grid2[dy-1][bx+1] == boxLeft && can2BoxesMove(bx-1, y-1, direction) && moveObject2(bx, dy, direction, false) && moveObject2(bx+1, dy, direction, false)) || (grid2[dy-1][bx] == boxRight && grid2[dy-1][bx+1] == empty && moveObject2(bx, dy, direction, false)) || (grid2[dy-1][bx+1] == left && grid2[dy-1][bx] == empty && moveObject2(bx, dy, direction, false)) {
				grid2[dy][bx] = empty
				grid2[dy][bx+1] = empty
				grid2[dy-1][bx] = boxLeft
				grid2[dy-1][bx+1] = boxRight
			} else {
				return false
			}
		case right:
			if grid2[dy][bx+2] == empty || (grid2[dy][bx+2] == boxLeft && moveObject2(bx+1, dy, direction, false)) {
				grid2[dy][bx] = empty
				grid2[dy][bx+1] = boxLeft
				grid2[dy][bx+2] = boxRight
			} else {
				return false
			}
		case down:
			if (grid2[dy+1][bx] == empty && grid2[dy+1][bx+1] == empty) || (grid2[dy+1][bx] == boxLeft && moveObject2(bx, dy, direction, false)) || (grid2[dy+1][bx] == boxRight && grid2[dy+1][bx+1] == boxLeft && can2BoxesMove(bx-2, dy+1, direction) && moveObject2(bx, dy, direction, false) && moveObject2(bx+1, dy, direction, false)) || (grid2[dy+1][bx] == boxRight && moveObject2(bx, dy, direction, false) && grid2[dy+1][bx+1] == empty) {
				grid2[dy][bx] = empty
				grid2[dy][bx+1] = empty
				grid2[dy+1][bx] = boxLeft
				grid2[dy+1][bx+1] = boxRight
			} else {
				return false
			}
		case left:
			if grid2[dy][bx-1] == empty || (grid2[dy][bx-1] == boxRight && moveObject2(bx-1, dy, direction, false)) {
				grid2[dy][bx+1] = empty
				grid2[dy][bx] = boxRight
				grid2[dy][bx-1] = boxLeft
			} else {
				return false
			}
		}
	}

	// All is well, move the object
	if isRobot {
		grid2[dy][dx] = grid2[y][x]
		grid2[y][x] = empty

		robot2.X = dx
		robot2.Y = dy
	}

	return true
}

func can2BoxesMove(x, y int, direction byte) bool {
	if direction == down {
		y += 2
	} else {
		y -= 2
	}

	if y < 0 || y > len(grid)-1 || x > len(grid2[y])-4 {
		return true
	}

	// Obstacles in the way?
	for i := 0; i < 4; i++ {
		if grid2[y][x+i] == boundry {
			return false
		}

	}

	// One+ box blocking?
	if x > len(grid2[y])-5 {
		for i := -1; i < 5; i++ {
			if grid[y][x+i] == boxLeft && grid[y][x+i+1] == boxRight && !moveObject2(x+i, y, direction, false) {
				return false
			}
		}
	}

	return true
}
