package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type elf struct {
	calories []int
}

func (e elf) totalCalories() int {
	total := 0
	for _, c := range e.calories {
		total += c
	}

	return total
}

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	line := ""

	var elves []elf
	var e elf

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			if line != "" {
				calories, _ := strconv.Atoi(line)
				e.calories = append(e.calories, calories)
				line = ""
			} else {
				elves = append(elves, e)
				e = elf{}
			}

			if err == io.EOF {
				elves = append(elves, e)
				break
			}

			continue
		}

		line += string(b[0])
	}

	// Part 1: Most calories
	maxCalories := 0
	for _, e := range elves {
		calories := e.totalCalories()
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Println(maxCalories)

	// Part 2: Top three calories
	slices.SortFunc(elves, func(i, j elf) int {
		return j.totalCalories() - i.totalCalories()
	})
	fmt.Println(elves[0].totalCalories() + elves[1].totalCalories() + elves[2].totalCalories())
}
