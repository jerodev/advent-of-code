package main

import (
	"advent-of-code/util"
	"fmt"
	"math"
)

func main() {
	stones := util.StringToInts(util.ReadFileFromArgs(), " ")

	pebbles := make(map[int]int, len(stones))
	for i := range stones {
		pebbles[stones[i]] += 1
	}

	for i := 0; i < 75; i++ {
		pebbles = blink(pebbles)
	}

	var count int
	for i := range pebbles {
		count += pebbles[i]
	}

	fmt.Println(count)
}

func blink(stones map[int]int) map[int]int {
	pebbles := map[int]int{}
	var l, left, mod int

	for k, v := range stones {
		if k == 0 {
			pebbles[1] += v
		} else if l = util.IntLength(k); l%2 == 0 {
			mod = int(math.Pow(10, float64(l/2)))
			left = int(k / mod)

			pebbles[left] += v
			pebbles[k-left*mod] += v
		} else {
			pebbles[k*2024] += v
		}
	}

	return pebbles
}
