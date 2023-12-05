package util

import (
	"strconv"
	"strings"
)

func StringToInts(s, delimiter string) []int {
	parts := strings.Split(strings.TrimSpace(s), delimiter)
	ints := make([]int, len(parts))

	for i, part := range parts {
		ints[i], _ = strconv.Atoi(part)
	}

	return ints
}
