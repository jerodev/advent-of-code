package util

import "golang.org/x/exp/constraints"

func Abs[I constraints.Integer](v I) I {
	if v < 0 {
		return -v
	}

	return v
}

func IntLength(v int) int {
	x, count := 10, 1

	for x <= v {
		x *= 10
		count++
	}

	return count
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
