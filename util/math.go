package util

import "golang.org/x/exp/constraints"

func Abs[I constraints.Integer](v I) I {
	if v < 0 {
		return -v
	}

	return v
}
