package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var timer int

func main() {
	var deers []deer
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l := strings.Split(line, " ")
		deers = append(deers, deer{
			l[0],
			util.MustAtoi(l[3]),
			util.MustAtoi(l[6]),
			util.MustAtoi(l[13]),
			0,
			0,
		})
	}

	for range 2503 {
		deers = advance(deers)
	}

	for _, d := range deers {
		fmt.Printf("%s:\t%v\t%v\n", d.name, d.position, d.points)
	}
}

type deer struct {
	name     string
	speed    int
	time     int
	rest     int
	position int
	points   int
}

func advance(deers []deer) []deer {
	timer++

	var tmp, winner int
	for i := range deers {
		// Determine resting state
		tmp = (timer - 1) % (deers[i].time + deers[i].rest)
		if tmp < deers[i].time {
			deers[i].position += deers[i].speed
		}

		if deers[i].position > winner {
			winner = deers[i].position
		}
	}

	for i := range deers {
		if deers[i].position == winner {
			deers[i].points++
		}
	}

	return deers
}
