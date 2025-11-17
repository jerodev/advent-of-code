package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"strings"
)

const (
	on     = 1
	off    = 2
	toggle = 3
)

//go:embed input.txt
var input string

func main() {
	data := strings.Split(input, "\n")

	var grid [1000][1000]bool
	var brightness [1000][1000]int
	var a, trim, x1, y1, x2, y2, dx int
	var tmp, tmp2 []string
	for i := range data {
		switch data[i][:7] {
		case "toggle ":
			a = toggle
			trim = 7
		case "turn on":
			a = on
			trim = 8
		case "turn of":
			a = off
			trim = 9
		}

		data[i] = data[i][trim:]

		// Now parse positions
		tmp = strings.Split(data[i], " through ")
		tmp2 = strings.Split(tmp[0], ",")
		x1, y1 = util.MustAtoi(tmp2[0]), util.MustAtoi(tmp2[1])
		tmp2 = strings.Split(tmp[1], ",")
		x2, y2 = util.MustAtoi(tmp2[0]), util.MustAtoi(tmp2[1])

		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}

		for y1 <= y2 {
			dx = x1

			for dx <= x2 {
				switch a {
				case on:
					grid[y1][dx] = true
					brightness[y1][dx]++
				case off:
					grid[y1][dx] = false
					brightness[y1][dx] = max(0, brightness[y1][dx]-1)
				case toggle:
					grid[y1][dx] = !grid[y1][dx]
					brightness[y1][dx] += 2
				}

				dx++
			}

			y1++
		}
	}

	var count, b int
	for x1 = range grid {
		for y1 = range grid[x1] {
			if grid[x1][y1] {
				count++
			}

			b += brightness[x1][y1]
		}
	}

	fmt.Println(count)
	fmt.Println(b)
}
