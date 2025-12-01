package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var count, direction, zeroCount, loopZeroCount int
	counter := 50
	for _, l := range strings.Split(input, "\n") {
		count = util.MustAtoi(l[1:])
		if l[0] == 'L' {
			loopZeroCount += countZero(counter-count, counter-1)
			direction = -1
		} else {
			loopZeroCount += countZero(counter+1, counter+count)
			direction = 1
		}

		counter += direction * count
		if counter%100 == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
	fmt.Println(loopZeroCount)
}

func countZero(lo, hi int) int {
	if lo > hi {
		return 0
	}

	return util.FloorDiv(hi, 100) - util.FloorDiv(lo-1, 100)
}
