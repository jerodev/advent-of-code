package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1113222113"
	iterations := 50

	var b strings.Builder

	var c byte
	var cc int
	for range iterations {
		b.Reset()
		c = 0
		cc = 0

		for i := range input {
			cc++

			if input[i] != c {
				if c > 0 {
					b.WriteString(strconv.Itoa(cc))
					b.WriteByte(c)
				}

				c = input[i]
				cc = 0
			}
		}

		b.WriteString(strconv.Itoa(cc + 1))
		b.WriteByte(c)

		input = b.String()
	}

	fmt.Println(len(input))
}
