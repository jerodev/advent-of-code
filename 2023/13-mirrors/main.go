package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
)

const (
	ASH  = '.'
	ROCK = '#'
)

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)

	rows := []string{}
	columns := []string{}
	line := ""
	columnNumber := 0

	result := 0
	resultPart2 := 0

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			if len(line) > 0 {
				rows = append(rows, line)
			}

			if (line == "" || err == io.EOF) && len(rows) > 0 {
				// Part 1
				result += findMirrors(rows) * 100
				result += findMirrors(columns)

				// Part 2
				m := findMirrorsWithSmudges(rows)
				if m > 0 {
					resultPart2 += m * 100
				} else {
					resultPart2 += findMirrorsWithSmudges(columns)
				}
				fmt.Println(resultPart2)

				rows = []string{}
				columns = []string{}
			}

			if err == io.EOF {
				break
			}

			line = ""
			columnNumber = 0
			continue
		}

		if len(columns) < columnNumber+1 {
			columns = append(columns, string(b))
		} else {
			columns[columnNumber] += string(b)
		}
		columnNumber++

		line += string(b)
	}

	fmt.Println(result)
	fmt.Println(resultPart2)
}

func findMirrors(rows []string) int {
	lastIndex := len(rows) - 1

	for i := 1; i <= lastIndex; i++ {
		di := i
		isMirror := true
		for j := i - 1; j >= 0 && di <= lastIndex; j-- {
			if rows[j] != rows[di] {
				isMirror = false
				break
			}

			di++
		}

		if isMirror {
			fmt.Println(rows, i)
			return i
		}
	}

	return 0
}

func findMirrorsWithSmudges(rows []string) int {
	lastIndex := len(rows) - 1

	smallest := 0

	for x := range rows {
		for y := range rows[x] {
			if rows[x][y] == ASH {
				rows[x] = rows[x][:y] + string(ROCK) + rows[x][y+1:]
			} else {
				rows[x] = rows[x][:y] + string(ASH) + rows[x][y+1:]
			}

			for i := 1; i <= lastIndex; i++ {
				di := i
				isMirror := true
				for j := i - 1; j >= 0 && di <= lastIndex; j-- {
					if rows[j] != rows[di] {
						isMirror = false
						break
					}

					di++
				}

				if isMirror {
					if smallest == 0 || i < smallest {
						smallest = i
					}

					break
				}
			}

			// Clean back up
			if rows[x][y] == ASH {
				rows[x] = rows[x][:y] + string(ROCK) + rows[x][y+1:]
			} else {
				rows[x] = rows[x][:y] + string(ASH) + rows[x][y+1:]
			}
		}
	}

	return smallest
}
