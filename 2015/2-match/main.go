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
	inputs := strings.Split(input, "\n")

	var surface, ribbon int
	numbers := make([]int, 3)
	dims := make([]int, 3)
	for i := range inputs {
		numbers = util.StringToInts(inputs[i], "x")
		dims[0] = numbers[0] * numbers[1]
		dims[1] = numbers[1] * numbers[2]
		dims[2] = numbers[2] * numbers[0]

		slices.Sort(dims)
		slices.Sort(numbers)
		surface += 2*dims[0] + 2*dims[1] + 2*dims[2] + dims[0]
		ribbon += 2*numbers[0] + 2*numbers[1] + numbers[0]*numbers[1]*numbers[2]
	}

	fmt.Println(surface)
	fmt.Println(ribbon)
}
