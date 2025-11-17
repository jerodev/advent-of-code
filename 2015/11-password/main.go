package main

import (
	"fmt"
	"slices"
)

func main() {
	pw := []byte("hxbxxyzz")

	for {
		if incrementPassword(&pw); isValidPassword(pw) {
			break
		}
	}

	fmt.Println(string(pw))
}

func incrementPassword(pw *[]byte) {
	for i := len(*pw) - 1; i >= 0; i-- {
		if (*pw)[i] < 'z' {
			(*pw)[i]++
			return
		} else {
			(*pw)[i] = 'a'
		}
	}
}

func isValidPassword(pw []byte) bool {
	var straight bool
	var pairs []byte

	for i := range pw {
		if pw[i] == 'i' || pw[i] == 'o' || pw[i] == 'l' {
			return false
		}

		if i > 0 && pw[i-1] == pw[i] && !slices.Contains(pairs, pw[i]) {
			pairs = append(pairs, pw[i])
		}

		straight = straight || (i > 1 && pw[i-2]+1 == pw[i-1] && pw[i-1]+1 == pw[i])
	}

	return straight && len(pairs) > 1
}
