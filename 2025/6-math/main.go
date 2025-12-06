package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var modifiers []byte
	for i := strings.LastIndex(input, "\n") + 1; i < len(input); i++ {
		if input[i] != ' ' {
			modifiers = append(modifiers, input[i])
		}
	}

	parsePartOne(modifiers)
	parsePartTwo(modifiers)
}

func parsePartOne(modifiers []byte) {
	var data [][]int

	for _, l := range strings.Split(input, "\n") {
		if strings.HasPrefix(l, "+") || strings.HasPrefix(l, "*") {
			break
		}

		data = append(data, util.StringToInts(l, " "))
	}

	var count, i int

	counts := data[0]
	for i = range data {
		if i == 0 {
			continue
		}

		for x := range data[0] {
			if modifiers[x] == '+' {
				counts[x] += data[i][x]
			} else if data[i][x] > 0 {
				counts[x] *= data[i][x]
			}
		}
	}

	for i = range counts {
		count += counts[i]
	}

	fmt.Println(counts)
	fmt.Println(count)
}

func parsePartTwo(modifiers []byte) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	maxLength := slices.MaxFunc(lines, func(a, b string) int {
		return len(a) - len(b)
	})

	data := [][]int{{}}
	for x := range maxLength {
		var nr int
		for y := range lines {
			if x >= len(lines[y]) {
				continue
			}

			if lines[y][x] != ' ' {
				nr = nr*10 + int(lines[y][x]-'0')
			}
		}

		if nr > 0 {
			data[len(data)-1] = append(data[len(data)-1], nr)
			continue
		}

		// No number found, start a new set
		data = append(data, []int{})
	}

	fmt.Println(data)

	var count, i, j int
	counts := make([]int, len(data))
	for i = range data {
		for j = range data[i] {
			if modifiers[i] == '+' {
				counts[i] += data[i][j]
			} else if data[i][j] > 0 {
				if j == 0 {
					counts[i] = data[i][j]
				} else {
					counts[i] *= data[i][j]
				}
			}
		}
	}

	for i = range counts {
		count += counts[i]
	}

	fmt.Println(counts)
	fmt.Println(count)
}
