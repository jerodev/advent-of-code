package main

import (
	"advent-of-code/util"
	"fmt"
	"regexp"
	"strings"
)

var mulRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

const (
	doFunc   = "do()"
	dontFunc = "don't()"
)

func main() {
	input := util.ReadFileFromArgs()

	// Part 1
	fmt.Println(execMuls(input))

	// Part 2
	var startIndex, stopIndex, sum int
	for {
		stopIndex = startIndex + strings.Index(input[startIndex:], dontFunc)
		if stopIndex-startIndex == -1 {
			sum += execMuls(input[startIndex:])
			break
		}

		sum += execMuls(input[startIndex:stopIndex])

		startIndex = stopIndex + strings.Index(input[stopIndex:], doFunc)
		if startIndex-stopIndex == -1 {
			break
		}
	}

	fmt.Println(sum)
}

func execMuls(input string) int {
	matches := mulRegex.FindAllStringSubmatch(input, -1)

	var sum int
	for i := range matches {
		sum += util.MustAtoi(matches[i][1]) * util.MustAtoi(matches[i][2])
	}

	return sum
}
