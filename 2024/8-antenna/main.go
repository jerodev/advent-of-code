package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"math"
	"slices"
)

const nothing = '.'

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	antennas := map[byte][]position{}

	var grid [][]byte
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

	var antinodes, points []position
	for k := range antennas {
		for i := range antennas[k] {
			for j := range antennas[k] {
				if i == j {
					continue
				}

				points = calculateAntinodes(antennas[k][i], antennas[k][j])
				for p := range points {
					if slices.IndexFunc(antinodes, func(pp position) bool { return points[p].eq(pp) }) == -1 {
						antinodes = append(antinodes, points[p])
					}
				}

			}
		}
	}

	fmt.Println(len(antinodes))
}

type position struct {
	X, Y int
}

func (p position) eq(q position) bool {
	return p.X == q.X && p.Y == q.Y
}

// calculateAntinodes will calculate all antinodes between two antennas
func calculateAntinodes(c1, c2 position) []position {
	var positions []position

	// sqrt((x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2));
	d := math.Sqrt(float64((c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y)))

	for r := range 20 {
		if d > float64(r+2*r) {
			continue // Circles don't touch
		}

		// TODO: calculate touch points
	}

	return positions
}
