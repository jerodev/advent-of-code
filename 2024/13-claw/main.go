package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"math"
	"regexp"
)

var positionRegex *regexp.Regexp

func main() {
	positionRegex = regexp.MustCompile(`X(\+|=)(\d+), Y(\+|=)(\d+)`)

	var tokens int

	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	var buttonA, buttonB, prize position
	var da, db, dTokens int
	for {
		if !scan.Scan() {
			break
		}
		buttonA = parsePosition(scan.Bytes())
		scan.Scan()
		buttonB = parsePosition(scan.Bytes())
		scan.Scan()
		prize = parsePosition(scan.Bytes())
		scan.Scan()

		// Part 2
		prize.X += 10000000000000
		prize.Y += 10000000000000

		da, db = simulate(buttonA, buttonB, prize)
		dTokens = da*3 + db

		fmt.Println(buttonA, buttonB, prize, dTokens)

		tokens += dTokens
	}

	fmt.Println(tokens)
}

type position struct {
	X, Y int
}

func parsePosition(line []byte) position {
	matches := positionRegex.FindSubmatch(line)

	return position{
		util.MustAtoi(string(matches[2])),
		util.MustAtoi(string(matches[4])),
	}
}

func simulate(buttonA, buttonB, prize position) (int, int) {
	var da, db float64

	db = float64(buttonA.Y*prize.X-buttonA.X*prize.Y) / float64(buttonB.X*buttonA.Y-buttonB.Y*buttonA.X)
	if db != math.Floor(db) {
		return 0, 0
	}

	da = float64(prize.Y-int(db)*buttonB.Y) / float64(buttonA.Y)

	return int(da), int(db)
}
