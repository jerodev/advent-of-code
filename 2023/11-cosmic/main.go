package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
)

const (
	SPACE_EMPTY  = byte('.')
	SPACE_GALAXY = byte('#')

	// Part 1
	// SPACE_EXPANSION = 2

	// Part 2
	SPACE_EXPANSION = 1e6
)

func main() {
	file := util.FileFromArgs()

	universe := [][]byte{
		{},
	}
	rowNumber := 0
	b := make([]byte, 1)

	emptyRows := []int{}
	emptyColumns := []int{}

	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}

		if b[0] == '\n' {
			empty := true
			for _, b := range universe[rowNumber] {
				if b != SPACE_EMPTY {
					empty = false
					break
				}
			}
			if empty {
				emptyRows = append(emptyRows, rowNumber)
			}

			rowNumber++
			universe = append(universe, []byte{})
			continue
		}

		universe[rowNumber] = append(universe[rowNumber], b[0])
	}

	// Find empty columns
	for i := 0; i < len(universe[0]); i++ {
		empty := true
		for _, row := range universe {
			if row[i] != SPACE_EMPTY {
				empty = false
				break
			}
		}

		if empty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	// Find galaxies in the universe
	var galaxies [][2]int
	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] == SPACE_GALAXY {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	fmt.Println(galaxies)

	totalDistance := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += distanceBetween(galaxies[i], galaxies[j], emptyRows, emptyColumns)
		}
	}

	fmt.Println(totalDistance)
}

func distanceBetween(a, b [2]int, emptyRows, emptyColumns []int) int {
	if a[0] > b[0] {
		a[0], b[0] = b[0], a[0]
	}
	if a[1] > b[1] {
		a[1], b[1] = b[1], a[1]
	}

	dx := b[1] - a[1]
	for _, i := range emptyColumns {
		if i > a[1] && i < b[1] {
			dx += SPACE_EXPANSION - 1
		}
	}

	dy := b[0] - a[0]
	for _, i := range emptyRows {
		if i > a[0] && i < b[0] {
			dy += SPACE_EXPANSION - 1
		}
	}

	return dx + dy
}
