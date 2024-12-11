package main

import (
	"advent-of-code/util"
	"fmt"
	"math"
)

func main() {
	stones := util.StringToInts(util.ReadFileFromArgs(), " ")

	for i := 0; i < 25; i++ {
		stones = blink(stones)
		fmt.Println(len(stones))
	}

	fmt.Println(len(stones))
}

func blink(stones []int) []int {
	result := make([]int, 0, len(stones))
	var l, left, mod int

	for i := range stones {
		if stones[i] == 0 {
			result = append(result, 1)
		} else if l = util.IntLength(stones[i]); l%2 == 0 {
			mod = int(math.Pow(10, float64(l/2)))
			left = int(stones[i] / mod)
			result = append(result, left)
			result = append(result, stones[i]-left*mod)
		} else {
			result = append(result, stones[i]*2024)
		}
	}

	return result
}
