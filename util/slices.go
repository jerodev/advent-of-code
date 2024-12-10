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

// Unique returns a new slice with only unique values
func Unique[V comparable](s []V) []V {
	ns := make([]V, 0, len(s))

	var duplicate bool
	for i := range s {
		duplicate = false

		for j := range ns {
			if s[i] == ns[j] {
				duplicate = true
				break
			}
		}

		if !duplicate {
			ns = append(ns, s[i])
		}
	}

	return ns
}
