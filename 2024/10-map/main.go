package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
)

var grid [][]byte

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	var starts []position

	for scan.Scan() {
		row := scan.Bytes()

		for i := range row {
			row[i] = row[i] - '0'

			if row[i] == 0 {
				starts = append(starts, position{i, len(grid)})
			}
		}

		grid = append(grid, row)
	}

	var trials, trialPoints int
	for i := range starts {
		trialheads := walkUp(starts[i])
		trials += len(util.Unique(trialheads))
		trialPoints += len(trialheads)
	}

	fmt.Println(trials)
	fmt.Println(trialPoints)
}

type position struct {
	X, Y int
}

// walkUp starts at a position and returns all sumits it can reach with a unique path
func walkUp(p position) []position {
	height := grid[p.Y][p.X]
	if height == 9 {
		return []position{p}
	}

	var summits []position
	if p.X > 0 && grid[p.Y][p.X-1] == height+1 { // LEFT
		summits = append(summits, walkUp(position{p.X - 1, p.Y})...)
	}
	if p.X < len(grid[0])-1 && grid[p.Y][p.X+1] == height+1 { // RIGHT
		summits = append(summits, walkUp(position{p.X + 1, p.Y})...)
	}
	if p.Y > 0 && grid[p.Y-1][p.X] == height+1 { // UP
		summits = append(summits, walkUp(position{p.X, p.Y - 1})...)
	}
	if p.Y < len(grid)-1 && grid[p.Y+1][p.X] == height+1 { // DOWN
		summits = append(summits, walkUp(position{p.X, p.Y + 1})...)
	}

	return summits
}
