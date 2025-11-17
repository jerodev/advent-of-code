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

var locations []string
var flights = map[string][]flight{}

func main() {
	lines := strings.Split(input, "\n")
	var parts []string
	var ok bool
	for _, l := range lines {
		parts = strings.Split(l, " ")
		if !slices.Contains(locations, parts[0]) {
			locations = append(locations, parts[0])
		}
		if !slices.Contains(locations, parts[2]) {
			locations = append(locations, parts[2])
		}

		_, ok = flights[parts[0]]
		if !ok {
			flights[parts[0]] = []flight{
				{parts[0], parts[2], util.MustAtoi(parts[4])},
			}
		} else {
			flights[parts[0]] = append(flights[parts[0]], flight{parts[0], parts[2], util.MustAtoi(parts[4])})
		}

		_, ok = flights[parts[2]]
		if !ok {
			flights[parts[2]] = []flight{
				{parts[2], parts[0], util.MustAtoi(parts[4])},
			}
		} else {
			flights[parts[2]] = append(flights[parts[2]], flight{parts[2], parts[0], util.MustAtoi(parts[4])})
		}
	}

	// Run the tree!
	// Try starting from each of the locations
	var shortestDistance, longestDistance int
	for i := range locations {
		d := flyShortest([]string{locations[i]}, 0, shortestDistance)
		if d > 0 && (shortestDistance == 0 || d < shortestDistance) {
			shortestDistance = d
		}

		d = flyLongest([]string{locations[i]}, 0, longestDistance)
		if d > longestDistance {
			longestDistance = d
		}
	}

	fmt.Println(shortestDistance)
	fmt.Println(longestDistance)
}

type flight struct {
	from, to string
	duration int
}

func flyShortest(traject []string, distance int, shortestDistance int) int {
	if len(traject) == len(locations) {
		// All locations visited!
		return distance
	}

	// Attempt all flights starting from this location
	for _, f := range flights[traject[len(traject)-1]] {
		if slices.Contains(traject, f.to) {
			// Been there, done that
			continue
		}

		// If our travel route is still shorter than the shortest route, keep going!
		if shortestDistance == 0 || shortestDistance > distance+f.duration {
			d := flyShortest(append(traject, f.to), distance+f.duration, shortestDistance)
			if shortestDistance == 0 || d < shortestDistance {
				shortestDistance = d
			}
		}
	}

	// No possible route found, we are stuck...
	return shortestDistance
}

func flyLongest(traject []string, distance int, longestDistance int) int {
	if len(traject) == len(locations) {
		// All locations visited!
		return distance
	}

	// Attempt all flights starting from this location
	for _, f := range flights[traject[len(traject)-1]] {
		if slices.Contains(traject, f.to) {
			// Been there, done that
			continue
		}

		d := flyLongest(append(traject, f.to), distance+f.duration, longestDistance)
		if d > longestDistance {
			longestDistance = d
		}
	}

	// No possible route found, we are stuck...
	return longestDistance
}
