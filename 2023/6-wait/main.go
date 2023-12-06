package main

import (
	"advent-of-code/util"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var numbersRegex = regexp.MustCompile(`\D`)
var spacesRegex = regexp.MustCompile(`\s+`)

type contest struct {
	Time     int
	Distance int
}

func main() {
	lines := strings.SplitN(util.ReadFileFromArgs(), "\n", 2)

	// Parse the lines
	times := spacesRegex.Split(strings.TrimSpace(lines[0][11:]), -1)
	distances := spacesRegex.Split(strings.TrimSpace(lines[1][11:]), -1)

	contests := []contest{}
	for i := range times {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		if time == 0 {
			continue
		}

		contests = append(contests, contest{
			Time:     time,
			Distance: distance,
		})
	}

	fmt.Println(contests)

	// Solve the problem
	posibilities := 1
	for _, c := range contests {
		numbers := make([]int, c.Time)
		for i := 0; i < c.Time; i++ {
			numbers[i] = i + 1
		}

		minIndex := slices.IndexFunc(numbers, func(hold int) bool {
			return isWinning(c, hold)
		})

		slices.Reverse(numbers)
		maxIndex := slices.IndexFunc(numbers, func(hold int) bool {
			return isWinning(c, hold)
		})
		maxIndex = len(numbers) - maxIndex

		posibilities *= maxIndex - minIndex
	}

	fmt.Println(posibilities)

	// Problem part 2
	time, _ := strconv.Atoi(numbersRegex.ReplaceAllString(lines[0], ""))
	destination, _ := strconv.Atoi(numbersRegex.ReplaceAllString(lines[1], ""))
	c := contest{
		Time:     time,
		Distance: destination,
	}

	fmt.Println(c)

	numbers := make([]int, c.Time)
	for i := 0; i < c.Time; i++ {
		numbers[i] = i + 1
	}

	minIndex := slices.IndexFunc(numbers, func(hold int) bool {
		return isWinning(c, hold)
	})

	slices.Reverse(numbers)
	maxIndex := slices.IndexFunc(numbers, func(hold int) bool {
		return isWinning(c, hold)
	})
	maxIndex = len(numbers) - maxIndex

	fmt.Println(maxIndex - minIndex)
}

func isWinning(contest contest, time int) bool {
	return (contest.Time-time)*time > contest.Distance
}
