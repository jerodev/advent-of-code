package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"math/rand"
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
			mid := int(len(ints) / 2)
			result += ints[mid]
		} else {
			for {
				rand.Shuffle(len(ints), func(i, j int) {
					ints[i], ints[j] = ints[j], ints[i]
				})

				if isValidRow(ints, rules) {
					mid := int(len(ints) / 2)
					invalidResult += ints[mid]
					break
				}
			}
		}
	}

	fmt.Println(result)
	fmt.Println(invalidResult)
}

func isValidRow(ints []int, rules [][2]int) bool {
	update := map[int]int{}
	for i := range ints {
		update[ints[i]] = i
	}

	var pos1, pos2 int
	var ok1, ok2 bool
	for i := range rules {
		pos1, ok1 = update[rules[i][0]]
		pos2, ok2 = update[rules[i][1]]

		if ok1 && ok2 && pos1 >= pos2 {
			return false
		}
	}

	return true
}
