package main

import (
	"advent-of-code/util"
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
)

const (
	bounds  = '#'
	empty   = '.'
	visited = '@'
	finish  = 'E'
)

var directions = map[byte]position{
	'N': {0, -1},
	'E': {1, 0},
	'S': {0, 1},
	'W': {-1, 0},
}

var grid [][]byte
var player path

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		x := bytes.IndexRune(scan.Bytes(), 'S')
		if x > -1 {
			player.X = x
			player.Y = len(grid)
			player.Direction = 'E'
		}

		gridRow := make([]byte, len(scan.Bytes()))
		copy(gridRow, scan.Bytes())
		grid = append(grid, gridRow)
	}

	paths := &pathHeap{
		player,
	}
	heap.Init(paths)

	var p path
	for {
		if paths.Len() == 0 {
			util.PrintMatrix(grid)
			fmt.Println(finish)
			panic("No path found...")
		}

		p = heap.Pop(paths).(path)
		if grid[p.Y][p.X] == finish {
			fmt.Println(p.Score)
			break
		}

		var dx, dy int
		for d, dd := range directions {
			dx, dy = p.X+dd.X, p.Y+dd.Y
			if grid[dy][dx] == bounds || grid[dy][dx] == visited {
				continue
			}
			if grid[dy][dx] == empty {
				grid[dy][dx] = visited
			}

			// Going straight on is free
			if p.Direction == d {
				heap.Push(paths, path{
					X:         dx,
					Y:         dy,
					Score:     p.Score + 1,
					Direction: d,
				})

				continue
			}

			// No going backsies
			if directions[p.Direction].X+dd.X == 0 && directions[p.Direction].Y+dd.Y == 0 {
				continue
			}

			// Turn here
			heap.Push(paths, path{
				X:         dx,
				Y:         dy,
				Score:     p.Score + 1001,
				Direction: d,
			})
		}
	}
}

type position struct {
	X, Y int
}

type path struct {
	X, Y      int
	Score     int
	Direction byte
}

type pathHeap []path

func (h pathHeap) Len() int { return len(h) }
func (h pathHeap) Less(i, j int) bool {
	return h[i].Score < h[j].Score
}
func (h pathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *pathHeap) Push(x interface{}) {
	*h = append(*h, x.(path))
}
func (h *pathHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]

	return x
}
