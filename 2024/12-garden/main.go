package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
)

var grid [][]byte
var history map[byte][]position

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		grid = append(grid, scan.Bytes())
	}

	history = map[byte][]position{}
	var area, perimeter, sum int
	for y := range grid {
		for x := range grid[y] {
			area, perimeter = measure(x, y)
			sum += area * perimeter

			// Debug: print the found area
			if area > 0 {
				fmt.Println(string(grid[y][x]), area, perimeter)
			}
		}
	}

	fmt.Println(sum)
}

type position struct {
	X, Y int
}

// measure returns the area and peremiter this position is part of
func measure(x, y int) (int, int) {
	// Have we passed here yet?
	if _, exists := history[grid[y][x]]; exists {
		for i := range history[grid[y][x]] {
			if history[grid[y][x]][i].X == x && history[grid[y][x]][i].Y == y {
				return 0, 0
			}
		}

		history[grid[y][x]] = append(history[grid[y][x]], position{x, y})
	} else {
		history[grid[y][x]] = []position{
			{x, y},
		}
	}

	// Measure shapes
	area, perimeter := 1, 0
	var dA, dP int

	if x > 0 && grid[y][x] == grid[y][x-1] {
		dA, dP = measure(x-1, y)
		area += dA
		perimeter += dP
	} else {
		perimeter++
	}

	if x < len(grid[0])-1 && grid[y][x] == grid[y][x+1] {
		dA, dP = measure(x+1, y)
		area += dA
		perimeter += dP
	} else {
		perimeter++
	}

	if y > 0 && grid[y][x] == grid[y-1][x] {
		dA, dP = measure(x, y-1)
		area += dA
		perimeter += dP
	} else {
		perimeter++
	}

	if y < len(grid)-1 && grid[y][x] == grid[y+1][x] {
		dA, dP = measure(x, y+1)
		area += dA
		perimeter += dP
	} else {
		perimeter++
	}

	return area, perimeter
}
