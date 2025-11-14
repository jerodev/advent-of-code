package main

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input.txt
var input []byte

func main() {
	var x, y, sx, sy, rx, ry int
	coords := [][2]int{{0, 0}}
	coords2 := [][2]int{{0, 0}}
	for i := range input {
		switch input[i] {
		case '>':
			x++
			if i%2 == 0 {
				sx++
			} else {
				rx++
			}
		case '<':
			x--
			if i%2 == 0 {
				sx--
			} else {
				rx--
			}
		case '^':
			y--
			if i%2 == 0 {
				sy--
			} else {
				ry--
			}
		case 'v':
			y++
			if i%2 == 0 {
				sy++
			} else {
				ry++
			}
		}

		if !slices.Contains(coords, [2]int{x, y}) {
			coords = append(coords, [2]int{x, y})
		}
		if i%2 == 0 {
			if !slices.Contains(coords2, [2]int{sx, sy}) {
				coords2 = append(coords2, [2]int{sx, sy})
			}
		} else {
			if !slices.Contains(coords2, [2]int{rx, ry}) {
				coords2 = append(coords2, [2]int{rx, ry})
			}
		}
	}

	fmt.Println(len(coords))
	fmt.Println(len(coords2))
}
