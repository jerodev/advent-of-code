package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strings"
)

var numberStrings = [...]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	file := util.FileFromArgs()

	left, right, sum := -1, -1, 0

	b := make([]byte, 1)
	buff := ""
	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}

		if b[0] >= '0' && b[0] <= '9' {
			// 48 is the byte value for 0
			if left == -1 {
				left = int(b[0]) - 48
			} else {
				right = int(b[0]) - 48
			}

			buff = ""

			continue
		}

		if b[0] == '\n' {
			if right == -1 {
				sum += left * 11
			} else {
				sum += left*10 + right
			}

			left, right = -1, -1
			buff = ""

			continue
		}

		buff += string(b[0])
		for i, numberString := range numberStrings {
			if strings.HasSuffix(buff, numberString) {
				if left == -1 {
					left = i + 1
				} else {
					right = i + 1
				}

				break
			}
		}
	}

	// Sum numbers once more
	if right == -1 {
		sum += left * 11
	} else {
		sum += left*10 + right
	}

	fmt.Println(sum)
}
