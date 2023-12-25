package main

import (
	"advent-of-code/util"
	"fmt"
	"image"
	"io"
	"slices"
)

const (
	PATH    = '.'
	FOREST  = '#'
	SLOPE_U = '^'
	SLOPE_R = '>'
	SLOPE_D = 'v'
	SLOPE_L = '<'
)

type step struct {
	X, Y      int
	direction byte
}

func main() {
	file := util.FileFromArgs()

	grid := [][]byte{
		{},
	}
	rowNumber := 0

	b := make([]byte, 1)

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

		grid[rowNumber] = append(grid[rowNumber], b[0])
	}

	fmt.Println(walk(grid, step{1, 1, SLOPE_D}, []string{"1,0"}, &map[string]int{}, false))
	fmt.Println(walk(grid, step{1, 1, SLOPE_D}, []string{"1,0"}, &map[string]int{}, true))
}

func walk(grid [][]byte, start step, pastPositions []string, walkCache *map[string]int, ignoreSlopes bool) int {
	destination := image.Point{
		X: len(grid[0]) - 2,
		Y: len(grid) - 1,
	}
	steps := 1

	var lastPosition image.Point
	var newPosition step

	switch start.direction {
	case SLOPE_U:
		lastPosition = image.Point{
			X: start.X,
			Y: start.Y + 1,
		}
	case SLOPE_R:
		lastPosition = image.Point{
			X: start.X - 1,
			Y: start.Y,
		}
	case SLOPE_D:
		lastPosition = image.Point{
			X: start.X,
			Y: start.Y - 1,
		}
	case SLOPE_L:
		lastPosition = image.Point{
			X: start.X + 1,
			Y: start.Y,
		}
	}

	for {
		pastPositions = append(pastPositions, fmt.Sprintf("%d,%d", start.X, start.Y))

		var possiblePositions []step
		for _, d := range []step{{-1, 0, SLOPE_L}, {0, -1, SLOPE_U}, {1, 0, SLOPE_R}, {0, 1, SLOPE_D}} {
			// We don't walk in forests
			if grid[start.Y+d.Y][start.X+d.X] == FOREST {
				continue
			}

			// We don't walk up slopes... Or do we?
			if !ignoreSlopes && grid[start.Y+d.Y][start.X+d.X] != PATH && grid[start.Y+d.Y][start.X+d.X] != d.direction {
				continue
			}

			// We don't walk the same path twice
			if slices.Contains(pastPositions, fmt.Sprintf("%d,%d", start.X+d.X, start.Y+d.Y)) {
				continue
			}

			newPosition = step{
				X:         start.X + d.X,
				Y:         start.Y + d.Y,
				direction: d.direction,
			}

			// We don't walk back
			if newPosition.X == lastPosition.X && newPosition.Y == lastPosition.Y {
				continue
			}

			// We can only take a slope in one direction
			possiblePositions = append(possiblePositions, newPosition)
		}

		if len(possiblePositions) == 0 {
			return 0
		}

		// If only one direction, keep walking
		if len(possiblePositions) == 1 {
			if possiblePositions[0].X == destination.X && possiblePositions[0].Y == destination.Y {
				return steps + 1
			}

			steps++
			lastPosition = image.Point{
				X: start.X,
				Y: start.Y,
			}
			start = possiblePositions[0]
			continue
		}

		// If multiple directions, walk them all
		maxDistance := 0
		var d int
		var ok bool
		for _, p := range possiblePositions {
			if p.X == destination.X && p.Y == destination.Y {
				return steps + 1
			}

			walkCacheKey := fmt.Sprintf("%d,%d,%v", p.X, p.Y, p.direction)
			d, ok = (*walkCache)[walkCacheKey]
			if !ok {
				d = walk(grid, p, pastPositions, walkCache, ignoreSlopes)
			}

			if d > maxDistance {
				maxDistance = d
			}

			(*walkCache)[walkCacheKey] = d
		}

		return steps + maxDistance
	}
}
