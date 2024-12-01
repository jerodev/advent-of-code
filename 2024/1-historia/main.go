package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func main() {
	file := util.FileFromArgs()
	defer file.Close()

	var lists [2][]int
	var numbers []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		numbers = strings.SplitN(scan.Text(), "   ", 2)
		lists[0] = append(lists[0], util.MustAtoi(numbers[0]))
		lists[1] = append(lists[1], util.MustAtoi(numbers[1]))
	}

	slices.Sort(lists[0])
	slices.Sort(lists[1])

	sum := 0
	simScore := 0
	for i := range lists[0] {
		sum += util.Abs(lists[0][i] - lists[1][i])
		simScore += lists[0][i] * util.CountOccurrences(lists[1], lists[0][i])
	}

	fmt.Println(sum)
	fmt.Println(simScore)
}
