package main

import (
	"advent-of-code/util"
	"container/heap"

	"fmt"
)

type destination struct {
	Direction      byte
	DirectionCount int
	Distance       int
	Y              int
	X              int
}

type location struct {
	Distance  int
	Y         int
	X         int
	Direction byte
}

type locationHeap []destination

func (h locationHeap) Len() int { return len(h) }
func (h locationHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}
func (h locationHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *locationHeap) Push(x interface{}) {
	*h = append(*h, x.(destination))
}
func (h *locationHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]

	return x
}

func main() {
	grid := util.GridFromArgsInt()

	queue := &locationHeap{
		destination{
			Direction:      'R',
			DirectionCount: 1,
			Distance:       grid[0][1],
			Y:              0,
			X:              1,
		},
		destination{
			Direction:      'D',
			DirectionCount: 1,
			Distance:       grid[1][0],
			Y:              1,
			X:              0,
		},
	}
	heap.Init(queue)

	distances := map[string]location{}
	fmt.Println(dijkstra(grid, queue, &distances))

	// Print a grid showing the result path
	agrid := make([][]byte, len(grid))
	for i := range agrid {
		agrid[i] = make([]byte, len(grid[0]))
		for j := range grid[i] {
			agrid[i][j] = '.'
		}
	}

	y, x := len(grid)-1, len(grid[0])-1
	for {
		if y == 0 && x == 0 {
			break
		}

		position := fmt.Sprintf("%d,%d", y, x)
		location := distances[position]
		agrid[y][x] = location.Direction

		switch location.Direction {
		case '^':
			y++
		case 'v':
			y--
		case '<':
			x++
		case '>':
			x--
		}
	}

	util.PrintMatrix(agrid)
}

func dijkstra(grid [][]int, q *locationHeap, distances *map[string]location) int {
	for {
		d := heap.Pop(q).(destination)

		// No more than 3 consecutive moves in the same direction
		if d.DirectionCount > 3 {
			continue
		}

		// If the position is known, we already have the shortest position to it
		position := fmt.Sprintf("%d,%d", d.Y, d.X)
		if loc, ok := (*distances)[position]; ok && loc.Distance <= d.Distance {
			continue
		}

		// Register the distance
		arrow := byte('.')
		switch d.Direction {
		case 'U':
			arrow = '^'
		case 'D':
			arrow = 'v'
		case 'L':
			arrow = '<'
		case 'R':
			arrow = '>'
		}
		(*distances)[position] = location{
			Distance:  d.Distance,
			Y:         d.Y,
			X:         d.X,
			Direction: arrow,
		}

		// We are at the end position!
		if d.Y == len(grid)-1 && d.X == len(grid[0])-1 {
			return d.Distance
		}

		// Push new destinations
		// Going up
		if d.Y > 0 {
			dCount := 1
			if d.Direction == 'U' {
				dCount += d.DirectionCount
			}

			heap.Push(q, destination{
				Direction:      'U',
				DirectionCount: dCount,
				Distance:       d.Distance + grid[d.Y-1][d.X],
				Y:              d.Y - 1,
				X:              d.X,
			})
		}
		// Going down
		if d.Y+1 < len(grid) {
			dCount := 1
			if d.Direction == 'D' {
				dCount += d.DirectionCount
			}

			heap.Push(q, destination{
				Direction:      'D',
				DirectionCount: dCount,
				Distance:       d.Distance + grid[d.Y+1][d.X],
				Y:              d.Y + 1,
				X:              d.X,
			})
		}
		// Going left
		if d.X > 0 {
			dCount := 1
			if d.Direction == 'L' {
				dCount += d.DirectionCount
			}

			heap.Push(q, destination{
				Direction:      'L',
				DirectionCount: dCount,
				Distance:       d.Distance + grid[d.Y][d.X-1],
				Y:              d.Y,
				X:              d.X - 1,
			})
		}
		// Going right
		if d.X+1 < len(grid[0]) {
			dCount := 1
			if d.Direction == 'R' {
				dCount += d.DirectionCount
			}

			heap.Push(q, destination{
				Direction:      'R',
				DirectionCount: dCount,
				Distance:       d.Distance + grid[d.Y][d.X+1],
				Y:              d.Y,
				X:              d.X + 1,
			})
		}
	}
}
