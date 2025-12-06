package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var count, count2 int

	var bank []int
	for _, l := range strings.Split(input, "\n") {
		bank = util.StringToInts(l, "")
		count += part1(bank)
		count2 += part2(bank)
	}

	fmt.Println(count)
	fmt.Println(count2)
}

func part1(bank []int) int {
	hi := slices.Max(bank[:len(bank)-1])
	lo := slices.Max(bank[slices.Index(bank, hi)+1:])

	return hi*10 + lo
}

func part2(bank []int) int {
	batteries := make([]int, 0, 12)

	var hi int
	for len(batteries) < 12 {
		hi = slices.Max(bank[:len(bank)-11+len(batteries)])

		batteries = append(batteries, hi)
		bank = bank[slices.Index(bank, hi)+1:]
	}

	// Make the byte array a number
	var sum int
	for i := range batteries {
		sum += batteries[i] * int(math.Pow(10, 11-float64(i)))
	}

	return sum
}
