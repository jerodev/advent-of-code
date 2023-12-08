package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
)

type location struct {
	Name  string
	Left  *location
	Right *location
}

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	line := ""
	directions := ""
	locationLines := []string{}
	locations := map[string]*location{}
	startLocations := []*location{}

	for {
		_, err := file.Read(b)
		if err == io.EOF || b[0] == '\n' {
			if line == "" {
				continue
			}
			if directions == "" {
				directions = line
				line = ""
				continue
			}

			// Remember raw location strings
			locationLines = append(locationLines, line)

			// Add location to the map
			locations[line[:3]] = &location{
				Name: line[:3],
			}

			if err == io.EOF {
				break
			}

			line = ""
			continue
		}

		line += string(b[0])
	}

	// Now loop once more to add the left and right locations
	for _, line := range locationLines {
		location := locations[line[:3]]
		location.Left = locations[line[7:10]]
		location.Right = locations[line[12:15]]

		if line[2] == 'A' {
			startLocations = append(startLocations, location)
		}
	}

	// Loop and loop and loop untill we find ZZZ
	position := locations["AAA"]
	steps := 0
outer:
	for {
		for _, LR := range directions {
			if LR == 'L' {
				position = position.Left
			} else {
				position = position.Right
			}
			steps++

			if position.Name == "ZZZ" {
				break outer
			}
		}
	}

	fmt.Println(steps)

	// Part two, loop for all start locations
	steps = 0
outer2:
	for {
		for _, LR := range directions {
			allZ := true
			for i, loc := range startLocations {
				if LR == 'L' {
					startLocations[i] = loc.Left
				} else {
					startLocations[i] = loc.Right
				}

				if startLocations[i].Name[2] != 'Z' {
					allZ = false

				}
			}

			steps++

			if allZ {
				break outer2
			}
		}
	}

	fmt.Println(steps)
}
