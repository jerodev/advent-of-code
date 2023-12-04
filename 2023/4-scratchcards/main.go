package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
	"slices"
	"strings"
)

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	var points float64 = 0
	line := ""

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			line = strings.SplitN(line, ":", 2)[1]
			parts := strings.SplitN(line, "|", 2)

			winners := strings.Split(strings.TrimSpace(parts[0]), " ")
			numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

			score := len(intersectSlice(winners, numbers))
			if score > 0 {
				points += math.Pow(2, float64(score)-1)
			}

			if err == io.EOF {
				break
			}

			line = ""
		}

		line += string(b[0])
	}

	fmt.Println(points)
}

// intersectSlice returns a slice with elements that exist in both arrays
func intersectSlice(a, b []string) []string {
	result := []string{}

	for _, x := range a {
		x = strings.TrimSpace(x)
		if x == "" {
			continue
		}

		for _, y := range b {
			y = strings.TrimSpace(y)
			if x == y && !slices.Contains(result, x) {
				result = append(result, x)
				break
			}
		}
	}

	fmt.Println(a, b, result)

	return result
}
