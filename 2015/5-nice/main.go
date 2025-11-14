package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	data := strings.Split(input, "\n")

	var nice, nice2, vowels int
	var repeat, forbidden, repeatBetween, repeatDouble bool

	for i := range data {
		repeat, forbidden, repeatBetween, repeatDouble = false, false, false, false
		vowels = 0

		for j := range data[i] {
			if j > 0 {
				if data[i][j] == data[i][j-1] {
					repeat = true
				} else if (data[i][j-1] == 'a' && data[i][j] == 'b') || (data[i][j-1] == 'c' && data[i][j] == 'd') || (data[i][j-1] == 'p' && data[i][j] == 'q') || (data[i][j-1] == 'x' && data[i][j] == 'y') {
					forbidden = true
				}
			}
			if j > 1 {
				if data[i][j] == data[i][j-2] {
					repeatBetween = true
				}
			}
			if j > 2 {
				// Find out if the last two chars have appeared already
				for k := range j - 2 {
					if data[i][k] == data[i][j-1] && data[i][k+1] == data[i][j] {
						repeatDouble = true
					}
				}
			}

			if data[i][j] == 'a' || data[i][j] == 'e' || data[i][j] == 'i' || data[i][j] == 'o' || data[i][j] == 'u' {
				vowels++
			}
		}

		if repeat && !forbidden && vowels > 2 {
			nice++
		}
		if repeatBetween && repeatDouble {
			nice2++
		}
	}

	fmt.Println(nice)
	fmt.Println(nice2)
}
