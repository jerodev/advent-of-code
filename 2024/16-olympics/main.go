package main

import (
	"advent-of-code/util"
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
)

const (
	bounds = '#'
	empty  = '.'
	finish = 'E'
)

var directions = map[byte]position{
	'N': {0, -1},
	'E': {1, 0},
	'S': {0, 1},
	'W': {-1, 0},
}

var grid [][]byte
var player path
var visited [][]int

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		visited = append(visited, make([]int, len(scan.Bytes())))
		x := bytes.IndexRune(scan.Bytes(), 'S')
		if x > -1 {
			player.X = x
			player.Y = len(grid)
			player.Direction = 'E'
			player.History = []position{
				{player.X, player.Y},
			}
		}

		gridRow := make([]byte, len(scan.Bytes()))
		copy(gridRow, scan.Bytes())
		grid = append(grid, gridRow)
	}

	paths := &pathHeap{
		player,
	}
	heap.Init(paths)

	var finishes []path
	var p path
	for {
		// All shortest paths are found, stop searching!
		if paths.Len() == 0 {
			break
		}

		p = heap.Pop(paths).(path)
		if grid[p.Y][p.X] == finish {
			if len(finishes) == 0 || p.Score == finishes[0].Score {
				fmt.Println(p.Score)           // Part 1
				finishes = append(finishes, p) // Part 2: Find all shortest routes
			}
			continue
		}

		// Don't bother paths that are higher than the finish
		if p.Score > 109516 { // 109516 is the input answer for part 1 ;)
			continue
		}

		var dx, dy int
	outer:
		for d, dd := range directions {
			dx, dy = p.X+dd.X, p.Y+dd.Y
			if grid[dy][dx] == bounds {
				continue
			}

			// Have we been here before?
			for _, h := range p.History {
				if h.X == dx && h.Y == dy {
					continue outer
				}
			}

			// Have we got enough points to be here?
			if visited[dy][dx] > 0 && p.Score > visited[dy][dx] {
				continue
			}
			visited[dy][dx] = p.Score

			if p.Direction == d { // Going straight on is free
				heap.Push(paths, path{
					X:         dx,
					Y:         dy,
					Score:     p.Score + 1,
					Direction: d,
					History:   makeHistory(p.History, position{dx, dy}),
				})
			} else { // Turning costs 1000
				heap.Push(paths, path{
					X:         dx,
					Y:         dy,
					Score:     p.Score + 1001,
					Direction: d,
					History:   makeHistory(p.History, position{dx, dy}),
				})
			}
		}
	}

	fmt.Println(finishes)

	// Part 2: now find unique points
	var count int

	for _, f := range finishes {
		for _, fh := range f.History {
			if grid[fh.Y][fh.X] != 'O' {
				count++
				grid[fh.Y][fh.X] = 'O'
			}
		}
	}

	util.PrintMatrix(grid)
	fmt.Println(count)
}

// makeHistory creates a brand new slice containing elements from h and p
func makeHistory(h []position, p position) []position {
	newSlice := make([]position, len(h)+1)

	var i int
	for i = range h {
		newSlice[i] = position{h[i].X, h[i].Y}
	}
	newSlice[i+1] = p

	return newSlice
}

type position struct {
	X, Y int
}

type path struct {
	X, Y      int
	Score     int
	Direction byte
	History   []position
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
