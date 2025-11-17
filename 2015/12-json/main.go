package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	var sum int

	for i := 0; i < len(input); i++ {
		if input[i] == '-' || (input[i] >= '0' && input[i] <= '9') {
			sum += parseNumber(&i)
		}
	}

	fmt.Println(sum)
	fmt.Println(parseRed())
}

func parseRed() int {
	var sum, tmp int
	var red bool

	for i := 0; i < len(input); i++ {
		if input[i] == '-' || (input[i] >= '0' && input[i] <= '9') {
			sum += parseNumber(&i)
		}

		if input[i] == '{' {
			tmp, red = parseBlock(&i, false)
			if !red {
				sum += tmp
			}
		}
	}

	return sum
}

// parseBlock parses an entire json object {}
// Returns the sum of the numbers and whether the block contains "red"
func parseBlock(i *int, red bool) (int, bool) {
	var sum, tmp int
	var tmpRed bool
	*i++

	for ; *i < len(input); *i++ {
		if input[*i] == '}' {
			return sum, red
		}
		if input[*i] == '{' {
			tmp, tmpRed = parseBlock(i, red)
			if !tmpRed {
				sum += tmp
			}

			continue
		}

		if *i > 5 && string(input[*i-5:*i+1]) == `:"red"` {
			red = true
			continue
		}

		if !red && (input[*i] == '-' || (input[*i] >= '0' && input[*i] <= '9')) {
			sum += parseNumber(i)
			*i--
		}
	}

	// This should never be reached
	panic("Unclosed block!")
	return sum, red
}

// parseNumber parses the number starting at index
// The index stops at the next character after the number
func parseNumber(i *int) int {
	var negative bool
	if input[*i] == '-' {
		negative = true
		*i++

		if input[*i] < '0' || input[*i] > '9' {
			return 0
		}
	}

	var number int
	for {
		if input[*i] < '0' || input[*i] > '9' {
			if negative {
				number *= -1
			}
			return number
		}

		number = number*10 + int(input[*i]-'0')
		*i++
	}
}
