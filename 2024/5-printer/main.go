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

	var result int
	var update map[int]int
	for scan.Scan() {
		valid := true

		parts = strings.Split(scan.Text(), ",")
		update = map[int]int{}
		for i := range parts {
			update[util.MustAtoi(parts[i])] = i
		}

		var pos1, pos2 int
		var ok1, ok2 bool
		for i := range rules {
			pos1, ok1 = update[rules[i][0]]
			pos2, ok2 = update[rules[i][1]]

			if ok1 && ok2 && pos1 >= pos2 {
				valid = false
				break
			}
		}

		if valid {
			mid := int(len(parts) / 2)
			result += util.MustAtoi(parts[mid])
		}
	}

	fmt.Println(result)
}
