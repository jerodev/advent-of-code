package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	var counter, basementPos int
	for i := range input {
		if input[i] == '(' {
			counter++
		} else {
			counter--
		}

		if counter < 0 && basementPos == 0 {
			basementPos = i + 1
		}
	}

	fmt.Println(counter)
	fmt.Println(basementPos)
}
