package main

import (
	"advent-of-code/util"
	"fmt"
	"math"
)

type position struct {
	Value  byte
	OnPath bool
}

type path struct {
	Position     [2]int
	LastPosition [2]int
}

func newPath(position [2]int, lastPosition [2]int) path {
	return path{
		Position:     position,
		LastPosition: lastPosition,
	}
}

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	grid := [][]position{
		{},
	}
	startPosition := [2]int{}
	rowNumber := 0

	for {
		_, err := file.Read(b)
		if err != nil {
			break
		}

		if b[0] == '\n' {
			rowNumber++
			grid = append(grid, []position{})
			continue
		}

		if b[0] == 'S' {
			startPosition = [2]int{
				len(grid[rowNumber]),
				rowNumber,
			}
		}

		grid[rowNumber] = append(grid[rowNumber], position{
			Value:  b[0],
			OnPath: b[0] == 'S',
		})
	}

	grid[startPosition[1]][startPosition[0]].Value = determineStartPipe(grid, startPosition[0], startPosition[1])

	connections := findConnections(grid, startPosition[0], startPosition[1])
	paths := make([]path, 2)
	for i, c := range connections {
		paths[i] = newPath(c, startPosition)
		grid[c[1]][c[0]].OnPath = true
	}

	steps := 1

	for {
		for i := range paths {
			connections = findConnections(grid, paths[i].Position[0], paths[i].Position[1])
			for _, c := range connections {
				if c[0] != paths[i].LastPosition[0] || c[1] != paths[i].LastPosition[1] {
					paths[i].LastPosition = paths[i].Position
					paths[i].Position = c
					grid[c[1]][c[0]].OnPath = true
					break
				}
			}
		}

		steps++

		fmt.Println(paths)

		if paths[0].Position[0] == paths[1].Position[0] && paths[0].Position[1] == paths[1].Position[1] {
			break
		}
	}

	fmt.Println(steps)

	//
	// Part 2: calculate surface
	surface := 0
	for y := 0; y < len(grid); y++ {
		intersections := 0.0
		for x := 0; x < len(grid[y]); x++ {
			if !grid[y][x].OnPath {
				if math.Mod(intersections, 2) != 0 {
					fmt.Println(y, x, intersections)
					surface++
				}
			} else {
				switch grid[y][x].Value {
				case '|':
					intersections++
				case 'J', 'F':
					intersections += 0.5
				case 'L', '7':
					intersections -= 0.5
				}
			}
		}
	}

	fmt.Println(surface)
}

func determineStartPipe(grid [][]position, x, y int) byte {
	if x > 0 && grid[y][x-1].Value == '-' && grid[y][x+1].Value == '-' {
		return '-'
	}

	if y > 0 && grid[y-1][x].Value == '|' && grid[y+1][x].Value == '|' {
		return '|'
	}

	goingLeft := grid[y][x-1].Value == '-' || grid[y][x-1].Value == 'L' || grid[y][x-1].Value == 'F'
	goingRight := !goingLeft && (grid[y][x+1].Value == '-' || grid[y][x+1].Value == 'J' || grid[y][x+1].Value == '7')

	// Bottom up
	if grid[y+1][x].Value == '|' || grid[y+1][x].Value == 'J' || grid[y+1][x].Value == 'L' {
		if goingLeft {
			return 'F'
		}
		if goingRight {
			return '7'
		}
	}

	// Going down
	if goingLeft {
		return 'J'
	}

	return 'L'
}

// findConnections returns all possible destinations from position x, y
func findConnections(grid [][]position, x, y int) [][2]int {
	pos := grid[y][x]
	var connections [][2]int

	// Top
	if y > 0 && (pos.Value == '|' || pos.Value == 'J' || pos.Value == 'L') && (grid[y-1][x].Value == '|' || grid[y-1][x].Value == '7' || grid[y-1][x].Value == 'F') {
		connections = append(connections, [2]int{x, y - 1})
	}

	// Bottom
	if y < len(grid)-1 && (pos.Value == '|' || pos.Value == '7' || pos.Value == 'F') && (grid[y+1][x].Value == '|' || grid[y+1][x].Value == 'J' || grid[y+1][x].Value == 'L') {
		connections = append(connections, [2]int{x, y + 1})
	}

	// Left
	if x > 0 && (pos.Value == '-' || pos.Value == 'J' || pos.Value == '7') && (grid[y][x-1].Value == '-' || grid[y][x-1].Value == 'L' || grid[y][x-1].Value == 'F') {
		connections = append(connections, [2]int{x - 1, y})
	}

	// Right
	if x < len(grid[0])-1 && (pos.Value == '-' || pos.Value == 'L' || pos.Value == 'F') && (grid[y][x+1].Value == '-' || grid[y][x+1].Value == 'J' || grid[y][x+1].Value == '7') {
		connections = append(connections, [2]int{x + 1, y})
	}

	return connections
}
