package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var facts = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	lines := strings.Split(input, "\n")

	var aunt, kv []string
	var exactAunt, rangedAunt string
	var tmp int
	var exact, rangedExact bool

	for _, l := range lines {
		exact = true
		rangedExact = true
		aunt = strings.Split(l, " ")

		for i := range aunt[1:] {
			kv = strings.SplitN(aunt[i+1], ":", 2)
			tmp = util.MustAtoi(kv[1])

			exact = exact && facts[kv[0]] == tmp
			switch kv[0] {
			case "cats", "trees":
				rangedExact = rangedExact && facts[kv[0]] < tmp
			case "pomeranians", "goldfish":
				rangedExact = rangedExact && facts[kv[0]] > tmp
			default:
				rangedExact = rangedExact && facts[kv[0]] == tmp
			}
		}

		// All matches, Aunt Sue found!
		if exact {
			exactAunt = aunt[0]
		}
		if rangedExact {
			rangedAunt = aunt[0]
		}
		if exactAunt != "" && rangedAunt != "" {
			break
		}
	}

	fmt.Println(exactAunt)
	fmt.Println(rangedAunt)
}
