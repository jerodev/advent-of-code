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
				m := findMirrors(rows, 0)
				if m > 0 {
					result += m * 100
				} else {
					result += findMirrors(columns, 0)
				}

				// Part 2
				m = findMirrors(rows, 1)
				if m > 0 {
					resultPart2 += m * 100
				} else {
					resultPart2 += findMirrors(columns, 1)
				}

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

func findMirrors(rows []string, smudgeCount int) int {
	lastIndex := len(rows) - 1

	for i := 1; i <= lastIndex; i++ {
		di := i
		smudges := smudgeCount
		isMirror := true
		for j := i - 1; j >= 0 && di <= lastIndex; j-- {
			if rows[j] != rows[di] {
				if smudges > 0 && countStringDiff(rows[j], rows[di]) == 1 {
					smudges--
				} else {
					isMirror = false
					break
				}
			}

			di++
		}

		if isMirror && smudges == 0 {
			return i
		}
	}

	return 0
}

func countStringDiff(a, b string) int {
	diff := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff
}
