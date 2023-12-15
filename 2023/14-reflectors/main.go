package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
)

const (
	BOULDER = 'O'
	EMPTY   = '.'
	ROCK    = '#'
)

type cycle struct {
	index    int
	grid     [][]byte
	boulders [][2]int
}

func main() {
	file := util.FileFromArgs()

	grid := [][]byte{
		{},
	}
	rowNumber := 0
	b := make([]byte, 1)
	boulders := [][2]int{}

	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		if b[0] == '\n' {
			rowNumber++
			grid = append(grid, []byte{})

			continue
		}

		if b[0] == BOULDER {
			boulders = append(boulders, [2]int{rowNumber, len(grid[rowNumber])})
		}

		grid[rowNumber] = append(grid[rowNumber], b[0])
	}

	// Part 1
	// part1Grid, part1Boulders := rollBoulders(grid, boulders)
	// fmt.Println(calculateLoad(part1Grid, part1Boulders))

	// // Part 2
	cycles := map[string]cycle{}
	var hash string
	for n := 0; n < 1e9; n++ {
		grid, _ = rollBoulders(grid, boulders)
		grid, boulders, _ = turnClockWise(grid)
		grid, _ = rollBoulders(grid, boulders)
		grid, boulders, _ = turnClockWise(grid)
		grid, _ = rollBoulders(grid, boulders)
		grid, boulders, _ = turnClockWise(grid)
		grid, _ = rollBoulders(grid, boulders)
		grid, boulders, hash = turnClockWise(grid)

		if _, ok := cycles[hash]; ok {
			fmt.Println("Found repetition at", n)

			rest := math.Mod(1e9, float64(n))

			fmt.Println(rest)
			for _, c := range cycles {
				if int(rest) == c.index {
					grid = c.grid
					boulders = c.boulders
					break
				}
			}

			break
		}

		cycles[hash] = cycle{
			index:    n,
			grid:     grid,
			boulders: boulders,
		}
	}

	fmt.Println(calculateLoad(grid, boulders))
}

func rollBoulders(grid [][]byte, boulders [][2]int) ([][]byte, [][2]int) {
	var newBoulders [][2]int

	// Loop the boulders and move them up as far as possible
	for _, b := range boulders {
		y, x := b[0], b[1]

		// Roll this boulder!
		var dy int
		for dy = y - 1; dy >= 0; dy-- {
			if grid[dy][x] != EMPTY {
				break
			}
		}

		grid[y][x] = EMPTY
		grid[dy+1][x] = BOULDER
		newBoulders = append(newBoulders, [2]int{dy + 1, x})
	}

	return grid, newBoulders
}

func calculateLoad(grid [][]byte, boulders [][2]int) int {
	rowCount := len(grid)
	load := 0

	for _, b := range boulders {
		load += rowCount - b[0]
	}

	return load
}

func turnClockWise(grid [][]byte) ([][]byte, [][2]int, string) {
	colCount := len(grid[0])
	hash := ""

	var boulders [][2]int
	for y := 0; y < len(grid)/2; y++ {
		for x := 0; x < len(grid)/2; x++ {
			temp := grid[y][x]
			grid[y][x] = grid[colCount-1-x][y]
			grid[colCount-1-x][y] = grid[colCount-1-y][colCount-1-x]
			grid[colCount-1-y][colCount-1-x] = grid[x][colCount-1-y]
			grid[x][colCount-y-1] = temp

			if grid[y][x] == BOULDER {
				boulders = append(boulders, [2]int{y, x})
			}
			if grid[colCount-1-x][y] == BOULDER {
				boulders = append(boulders, [2]int{colCount - 1 - x, y})
			}
			if grid[colCount-1-y][colCount-1-x] == BOULDER {
				boulders = append(boulders, [2]int{colCount - 1 - y, colCount - 1 - x})
			}
			if grid[x][colCount-1-y] == BOULDER {
				boulders = append(boulders, [2]int{x, colCount - 1 - y})
			}

			hash += string(grid[y][x]) + string(grid[colCount-1-x][y]) + string(grid[colCount-1-y][colCount-1-x]) + string(grid[x][colCount-1-y])
		}
	}

	return grid, boulders, hash
}
