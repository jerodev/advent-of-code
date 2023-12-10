package main

import (
	"advent-of-code/util"
	"fmt"
)

type path struct {
	Position     [2]int
	LastPosition [2]int
}

func newPath(position [2]int) path {
	return path{
		Position:     position,
		LastPosition: [2]int{-1, -1},
	}
}

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	grid := [][]byte{
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
			grid = append(grid, []byte{})
			continue
		}

		if b[0] == 'S' {
			startPosition = [2]int{
				len(grid[rowNumber]),
				rowNumber,
			}
		}

		grid[rowNumber] = append(grid[rowNumber], b[0])
	}

	connections := findConnections(grid, startPosition[0], startPosition[1])
	paths := make([]path, 2)
	for i, c := range connections {
		paths[i] = newPath(c)
	}

	fmt.Println(paths)
	steps := 1

	for {
		for i := range paths {
			connections = findConnections(grid, paths[i].Position[0], paths[i].Position[1])
			for _, c := range connections {
				if c[0] != paths[i].LastPosition[0] || c[1] != paths[i].LastPosition[1] {
					paths[i].LastPosition = paths[i].Position
					paths[i].Position = c
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
}

// findConnections returns all possible destinations from position x, y
func findConnections(grid [][]byte, x, y int) [][2]int {
	pos := grid[y][x]
	var connections [][2]int

	// Top
	if y > 0 && (pos == '|' || pos == 'J' || pos == 'L' || pos == 'S') && (grid[y-1][x] == '|' || grid[y-1][x] == '7' || grid[y-1][x] == 'F') {
		connections = append(connections, [2]int{x, y - 1})
	}

	// Bottom
	if y < len(grid[0])-1 && (pos == '|' || pos == '7' || pos == 'F' || pos == 'S') && (grid[y+1][x] == '|' || grid[y+1][x] == 'J' || grid[y+1][x] == 'L') {
		connections = append(connections, [2]int{x, y + 1})
	}

	// Left
	if x > 0 && (pos == '-' || pos == 'J' || pos == '7' || pos == 'S') && (grid[y][x-1] == '-' || grid[y][x-1] == 'L' || grid[y][x-1] == 'F') {
		connections = append(connections, [2]int{x - 1, y})
	}

	// Right
	if x < len(grid)-1 && (pos == '-' || pos == 'L' || pos == 'F' || pos == 'S') && (grid[y][x+1] == '-' || grid[y][x+1] == 'J' || grid[y][x+1] == '7') {
		connections = append(connections, [2]int{x + 1, y})
	}

	return connections
}
