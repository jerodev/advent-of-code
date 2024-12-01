package util

func CountOccurrences[V comparable](s []V, value V) int {
	var count int
	for i := range s {
		if s[i] == value {
			count++
		}
	}

	return count
}
