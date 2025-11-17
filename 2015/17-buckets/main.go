package main

import (
	"fmt"
	"math"
)

var buckets = []int{33, 14, 18, 20, 45, 35, 16, 35, 1, 13, 18, 13, 50, 44, 48, 6, 24, 41, 30, 42}
var minCount [16]int

const target = 150

func main() {
	possibilities := int(math.Pow(2, float64(len(buckets))))
	var count, bucket, bucketCount, tmp int

	for i := range possibilities {
		bucket, bucketCount = 0, 0

		for j := range buckets {
			tmp = int(math.Pow(2, float64(j)))
			if int(math.Pow(2, float64(j)))&i == tmp {
				bucket += buckets[j]
				bucketCount++

				// Overflow!
				if bucket > target {
					break
				}
			}
		}

		if bucket == target {
			count++
			minCount[bucketCount]++
		}
	}

	fmt.Println(count)
	fmt.Println(minCount)
}
