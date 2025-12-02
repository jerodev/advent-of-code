package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var ids []int
	var count, count2 int

	for _, line := range strings.Split(input, ",") {
		ids = util.StringToInts(line, "-")

		for ; ids[0] <= ids[1]; ids[0]++ {
			if isInvalid(ids[0]) {
				count += ids[0]
				count2 += ids[0]
			} else if isInvalidPart2(ids[0]) {
				fmt.Println(ids[0])
				count2 += ids[0]
			}
		}
	}

	fmt.Println(count)
	fmt.Println(count2)
}

// IsInvalid checks if a number is invalid
// A number is invalid if it exists of a specific sequence of numbers repeated twice
func isInvalid(v int) bool {
	l := int(math.Log10(float64(v))) + 1
	if l%2 == 1 {
		return false
	}

	half := l / 2
	for i := range half {
		if util.NthDigit(v, i+1) != util.NthDigit(v, half+i+1) {
			return false
		}
	}

	return true
}

// IsInvalidPart2 checks if a number is invalid
// A number is invalid if it exists of a specific sequence of numbers repeated any number of times
func isInvalidPart2(v int) bool {
	l := int(math.Log10(float64(v))) + 1
	if l == 1 {
		return false
	}

	// Special case where all digits are equal
	allEqual := true
	for i := range l {
		if i == 0 {
			continue
		}

		if util.NthDigit(v, i) != util.NthDigit(v, i+1) {
			allEqual = false
			break
		}
	}
	if allEqual {
		return true
	}

	seq := int(math.Floor(float64(l) / 2))
	var parts int
seqloop:
	for {
		seq -= 1
		if seq <= 0 {
			break
		}
		if l%seq > 0 {
			continue
		}
		parts = l/seq - 1

		for i := range seq {
			for p := range parts {
				if util.NthDigit(v, p*seq+i+1) != util.NthDigit(v, (p+1)*seq+i+1) {
					continue seqloop
				}
			}
		}

		return true
	}

	return false
}
