package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	var rules [][2]int

	var line string
	var parts []string
	for scan.Scan() {
		line = scan.Text()
		if line == "" {
			break
		}

		parts = strings.SplitN(line, "|", 2)

		rules = append(rules, [2]int{
			util.MustAtoi(parts[0]),
			util.MustAtoi(parts[1]),
		})
	}

	var result, invalidResult int
	for scan.Scan() {
		row := scan.Text()
		ints := util.StringToInts(row, ",")

		if isValidRow(ints, rules) {
			mid := len(ints) / 2
			result += ints[mid]
		} else {
			ints = fixInvalidRow(ints, rules)

			mid := len(ints) / 2
			invalidResult += ints[mid]
		}
	}

	fmt.Println(result)
	fmt.Println(invalidResult)
}

func isValidRow(ints []int, rules [][2]int) bool {
	position := map[int]int{}
	for i := range ints {
		position[ints[i]] = i
	}

	var pos1, pos2 int
	var ok1, ok2 bool
	for i := range rules {
		pos1, ok1 = position[rules[i][0]]
		pos2, ok2 = position[rules[i][1]]

		if ok1 && ok2 && pos1 >= pos2 {
			return false
		}
	}

	return true
}

func fixInvalidRow(ints []int, rules [][2]int) []int {
outer:
	position := map[int]int{}
	for i := range ints {
		position[ints[i]] = i
	}

	var pos1, pos2 int
	var ok1, ok2 bool
	for i := range rules {
		pos1, ok1 = position[rules[i][0]]
		pos2, ok2 = position[rules[i][1]]

		if ok1 && ok2 && pos1 >= pos2 {
			ints[pos1], ints[pos2] = ints[pos2], ints[pos1]
			goto outer
		}
	}

	return ints
}
