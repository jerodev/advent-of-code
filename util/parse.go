package util

import (
	"strconv"
	"strings"
)

func IntsToStrings(ints []int) []string {
	strings := make([]string, len(ints))

	for i := range ints {
		strings[i] = strconv.Itoa(ints[i])
	}

	return strings
}

func MustAtoi(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

func StringToInts(s, delimiter string) []int {
	if delimiter == "" {
		ints := make([]int, len(s))
		for i, b := range s {
			ints[i] = int(b - '0')
		}

		return ints
	}

	parts := strings.Split(strings.TrimSpace(s), delimiter)
	ints := make([]int, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		ints = append(ints, MustAtoi(part))
	}

	return ints
}
