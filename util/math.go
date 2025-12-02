package util

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Abs[I constraints.Integer](v I) I {
	if v < 0 {
		return -v
	}

	return v
}

func FloorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return (a - b + 1) / b
}

func IntLength[I constraints.Integer](v I) int {
	var x I = 10
	count := 1

	for x <= v {
		x *= 10
		count++
	}

	return count
}

func IntPow[I constraints.Integer](n, m I) I {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	var i I
	for i = 2; i <= m; i++ {
		result *= n
	}
	return result
}

// NthDigit returns the nth digit of a number starting from the right
func NthDigit(v, i int) int {
	return int(math.Floor(float64(v)/math.Pow(10, float64(i)-1))) % 10
}
