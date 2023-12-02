package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	sum := 0
	powerSum := 0
	gameIndex := 1
	line := ""
	inPrefix := true

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			if lineIsPossible(line) {
				sum += gameIndex
			}
			powerSum += linePower(line)

			if err == io.EOF {
				break
			}

			gameIndex++
			line = ""
			inPrefix = true
		}

		if inPrefix {
			if b[0] == ':' {
				_, _ = file.Read(b)
				inPrefix = false
			}

			continue
		}

		line += string(b[0])
	}

	fmt.Println(sum)
	fmt.Println(powerSum)
}

// lineIsPossible checks if the given amounts of cubes in the line are possible
func lineIsPossible(line string) bool {
	groups := strings.Split(line, ";")
	for _, group := range groups {
		cubes := strings.Split(group, ",")

		for _, cube := range cubes {
			split := strings.SplitN(strings.Trim(cube, " "), " ", 2)
			amount, _ := strconv.Atoi(split[0])

			if split[1] == "blue" && amount > blue {
				return false
			}
			if split[1] == "green" && amount > green {
				return false
			}
			if split[1] == "red" && amount > red {
				return false
			}
		}
	}

	return true
}

// linePower calculates the power for a given line
func linePower(line string) int {
	minBlue, minGreen, minRed := 1, 1, 1

	groups := strings.Split(line, ";")
	for _, group := range groups {
		cubes := strings.Split(group, ",")

		for _, cube := range cubes {
			split := strings.SplitN(strings.Trim(cube, " "), " ", 2)
			amount, _ := strconv.Atoi(split[0])

			if split[1] == "blue" && amount > minBlue {
				minBlue = amount
			} else if split[1] == "green" && amount > minGreen {
				minGreen = amount
			} else if split[1] == "red" && amount > minRed {
				minRed = amount
			}
		}
	}

	return minBlue * minGreen * minRed
}
