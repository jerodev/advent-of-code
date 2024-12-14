package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	gridWidth  = 101
	gridHeight = 103
	gridMidX   = int(gridWidth / 2)
	gridMidY   = int(gridHeight / 2)
	iterations = 100
)

var inputRegex *regexp.Regexp

func main() {
	inputRegex = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	var robots []robot
	var matches []string
	for scan.Scan() {
		matches = inputRegex.FindStringSubmatch(scan.Text())

		robots = append(robots, robot{
			p: position{
				util.MustAtoi(matches[1]),
				util.MustAtoi(matches[2]),
			},
			v: position{
				util.MustAtoi(matches[3]),
				util.MustAtoi(matches[4]),
			},
		})
	}

	// A grid for displaying easter eggs
	grid := make([][]byte, gridHeight)
	for g := range gridHeight {
		grid[g] = make([]byte, gridWidth)
	}

	// Move the robots
	//for i := 0; i >= 0; i++ { // Part 2
	for i := 0; i < iterations; i++ {
		for r := range robots {
			grid[robots[r].p.Y][robots[r].p.X] = ' '

			robots[r].p.X += robots[r].v.X
			robots[r].p.Y += robots[r].v.Y

			robots[r].p.X = robots[r].p.X % gridWidth
			if robots[r].p.X < 0 {
				robots[r].p.X += gridWidth
			}

			robots[r].p.Y = robots[r].p.Y % gridHeight
			if robots[r].p.Y < 0 {
				robots[r].p.Y += gridHeight
			}

			grid[robots[r].p.Y][robots[r].p.X] = '#'
		}

		printPotentialTree(i+1, grid)
	}

	// 0 | 1
	// - + -
	// 2 | 3
	var quadrants [4]int

	// Calculate quadrants
	for _, r := range robots {
		if r.p.X < gridMidX {
			if r.p.Y < gridMidY {
				quadrants[0]++
			} else if r.p.Y > gridMidY {
				quadrants[2]++
			}
		} else if r.p.X > gridMidX {
			if r.p.Y < gridMidY {
				quadrants[1]++
			} else if r.p.Y > gridMidY {
				quadrants[3]++
			}
		}
	}

	fmt.Println(quadrants)
	fmt.Println(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
}

type position struct {
	X, Y int
}

type robot struct {
	p position
	v position
}

func printPotentialTree(iteration int, grid [][]byte) {
	// Search the grid for a continuous line of 10 #
	var found bool

outer:
	for y := range gridHeight {
		var count int
		for x := range gridWidth {
			if grid[y][x] == '#' {
				count++
			} else {
				count = 0
			}

			if count == 20 {
				found = true
				break outer
			}
		}
	}

	if found {
		util.PrintMatrix(grid)
		fmt.Println(iteration)

		b := make([]byte, 1)
		_, _ = os.Stdin.Read(b)
	}
}
