package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
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

		if buttonA.X+buttonA.Y > buttonB.X+buttonB.Y {
			da, db = simulate(buttonA, buttonB, prize)
			dTokens = da*3 + db
		} else {
			da, db = simulate(buttonB, buttonA, prize)
			dTokens = db*3 + da
		}

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
	maxA := min(100, prize.X/buttonA.X, prize.Y/buttonA.Y)
	maxB := min(100, prize.X/buttonB.X, prize.Y/buttonB.Y)
	var da, db, dPresses int
	var p position
	var found bool

outer:
	for da = maxA; da >= 0; da-- {
		dPresses = da
		p.X = buttonA.X * dPresses
		p.Y = buttonA.Y * dPresses

		for db = 0; db < maxB; db++ {
			if p.X+db*buttonB.X == prize.X && p.Y+db*buttonB.Y == prize.Y {
				dPresses += db
				found = true
				break outer
			}
		}

	}

	if !found {
		return 0, 0
	}

	return da, db
}
