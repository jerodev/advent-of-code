package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"math"
)

const nothing = '.'

var grid [][]byte
var antennas map[byte][]position

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	antennas = map[byte][]position{}

	var row []byte
	for scan.Scan() {
		row = scan.Bytes()
		for i := range row {
			if row[i] != nothing {
				if _, ok := antennas[row[i]]; !ok {
					antennas[row[i]] = []position{}
				}

				antennas[row[i]] = append(antennas[row[i]], position{i, len(grid)})
			}
		}

		grid = append(grid, row)
	}

	var antinodeCount, antinodeHarmonicCount int
	for y := range len(grid) {
		for x := range len(grid[0]) {
			if isAntinode(x, y, false) {
				antinodeCount++
				antinodeHarmonicCount++
			} else if isAntinode(x, y, true) {
				antinodeHarmonicCount++
			}
		}
	}

	fmt.Println(antinodeCount)
	fmt.Println(antinodeHarmonicCount)
}

type position struct {
	X, Y int
}

func (p position) eq(q position) bool {
	return p.X == q.X && p.Y == q.Y
}

func isAntinode(x, y int, harmonics bool) bool {
	var diff float64

	for k := range antennas {
		for i := range antennas[k] {
			for j := range antennas[k] {
				if i == j {
					continue
				}

				// Are towers on double distances?
				if !harmonics {
					diff = math.Sqrt(float64(util.IntPow(x-antennas[k][i].X, 2)+util.IntPow(y-antennas[k][i].Y, 2))) /
						math.Sqrt(float64(util.IntPow(x-antennas[k][j].X, 2)+util.IntPow(y-antennas[k][j].Y, 2)))
				}
				if harmonics || diff == 2 || diff == .5 {

					// Test if the triangle create by these points has surface 0, this means they are all on a single line
					if x*(antennas[k][i].Y-antennas[k][j].Y)+antennas[k][i].X*(antennas[k][j].Y-y)+antennas[k][j].X*(y-antennas[k][i].Y) == 0 {
						return true
					}
				}
			}
		}
	}

	return false
}
