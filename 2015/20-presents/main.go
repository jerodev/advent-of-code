package main

import "fmt"

const input = 29_000_000
const limit = input / 10

func main() {
	houses := make([]int, limit)

	// Part 1
	for e := 1; e < limit; e++ {
		for h := e; h < limit; h += e {
			houses[h] += e * 10
		}
	}
	for i := range houses {
		if houses[i] >= input {
			fmt.Println(i)
			break
		}
	}

	// Part 2
	clear(houses)
	for e := 1; e < limit; e++ {
		for h := e; h < limit && h <= e*50; h += e {
			houses[h] += e * 11
		}
	}
	for i := range houses {
		if houses[i] >= input {
			fmt.Println(i)
			break
		}
	}
}
