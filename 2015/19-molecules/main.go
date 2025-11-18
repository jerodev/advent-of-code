package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

var transformations = map[string][]string{}
var transfs [][2]string
var electrons []string

func main() {
	parse()
	part1()
	part2()
}

func part1() {
	molecules := map[string]bool{}

	// Loop over all characters of the string to try and replace atoms
	for i := 0; i < len(input); i++ {
		for p, t := range transformations {
			if strings.HasPrefix(input[i:], p) {
				for _, tt := range t {
					molecules[input[:i]+tt+input[i+len(p):]] = true
				}
			}
		}
	}

	fmt.Println(len(molecules))
}

func part2() {
	// Sort the transformations so the biggest changes are first
	slices.SortFunc(transfs, func(a, b [2]string) int {
		return len(b[0]) - len(a[0])
	})

	var count int
	var changed bool
	for {
		for i := range transfs {
			if strings.Contains(input, transfs[i][0]) {
				input = strings.Replace(input, transfs[i][0], transfs[i][1], 1)
				changed = true
				count++
				break
			}
		}

		if !changed {
			fmt.Println("INF")
			return
		}

		fmt.Println(input)

		if slices.Contains(electrons, input) {
			break
		}
	}

	fmt.Println(electrons)
	fmt.Println(count)
}

func parse() {
	var i int
	var line string
	var parts []string
	var ok bool

	lines := strings.Split(input, "\n")
	for i, line = range lines {
		if line == "" {
			i++
			break
		}

		parts = strings.Split(line, " => ")
		if parts[0] == "e" {
			electrons = append(electrons, parts[1])
		} else {
			transfs = append(transfs, [2]string{parts[1], parts[0]})

			_, ok = transformations[parts[0]]
			if ok {
				transformations[parts[0]] = append(transformations[parts[0]], parts[1])
			} else {
				transformations[parts[0]] = []string{parts[1]}
			}
		}
	}

	input = lines[i]
}
